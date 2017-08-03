package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	retryMaxN = 5
)

//Client 主动调用的客户端
type Client struct {
	corpID      string
	secret      string
	httpClient  *http.Client //
	tokenServer AccesstokenServer
}

//ResponseError 微信返回的错误信息
type ResponseError struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type reply struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

//NewClient 返回一个client实例
//如果没有指定token服务器，就自己启动一个，请确保不要重复启动token服务器
//如果htc=nil，则默认使用http.DefaultClient，如果要实现代理等，可自行传入client
func NewClient(corpid, secret string, srv AccesstokenServer, htc *http.Client) (*Client, error) {
	if corpid == "" || secret == "" {
		return nil, errors.New("企业号ID或者密钥为空")
	}

	var c = &Client{corpID: corpid, secret: secret}
	if htc == nil {
		htc = http.DefaultClient
	}
	c.httpClient = htc
	if srv == nil {
		c.tokenServer = NewDefaultAccessTokenServer(c.corpID, c.secret, c.httpClient)

	}
	return c, nil
}

//AccessToken 获取token
func (c *Client) AccessToken() (string, error) {
	return c.tokenServer.Token()
}

//CorpID CorpID
func (c *Client) CorpID() string {
	return c.corpID
}

//SendGetRequest 发送get请求
//requrl 是完整的格式字符串，例如
//https://qyapi.weixin.qq.com/cgi-bin/user/getuserinfo?access_token=%s&code=CODE
func (c *Client) SendGetRequest(reqURL string) ([]byte, error) {
	for i := 0; i < retryMaxN; i++ {
		token, err := c.tokenServer.Token()
		if err != nil {
			c.tokenServer.RefreshToken()
			return nil, fmt.Errorf("获取token失败：%v\r\n ", err)
		}

		newURL := fmt.Sprintf(reqURL, token)
		if _, err := url.Parse(newURL); err != nil {
			return nil, fmt.Errorf("url is wrong:%s       %v", reqURL, err)
		}

		r, err := c.httpClient.Get(newURL)
		if err != nil {
			return nil, fmt.Errorf("get请求失败：%v", err)
		}
		defer r.Body.Close()
		reply, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}
		var result ResponseError
		if err := json.Unmarshal(reply, &result); err != nil {
			return nil, err
		}

		switch result.ErrCode {
		case 0:
			return reply, nil
		case -1:
			i--
			continue
		case 42001, 40001, 40014: // access_token timeout and retry
			c.tokenServer.RefreshToken()
			continue
		default:
			return nil, fmt.Errorf("WeiXin send get request reply[%d]: %s", result.ErrCode, result.ErrMsg)
		}
	}

	return nil, fmt.Errorf("WeiXin post request too many times:%s" + reqURL)
}

//PostJSON post json数据到reqURL，注意，Url中access_token要占位下来,
//requrl 是完整的格式字符串，例如
//https://qyapi.weixin.qq.com/cgi-bin/user/getuserinfo?access_token=%s&code=CODE
func (c *Client) PostJSON(reqURL string, v interface{}) ([]byte, error) {
	//所有post请求都要用到accesstoken，如果没有启用accesstoken，返回错误
	data, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("需要post的数据不能通过json转化：%v", err)
	}

	for i := 0; i < retryMaxN; i++ {
		token, err := c.tokenServer.Token()
		if err != nil {
			c.tokenServer.RefreshToken()
			return nil, fmt.Errorf("获取token失败：%v\r\n ", err)
		}

		newURL := fmt.Sprintf(reqURL, token)
		if _, err := url.Parse(newURL); err != nil {
			return nil, errors.New("reqURL is wrong")
		}
		resp, err := c.httpClient.Post(newURL, "application/json; charset=utf-8", bytes.NewReader(data))
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()

		reply, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		var result ResponseError
		if err := json.Unmarshal(reply, &result); err != nil {
			return nil, err
		}
		switch result.ErrCode {
		case -1:
			i--
			continue
		case 0:
			return reply, nil
		case 42001, 40001, 40014: // access_token timeout and retry
			c.tokenServer.RefreshToken()
			continue
		default:
			return nil, fmt.Errorf("WeiXin reply[%d]: %s", result.ErrCode, result.ErrMsg)
		}

	}

	return nil, errors.New("WeiXin post request too many times:" + reqURL)
}

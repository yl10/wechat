package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var _AccesstokenServer = DefaultAccessTokenServer{}

//AccessToken AccessToken
//expires为到期时间
type AccessToken struct {
	AccessToken string    `json:"access_token"`
	Expires     time.Time //有效期
}

//AccesstokenServer AccesstokenServer接口,如果是第三方获取，只需要实现这个接口
type AccesstokenServer interface {
	Token() (string, error)           //获取token
	Check(corpID, secret string) bool //检查是否一个管理组
	BadToken()                        //如果token没有过期但是被破坏了不能使用，则通知服务进行重新拉取
}

//DefaultAccessTokenServer 默认的本地token服务
type DefaultAccessTokenServer struct {
	corpID    string
	secret    string
	lost      bool
	tokenChan chan AccessToken
	err       error //读取token时最后一次错误信息
	isRun     bool
}

//Check 实现接口
func (dats DefaultAccessTokenServer) Check(corpID, secret string) bool {

	return dats.corpID == corpID && dats.secret == secret
}

//BadToken 实现接口
func (dats DefaultAccessTokenServer) BadToken() {

	dats.lost = true
}

//Token 获取token,如果返回错误的话，可以设置BadToken()后再执行
func (dats DefaultAccessTokenServer) Token() (string, error) {

	if !dats.isRun {

		go dats.createAccessToken()
	}

	token := <-dats.tokenChan
	if token.AccessToken == "" {
		return "", fmt.Errorf("没能获取到token：%s", dats.err)
	}

	return token.AccessToken, nil
}

func (dats *DefaultAccessTokenServer) createAccessToken() {
	token := AccessToken{"", time.Now()}
	dats.isRun = true
	for {

		if time.Since(token.Expires).Seconds() >= 0 || dats.lost {

			var expires time.Duration

			token.AccessToken, expires, dats.err = AuthAccessToken(dats.corpID, dats.secret)
			token.Expires = time.Now().Add(expires - 100) //减少100秒
			dats.lost = false

		}

		dats.tokenChan <- token

	}
}

//NewDefaultAccessTokenServer 初始化默认tokenserver
func NewDefaultAccessTokenServer(corpID, secret string, client *http.Client) *DefaultAccessTokenServer {
	ts := DefaultAccessTokenServer{
		corpID:    corpID,
		secret:    secret,
		tokenChan: make(chan AccessToken),
	}
	if client == nil {
		client = http.DefaultClient
	}
	go ts.createAccessToken()
	return &ts
}

//AuthAccessToken 获取token的基本方法
func AuthAccessToken(appid string, secret string) (token string, tlong time.Duration, errResult error) {
	var res struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int64  `json:"expires_in"`
	}

	resp, err := http.Get(fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s", appid, secret))
	if err != nil {
		errResult = fmt.Errorf("Get access token failed: %v", err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errResult = fmt.Errorf("Read access token failed: %v", err)

		return

	}
	if err = json.Unmarshal(body, &res); err != nil {
		var clienterr ResponseError
		err = json.Unmarshal(body, &clienterr)
		if err == nil {
			errResult = fmt.Errorf("获取Accesstoken失败：%v\n", clienterr)
			return
		}
		errResult = fmt.Errorf("Parse access token failed:%v ", err)
		return
	}
	return res.AccessToken, time.Duration(res.ExpiresIn * 1000 * 1000 * 1000), nil

}

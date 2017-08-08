package jssdk

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/yl10/wechat/qy/client"
	"github.com/yl10/wechat/util"
)

const (
	jsapiURI     = "https://qyapi.weixin.qq.com/cgi-bin/get_jsapi_ticket?access_token=%s"
	ticketErrMsg = "获取tikect失败：%v"
)

//JsapiTicket JsapiTicket
type JsapiTicket struct {
	Ticket  string
	Expires time.Time
}

//JsWxConfig JsWxConfig
type JsWxConfig struct {
	AppID     string `json:"appId"`
	NonceStr  string `json:"nonceStr"`
	TimeStamp int64  `json:"timestamp"`
	Signature string `json:"signature"`
}

//JsTicketServer ticket服务
type JsTicketServer struct {
	isRun      bool
	client     *client.Client
	ticketChan chan JsapiTicket
	err        error
}

//RunTicketServer 启动jstikect服务
func RunTicketServer(c *client.Client) (*JsTicketServer, error) {
	if c == nil {
		return nil, fmt.Errorf("没有传入微信client")
	}
	jts := &JsTicketServer{
		client:     c,
		ticketChan: make(chan JsapiTicket),
	}
	return jts, nil
}

//Ticket 输出ticket
func (j JsTicketServer) Ticket() (string, error) {

	if !j.isRun {
		go j.createJsAPITicket()
	}

	ticket := <-j.ticketChan

	if time.Since(ticket.Expires).Seconds() >= 0 {
		ticket = <-j.ticketChan
	}

	if ticket.Ticket == "" {
		return "", fmt.Errorf("未能获取到ticket：%v", j.err)
	}
	return ticket.Ticket, nil

}

func (j JsTicketServer) createJsAPITicket() {
	ticket := JsapiTicket{"", time.Now()}
	j.isRun = true

	for {
		if time.Since(ticket.Expires).Seconds() >= 0 {
			var expires time.Duration
			ticket.Ticket, expires, j.err = j.getTicket()
			ticket.Expires = time.Now().Add(expires - 100)

		}
		j.ticketChan <- ticket
	}

}

//GetTicket 获取ticket
func (j JsTicketServer) getTicket() (ticket string, tlong time.Duration, errResult error) {

	var res struct {
		Ticket    string `json:"ticket"`
		ExpiresIn int64  `json:"expires_in"`
	}

	if j.client == nil {
		errResult = fmt.Errorf(ticketErrMsg, "微信clent没有初始化")
		return
	}
	resp, err := j.client.SendGetRequest(jsapiURI)
	if err != nil {
		errResult = fmt.Errorf(ticketErrMsg, err)
		return
	}

	err = json.Unmarshal(resp, &res)
	if err != nil {
		errResult = fmt.Errorf(ticketErrMsg, err)
		return
	}

	return res.Ticket, time.Duration(res.ExpiresIn * 1000 * 1000 * 1000), nil

}

/*Jssha1 js签名方法
签名生成规则如下：

参与签名的字段包括有效的 jsapi_ticket（获取方式详见企业微信 JSSDK 文档）， noncestr （随机字符串，由开发者随机生成），timestamp （由开发者生成的当前时间戳）， url（当前网页的URL，不包含#及其后面部分。注意：对于没有只有域名没有 path 的 URL ，浏览器会自动加上 / 作为 path，如打开 http://qq.com 则获取到的 URL 为 http://qq.com/）。
对所有待签名参数按照字段名的 ASCII 码从小到大排序（字典序）后，使用 URL 键值对的格式（即key1=value1&key2=value2…）拼接成字符串 string1。这里需要注意的是所有参数名均为小写字符。
接下来对 string1 作 sha1 加密，字段名和字段值都采用原始值，不进行 URL 转义。即 signature=sha1(string1)。
*/
func Jssha1(noncestr, jsapiTicket string, timestamp int64, url string) string {
	sting1 := "jsapi_ticket=" + jsapiTicket + "&noncestr=" + noncestr + "&timestamp=" + strconv.FormatInt(timestamp, 10) + "&url=" + url
	h := sha1.New()
	h.Write([]byte(sting1))
	return fmt.Sprintf("%x", h.Sum(nil))
}

//GetJsWxconfig 获取微信jssdk配置文件
func (j JsTicketServer) GetJsWxconfig(urlstring string) (JsWxConfig, error) {
	cfg := JsWxConfig{}
	ticket, err := j.Ticket()
	if err != nil {
		return cfg, err
	}
	//判断url是否有错，并去除#后面的部分
	u, err := url.Parse(urlstring)
	if err != nil {
		return cfg, err
	}
	u.Fragment = ""

	cfg.AppID = j.client.CorpID()
	cfg.NonceStr = util.RandomAlphanumeric(10)
	cfg.TimeStamp = time.Now().Unix()
	cfg.Signature = Jssha1(cfg.NonceStr, ticket, cfg.TimeStamp, u.String())
	return cfg, nil
}

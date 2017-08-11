package qy

import (
	"net/http"

	"github.com/yl10/wechat/qy/client"
	"github.com/yl10/wechat/qy/oauth"
)

const (
	noncestr = "WxjahsoxMsklakzzzXLL"
)

//Wx 企业微信
type Wx struct {
	client *client.Client
}

/*
type JsApiTicket struct {
	CorpId    string
	NonceStr  string
	TimeStamp string
	Signature string
}*/

type JsApiTicket struct {
	CorpId   string
	JsTicket string
}

//NewWx 初始化一个微信
func NewWx(corpid, secret string, srv client.AccesstokenServer, htc *http.Client) (*Wx, error) {
	c, err := client.NewClient(corpid, secret, srv, htc)
	wx := &Wx{client: c}
	return wx, err

}

//GetOauthUserInfo 通过code获取用户信息
func (wx *Wx) GetOauthUserInfo(code string) (*oauth.User, error) {
	return oauth.GetUserInfo(wx.client, code)
}

//GetOAuthURI 获取跳转的uri
func (wx *Wx) GetOAuthURI(redirectURI string, state ...string) string {
	return oauth.GetOAuthURI(wx.client.CorpID(), redirectURI, state...)
}

//GetClient 获取 Client
func (wx *Wx) GetClient() *client.Client {
	return wx.client
}

//GetJsAPITicket 获取jsapi_ticket
func (wx *Wx) GetJsAPITicket(url string) (*JsApiTicket, error) {
	jsticket, err := wx.client.GetJsTicket()
	if err != nil {
		return nil, err
	}
	//return wx.Signature(jsticket, url), nil
	var ret = new(JsApiTicket)
	ret.JsTicket = jsticket
	ret.CorpId = wx.client.CorpID()
	return ret, nil
}

/*
//Signature 生成Signature
func (wx *Wx) Signature(str, url string) *JsApiTicket {
	cur := time.Now()
	timestamp := strconv.FormatInt(cur.UnixNano()/1000000000, 10)
	string1 := fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%s&url=%s", str, noncestr, timestamp, url)
	h := sha1.New()
	h.Write([]byte(string1))
	bs := h.Sum(nil)
	siginature := fmt.Sprintf("%x", bs)
	var ret = new(JsApiTicket)
	ret.CorpId = wx.client.CorpID()
	ret.NonceStr = noncestr
	ret.TimeStamp = timestamp
	ret.Signature = siginature
	return ret
}*/

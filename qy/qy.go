package qy

import (
	"net/http"

	"github.com/yl10/wechat/qy/client"
	"github.com/yl10/wechat/qy/jssdk"
	"github.com/yl10/wechat/qy/oauth"
)

//Wx 企业微信
type Wx struct {
	client            *client.Client
	jsapiticketserver *jssdk.JsTicketServer
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

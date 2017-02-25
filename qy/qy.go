package qy

import (
	"github.com/yl10/tencent/wechat/qy/client"
	"github.com/yl10/tencent/wechat/qy/message"
	"github.com/yl10/tencent/wechat/qy/oauth"
)

//Wx 企业微信
type Wx struct {
	client *client.Client
}

//SendMessage 发送消息
func (wx *Wx) SendMessage(msg message.Message) (*message.Result, error) {
	return message.SendMessage(wx.client, msg)
}

//GetOauthUserInfo 通过code获取用户信息
func (wx *Wx) GetOauthUserInfo(code string) (*oauth.User, error) {
	return oauth.GetUserInfo(wx.client, code)
}

//GetOAuthURI 获取跳转的uri
func (wx *Wx) GetOAuthURI(redirectURI string, state ...string) string {
	return oauth.GetOAuthURI(wx.client.CorpID(), redirectURI, state...)
}

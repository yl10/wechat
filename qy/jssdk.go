package qy

import (
	"github.com/yl10/wechat/qy/jssdk"
)

//RunJsAPITicketServer 启动jsapiticket服务
func (wx *Wx) RunJsAPITicketServer() (err error) {

	wx.jsapiticketserver, err = jssdk.RunTicketServer(wx.client)
	return
}

//GetJsWxConfig 获取GetJsWxConfig
func (wx *Wx) GetJsWxConfig(url string) (jssdk.JsWxConfig, error) {
	if wx.jsapiticketserver == nil {
		err := wx.RunJsAPITicketServer()
		return jssdk.JsWxConfig{}, err
	}
	return wx.jsapiticketserver.GetJsWxconfig(url)
}

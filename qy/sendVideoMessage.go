package qy

import (
	"github.com/yl10/wechat/qy/message"
)

//SendVideo 发送图像消息
func (wx *Wx) SendVideo(agenid int64, toAll bool, user, party, tag []string, mediaID, title, description string) (*message.Result, error) {
	msg := message.NewVideo(agenid, toAll, false, user, party, tag, mediaID, title, description)
	return message.SendMessage(wx.client, *msg)
}

//SendSafeVideo 发送保密图像消息
func (wx *Wx) SendSafeVideo(agenid int64, toAll, issafe bool, user, party, tag []string, mediaID, title, description string) (*message.Result, error) {
	msg := message.NewVideo(agenid, toAll, issafe, user, party, tag, mediaID, title, description)
	return message.SendMessage(wx.client, *msg)
}

//SendVideoToAll 发送图像消息给所有人
func (wx *Wx) SendVideoToAll(agenid int64, mediaID, title, description string) (*message.Result, error) {
	msg := message.NewVideo(agenid, true, false, nil, nil, nil, mediaID, title, description)
	return message.SendMessage(wx.client, *msg)
}

//SendSafeVideoToAll 发送保密图像消息给所有人
func (wx *Wx) SendSafeVideoToAll(agenid int64, mediaID, title, description string) (*message.Result, error) {
	msg := message.NewVideo(agenid, true, true, nil, nil, nil, mediaID, title, description)
	return message.SendMessage(wx.client, *msg)
}

//SendVideoToUser 发送图像消息给用户
func (wx *Wx) SendVideoToUser(agenid int64, mediaID, title, description string, user ...string) (*message.Result, error) {
	msg := message.NewVideo(agenid, false, false, user, nil, nil, mediaID, title, description)
	return message.SendMessage(wx.client, *msg)
}

//SendSafeVideoToUser 发送保密图像消息给用户
func (wx *Wx) SendSafeVideoToUser(agenid int64, mediaID, title, description string, user ...string) (*message.Result, error) {
	msg := message.NewVideo(agenid, false, true, user, nil, nil, mediaID, title, description)
	return message.SendMessage(wx.client, *msg)
}

//SendVideoToParty 发送图像消息给部门
func (wx *Wx) SendVideoToParty(agenid int64, mediaID, title, description string, party ...string) (*message.Result, error) {
	msg := message.NewVideo(agenid, false, false, nil, party, nil, mediaID, title, description)
	return message.SendMessage(wx.client, *msg)
}

//SendSafeVideoToParty 发送保密图像消息给部门
func (wx *Wx) SendSafeVideoToParty(agenid int64, mediaID, title, description string, party ...string) (*message.Result, error) {
	msg := message.NewVideo(agenid, false, true, nil, party, nil, mediaID, title, description)
	return message.SendMessage(wx.client, *msg)
}

//SendVideoToTags 发送图像消息给用户
func (wx *Wx) SendVideoToTags(agenid int64, mediaID, title, description string, tag ...string) (*message.Result, error) {
	msg := message.NewVideo(agenid, false, false, nil, nil, tag, mediaID, title, description)
	return message.SendMessage(wx.client, *msg)
}

//SendSafeVideoToTags 发送保密图像消息给用户组
func (wx *Wx) SendSafeVideoToTags(agenid int64, mediaID, title, description string, tag ...string) (*message.Result, error) {
	msg := message.NewVideo(agenid, false, true, nil, nil, tag, mediaID, title, description)
	return message.SendMessage(wx.client, *msg)
}

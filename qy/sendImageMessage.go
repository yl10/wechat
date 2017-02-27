package qy

import (
	"github.com/yl10/wechat/qy/message"
)

//SendImage 发送图像消息
func (wx *Wx) SendImage(agenid int64, toAll bool, user, party, tag []string, mediaID string) (*message.Result, error) {
	msg := message.NewImage(agenid, toAll, false, user, party, tag, mediaID)
	return message.SendMessage(wx.client, *msg)
}

//SendSafeImage 发送保密图像消息
func (wx *Wx) SendSafeImage(agenid int64, toAll, issafe bool, user, party, tag []string, mediaID string) (*message.Result, error) {
	msg := message.NewImage(agenid, toAll, issafe, user, party, tag, mediaID)
	return message.SendMessage(wx.client, *msg)
}

//SendImageToAll 发送图像消息给所有人
func (wx *Wx) SendImageToAll(agenid int64, mediaID string) (*message.Result, error) {
	msg := message.NewImage(agenid, true, false, nil, nil, nil, mediaID)
	return message.SendMessage(wx.client, *msg)
}

//SendSafeImageToAll 发送保密图像消息给所有人
func (wx *Wx) SendSafeImageToAll(agenid int64, mediaID string) (*message.Result, error) {
	msg := message.NewImage(agenid, true, true, nil, nil, nil, mediaID)
	return message.SendMessage(wx.client, *msg)
}

//SendImageToUser 发送图像消息给用户
func (wx *Wx) SendImageToUser(agenid int64, mediaID string, user ...string) (*message.Result, error) {
	msg := message.NewImage(agenid, false, false, user, nil, nil, mediaID)
	return message.SendMessage(wx.client, *msg)
}

//SendSafeImageToUser 发送保密图像消息给用户
func (wx *Wx) SendSafeImageToUser(agenid int64, mediaID string, user ...string) (*message.Result, error) {
	msg := message.NewImage(agenid, false, true, user, nil, nil, mediaID)
	return message.SendMessage(wx.client, *msg)
}

//SendImageToParty 发送图像消息给部门
func (wx *Wx) SendImageToParty(agenid int64, mediaID string, party ...string) (*message.Result, error) {
	msg := message.NewImage(agenid, false, false, nil, party, nil, mediaID)
	return message.SendMessage(wx.client, *msg)
}

//SendSafeImageToParty 发送保密图像消息给部门
func (wx *Wx) SendSafeImageToParty(agenid int64, mediaID string, party ...string) (*message.Result, error) {
	msg := message.NewImage(agenid, false, true, nil, party, nil, mediaID)
	return message.SendMessage(wx.client, *msg)
}

//SendImageToTags 发送图像消息给用户
func (wx *Wx) SendImageToTags(agenid int64, mediaID string, tag ...string) (*message.Result, error) {
	msg := message.NewImage(agenid, false, false, nil, nil, tag, mediaID)
	return message.SendMessage(wx.client, *msg)
}

//SendSafeImageToTags 发送保密图像消息给用户组
func (wx *Wx) SendSafeImageToTags(agenid int64, mediaID string, tag ...string) (*message.Result, error) {
	msg := message.NewImage(agenid, false, true, nil, nil, tag, mediaID)
	return message.SendMessage(wx.client, *msg)
}

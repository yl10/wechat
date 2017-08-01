package qy

import (
	"github.com/yl10/wechat/qy/message"
)

//SendTextCard 发送图像消息
func (wx *Wx) SendTextCard(agenid int64, toAll bool, user, party, tag []string, title, description, url, btntxt string) (*message.Result, error) {
	msg := message.NewTextCard(agenid, toAll, false, user, party, tag, title, description, url, btntxt)
	return message.SendMessage(wx.client, *msg)
}

//SendSafeTextCard 发送保密图像消息
func (wx *Wx) SendSafeTextCard(agenid int64, toAll, issafe bool, user, party, tag []string, title, description, url, btntxt string) (*message.Result, error) {
	msg := message.NewTextCard(agenid, toAll, issafe, user, party, tag, title, description, url, btntxt)
	return message.SendMessage(wx.client, *msg)
}

//SendTextCardToAll 发送图像消息给所有人
func (wx *Wx) SendTextCardToAll(agenid int64, title, description, url, btntxt string) (*message.Result, error) {
	msg := message.NewTextCard(agenid, true, false, nil, nil, nil, title, description, url, btntxt)
	return message.SendMessage(wx.client, *msg)
}

//SendSafeTextCardToAll 发送保密图像消息给所有人
func (wx *Wx) SendSafeTextCardToAll(agenid int64, title, description, url, btntxt string) (*message.Result, error) {
	msg := message.NewTextCard(agenid, true, true, nil, nil, nil, title, description, url, btntxt)
	return message.SendMessage(wx.client, *msg)
}

//SendTextCardToUser 发送图像消息给用户
func (wx *Wx) SendTextCardToUser(agenid int64, title, description, url, btntxt string, user ...string) (*message.Result, error) {
	msg := message.NewTextCard(agenid, false, false, user, nil, nil, title, description, url, btntxt)
	return message.SendMessage(wx.client, *msg)
}

//SendSafeTextCardToUser 发送保密图像消息给用户
func (wx *Wx) SendSafeTextCardToUser(agenid int64, title, description, url, btntxt string, user ...string) (*message.Result, error) {
	msg := message.NewTextCard(agenid, false, true, user, nil, nil, title, description, url, btntxt)
	return message.SendMessage(wx.client, *msg)
}

//SendTextCardToParty 发送图像消息给部门
func (wx *Wx) SendTextCardToParty(agenid int64, title, description, url, btntxt string, party ...string) (*message.Result, error) {
	msg := message.NewTextCard(agenid, false, false, nil, party, nil, title, description, url, btntxt)
	return message.SendMessage(wx.client, *msg)
}

//SendSafeTextCardToParty 发送保密图像消息给部门
func (wx *Wx) SendSafeTextCardToParty(agenid int64, title, description, url, btntxt string, party ...string) (*message.Result, error) {
	msg := message.NewTextCard(agenid, false, true, nil, party, nil, title, description, url, btntxt)
	return message.SendMessage(wx.client, *msg)
}

//SendTextCardToTags 发送图像消息给用户
func (wx *Wx) SendTextCardToTags(agenid int64, title, description, url, btntxt string, tag ...string) (*message.Result, error) {
	msg := message.NewTextCard(agenid, false, false, nil, nil, tag, title, description, url, btntxt)
	return message.SendMessage(wx.client, *msg)
}

//SendSafeTextCardToTags 发送保密图像消息给用户组
func (wx *Wx) SendSafeTextCardToTags(agenid int64, title, description, url, btntxt string, tag ...string) (*message.Result, error) {
	msg := message.NewTextCard(agenid, false, true, nil, nil, tag, title, description, url, btntxt)
	return message.SendMessage(wx.client, *msg)
}

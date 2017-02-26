package qy

import (
	"github.com/yl10/wechat/qy/message"
)

//SendMessage 发送消息
func (wx *Wx) SendMessage(msg message.Message) (*message.Result, error) {
	return message.SendMessage(wx.client, msg)
}

//SendText 发送文本消息
func (wx *Wx) SendText(agenid int64, toAll bool, user, party, tag []string, textcontent string) (*message.Result, error) {
	msg := message.NewText(agenid, toAll, false, user, party, tag, textcontent)
	return message.SendMessage(wx.client, *msg)
}

//SendSafeText 发送保密文本消息
func (wx *Wx) SendSafeText(agenid int64, toAll, issafe bool, user, party, tag []string, textcontent string) (*message.Result, error) {
	msg := message.NewText(agenid, toAll, issafe, user, party, tag, textcontent)
	return message.SendMessage(wx.client, *msg)
}

//SendTextToAll 发送文本消息给所有人
func (wx *Wx) SendTextToAll(agenid int64, textcontent string) (*message.Result, error) {
	msg := message.NewText(agenid, true, false, nil, nil, nil, textcontent)
	return message.SendMessage(wx.client, *msg)
}

//SendSafeTextToAll 发送保密文本消息给所有人
func (wx *Wx) SendSafeTextToAll(agenid int64, textcontent string) (*message.Result, error) {
	msg := message.NewText(agenid, true, true, nil, nil, nil, textcontent)
	return message.SendMessage(wx.client, *msg)
}

//SendTextToUser 发送文本消息给用户
func (wx *Wx) SendTextToUser(agenid int64, textcontent string, user ...string) (*message.Result, error) {
	msg := message.NewText(agenid, false, false, user, nil, nil, textcontent)
	return message.SendMessage(wx.client, *msg)
}

//SendSafeTextToUser 发送保密文本消息给用户
func (wx *Wx) SendSafeTextToUser(agenid int64, textcontent string, user ...string) (*message.Result, error) {
	msg := message.NewText(agenid, false, true, user, nil, nil, textcontent)
	return message.SendMessage(wx.client, *msg)
}

//SendTextToParty 发送文本消息给部门
func (wx *Wx) SendTextToParty(agenid int64, textcontent string, party ...string) (*message.Result, error) {
	msg := message.NewText(agenid, false, false, nil, party, nil, textcontent)
	return message.SendMessage(wx.client, *msg)
}

//SendSafeTextToParty 发送保密文本消息给部门
func (wx *Wx) SendSafeTextToParty(agenid int64, textcontent string, party ...string) (*message.Result, error) {
	msg := message.NewText(agenid, false, true, nil, party, nil, textcontent)
	return message.SendMessage(wx.client, *msg)
}

//SendTextToTags 发送文本消息给用户
func (wx *Wx) SendTextToTags(agenid int64, textcontent string, tag ...string) (*message.Result, error) {
	msg := message.NewText(agenid, false, false, nil, nil, tag, textcontent)
	return message.SendMessage(wx.client, *msg)
}

//SendSafeTextToTags 发送保密文本消息给用户组
func (wx *Wx) SendSafeTextToTags(agenid int64, textcontent string, tag ...string) (*message.Result, error) {
	msg := message.NewText(agenid, false, true, nil, nil, tag, textcontent)
	return message.SendMessage(wx.client, *msg)
}

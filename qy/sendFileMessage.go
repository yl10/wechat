package qy

import (
	"github.com/yl10/wechat/qy/message"
)

//SendFile 发送文件消息
func (wx *Wx) SendFile(agenid int64, toAll bool, user, party, tag []string, mediaID string) (*message.Result, error) {
	msg := message.NewFile(agenid, toAll, false, user, party, tag, mediaID)
	return message.SendMessage(wx.client, *msg)
}

//SendSafeFile 发送保密文件消息
func (wx *Wx) SendSafeFile(agenid int64, toAll, issafe bool, user, party, tag []string, mediaID string) (*message.Result, error) {
	msg := message.NewFile(agenid, toAll, issafe, user, party, tag, mediaID)
	return message.SendMessage(wx.client, *msg)
}

//SendFileToAll 发送文件消息给所有人
func (wx *Wx) SendFileToAll(agenid int64, mediaID string) (*message.Result, error) {
	msg := message.NewFile(agenid, true, false, nil, nil, nil, mediaID)
	return message.SendMessage(wx.client, *msg)
}

//SendSafeFileToAll 发送保密文件消息给所有人
func (wx *Wx) SendSafeFileToAll(agenid int64, mediaID string) (*message.Result, error) {
	msg := message.NewFile(agenid, true, true, nil, nil, nil, mediaID)
	return message.SendMessage(wx.client, *msg)
}

//SendFileToUser 发送文件消息给用户
func (wx *Wx) SendFileToUser(agenid int64, mediaID string, user ...string) (*message.Result, error) {
	msg := message.NewFile(agenid, false, false, user, nil, nil, mediaID)
	return message.SendMessage(wx.client, *msg)
}

//SendSafeFileToUser 发送保密文件消息给用户
func (wx *Wx) SendSafeFileToUser(agenid int64, mediaID string, user ...string) (*message.Result, error) {
	msg := message.NewFile(agenid, false, true, user, nil, nil, mediaID)
	return message.SendMessage(wx.client, *msg)
}

//SendFileToParty 发送文件消息给部门
func (wx *Wx) SendFileToParty(agenid int64, mediaID string, party ...string) (*message.Result, error) {
	msg := message.NewFile(agenid, false, false, nil, party, nil, mediaID)
	return message.SendMessage(wx.client, *msg)
}

//SendSafeFileToParty 发送保密文件消息给部门
func (wx *Wx) SendSafeFileToParty(agenid int64, mediaID string, party ...string) (*message.Result, error) {
	msg := message.NewFile(agenid, false, true, nil, party, nil, mediaID)
	return message.SendMessage(wx.client, *msg)
}

//SendFileToTags 发送文件消息给用户
func (wx *Wx) SendFileToTags(agenid int64, mediaID string, tag ...string) (*message.Result, error) {
	msg := message.NewFile(agenid, false, false, nil, nil, tag, mediaID)
	return message.SendMessage(wx.client, *msg)
}

//SendSafeFileToTags 发送保密文件消息给用户组
func (wx *Wx) SendSafeFileToTags(agenid int64, mediaID string, tag ...string) (*message.Result, error) {
	msg := message.NewFile(agenid, false, true, nil, nil, tag, mediaID)
	return message.SendMessage(wx.client, *msg)
}

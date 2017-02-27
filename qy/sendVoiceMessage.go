package qy

import (
	"github.com/yl10/wechat/qy/message"
)

//SendVoice 发送语音消息
func (wx *Wx) SendVoice(agenid int64, toAll bool, user, party, tag []string, mediaID string) (*message.Result, error) {
	msg := message.NewVoice(agenid, toAll, false, user, party, tag, mediaID)
	return message.SendMessage(wx.client, *msg)
}

//SendSafeVoice 发送保密语音消息
func (wx *Wx) SendSafeVoice(agenid int64, toAll, issafe bool, user, party, tag []string, mediaID string) (*message.Result, error) {
	msg := message.NewVoice(agenid, toAll, issafe, user, party, tag, mediaID)
	return message.SendMessage(wx.client, *msg)
}

//SendVoiceToAll 发送语音消息给所有人
func (wx *Wx) SendVoiceToAll(agenid int64, mediaID string) (*message.Result, error) {
	msg := message.NewVoice(agenid, true, false, nil, nil, nil, mediaID)
	return message.SendMessage(wx.client, *msg)
}

//SendSafeVoiceToAll 发送保密语音消息给所有人
func (wx *Wx) SendSafeVoiceToAll(agenid int64, mediaID string) (*message.Result, error) {
	msg := message.NewVoice(agenid, true, true, nil, nil, nil, mediaID)
	return message.SendMessage(wx.client, *msg)
}

//SendVoiceToUser 发送语音消息给用户
func (wx *Wx) SendVoiceToUser(agenid int64, mediaID string, user ...string) (*message.Result, error) {
	msg := message.NewVoice(agenid, false, false, user, nil, nil, mediaID)
	return message.SendMessage(wx.client, *msg)
}

//SendSafeVoiceToUser 发送保密语音消息给用户
func (wx *Wx) SendSafeVoiceToUser(agenid int64, mediaID string, user ...string) (*message.Result, error) {
	msg := message.NewVoice(agenid, false, true, user, nil, nil, mediaID)
	return message.SendMessage(wx.client, *msg)
}

//SendVoiceToParty 发送语音消息给部门
func (wx *Wx) SendVoiceToParty(agenid int64, mediaID string, party ...string) (*message.Result, error) {
	msg := message.NewVoice(agenid, false, false, nil, party, nil, mediaID)
	return message.SendMessage(wx.client, *msg)
}

//SendSafeVoiceToParty 发送保密语音消息给部门
func (wx *Wx) SendSafeVoiceToParty(agenid int64, mediaID string, party ...string) (*message.Result, error) {
	msg := message.NewVoice(agenid, false, true, nil, party, nil, mediaID)
	return message.SendMessage(wx.client, *msg)
}

//SendVoiceToTags 发送语音消息给用户
func (wx *Wx) SendVoiceToTags(agenid int64, mediaID string, tag ...string) (*message.Result, error) {
	msg := message.NewVoice(agenid, false, false, nil, nil, tag, mediaID)
	return message.SendMessage(wx.client, *msg)
}

//SendSafeVoiceToTags 发送保密语音消息给用户组
func (wx *Wx) SendSafeVoiceToTags(agenid int64, mediaID string, tag ...string) (*message.Result, error) {
	msg := message.NewVoice(agenid, false, true, nil, nil, tag, mediaID)
	return message.SendMessage(wx.client, *msg)
}

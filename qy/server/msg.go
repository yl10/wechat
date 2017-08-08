package server

/*
消息
*/

import (
	. "github.com/yl10/wechat/util"
)

const (
	ActionTextMsg     HanderAction = "text"
	ActionVoiceMsg                 = "voice"
	ActionVideoMsg                 = "video"
	ActionLocationMsg              = "location"
	ActionLinkMsg                  = "link"
)

type TextMsg struct {
	RequestMsgHeader
	MsgID   int `xml:"MsgId"`
	Content CDATA
}

func (o TextMsg) getheader() RequestMsgHeader {
	return o.RequestMsgHeader
}

type VoiceMsg struct {
	RequestMsgHeader
	MsgID   int `xml:"MsgId"`
	Format  CDATA
	MediaID CDATA `xml:"MediaId"`
}

func (o VoiceMsg) getheader() RequestMsgHeader {
	return o.RequestMsgHeader
}

type VideoMsg struct {
	RequestMsgHeader
	MsgID        int   `xml:"MsgId"`
	MediaID      CDATA `xml:"MediaId"`
	ThumbMediaID CDATA `xml:"ThumbMediaId"`
}

func (o VideoMsg) getheader() RequestMsgHeader {
	return o.RequestMsgHeader
}

type LocationMsg struct {
	RequestMsgHeader
	LocationX float64 `xml:"Location_X"`
	LocationY float64 `xml:"Location_Y"`
	Scale     int
	Label     CDATA
	MsgID     int `xml:"MsgId"`
}

func (o LocationMsg) getheader() RequestMsgHeader {
	return o.RequestMsgHeader
}

type LinkMsg struct {
	RequestMsgHeader
	Title       CDATA
	Description CDATA
	PicURL      CDATA `xml:"PicUrl"`

	MsgID int `xml:"MsgId"`
}

func (o LinkMsg) getheader() RequestMsgHeader {
	return o.RequestMsgHeader
}

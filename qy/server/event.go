package server

/*
除菜单和通讯录以外的event
*/

import (
	. "github.com/yl10/wechat/util"
)

const (
	//ActionSubscribe 订阅
	ActionSubscribe HanderAction = "subscribe"
	//ActionUnSubscribe 取消订阅
	ActionUnSubscribe = "sunubscribe"

	//ActionEnterAgent 进入应用
	ActionEnterAgent = "enter_agent"
	//ActionLocation 上报地理位置
	ActionLocation = "LOCATION"

	//ActionSyncUser 异步任务完成事件推送 增量更新成员
	ActionSyncUser = "sync_user"
	//ActionReplaceUser (全量覆盖成员
	ActionReplaceUser = "replace_user"
	//ActionInviteUser 邀请成员关注
	ActionInviteUser = "invite_user"
	//ActionReplaceParty 全量覆盖部门
	ActionReplaceParty = "replace_party"
	//ActionClick 点击点击菜单拉取消息的事件推送
	ActionClick = "click"
	//ActionView 点击菜单跳转链接的事件推送
	ActionView = "view"

	//ActionScancodePush 扫码推事件的事件推送
	ActionScancodePush = "scancode_push"
	//ActionScancodeWaitMsg 扫码推事件且弹出“消息接收中”提示框的事件推送
	ActionScancodeWaitMsg = "scancode_waitmsg"

	//ActionPicSysphoto 弹出系统拍照发图的事件推送
	ActionPicSysphoto = "pic_sysphoto"
	//ActionPicPhotoOrAlbum 弹出拍照或者相册发图的事件推送
	ActionPicPhotoOrAlbum = "pic_photo_or_album"

	//ActionPicWeixin 弹出微信相册发图器的事件推送
	ActionPicWeixin = "pic_weixin"

	//ActionLocationSelect 弹出地理位置选择器的事件推送
	ActionLocationSelect = "location_select"
)

//SubscribeRequestMsg 订阅
/*
成员关注/取消关注事件
本事件在成员关注（取消关注）应用时触发
消息类型，此时固定为：event
事件类型，subscribe(订阅)、unsubscribe(取消订阅)
事件KEY值，此事件该值为空
*/
type SubscribeRequestMsg struct {
	RequestMsgHeader
}

func (o SubscribeRequestMsg) getheader() RequestMsgHeader {
	return o.RequestMsgHeader
}

/*EnterAgentRequestMsg 进入应用
消息类型，此时固定为：event
事件类型：enter_agent
事件KEY值，此事件该值为空

*/
type EnterAgentRequestMsg struct {
	RequestMsgHeader
}

func (o EnterAgentRequestMsg) getheader() RequestMsgHeader {
	return o.RequestMsgHeader
}

//LocationRequestMsg 上报地理位置
/*

成员同意上报地理位置后，每次在进入应用会话时都会上报一次地理位置，或在进入应用会话后每5秒上报一次地理位置。
企业可以在管理端修改应用是否需要获取地理位置权限。
消息类型，此时固定为：event
事件类型：LOCATION
*/
type LocationRequestMsg struct {
	RequestMsgHeader
	Latitude  float64 //地理位置纬度
	Longitude float64 //地理位置经度
	Precision float64 //地理位置精度
}

func (o LocationRequestMsg) getheader() RequestMsgHeader {
	return o.RequestMsgHeader
}

//BatchJobResultRequestMsg 异步任务完成事件推送
/*


本事件是成员在使用异步任务接口时，用于接收任务执行完毕的结果通知。

消息类型，此时固定为：event
事件类型：batch_job_result

*/
type BatchJobResultRequestMsg struct {
	RequestMsgHeader
	JobID   CDATA `xml:"BatchJob>JobId"`   //异步任务id，最大长度为64字符
	JobType CDATA `xml:"BatchJob>JobType"` //操作类型，字符串，目前分别有：sync_user(增量更新成员)、 replace_user(全量覆盖成员）、invite_user(邀请成员关注）、replace_party(全量覆盖部门)
	ErrCode CDATA `xml:"BatchJob>ErrCode"` //
	ErrMsg  CDATA `xml:"BatchJob>ErrMsg"`  //对返回码的文本描述内容
}

func (o BatchJobResultRequestMsg) getheader() RequestMsgHeader {
	return o.RequestMsgHeader
}

//PicWeixinRequestMsg 弹出微信相册发图器的事件推送
type PicWeixinRequestMsg struct {
	RequestMsgHeader
	EventKey     CDATA //`xml:",cdata"` //	事件KEY值，与自定义菜单接口中KEY值对应
	SendPicsInfo struct {
		Count   int       //发送的图片数量
		PicList []PicItem `xml:"PicList"` //发送的图片信息
	}
	AgentID CDATA //企业应用的id，整型。可在应用的设置页面查看
}

func (o PicWeixinRequestMsg) getheader() RequestMsgHeader {
	return o.RequestMsgHeader
}

//PicItem 图像iteam
type PicItem struct {
	PicMd5Sum CDATA `xml:"item>PicMd5Sum"` //图片的MD5值，开发者若需要，可用于验证接收到图片
}

//LocationSelectRequestMsg 弹出地理位置选择器的事件推送
//消息类型，此时固定为：event
//事件类型：location_select
type LocationSelectRequestMsg struct {
	RequestMsgHeader
	SendLocationInfo LocationInfo //发送的位置信息

}

func (o LocationSelectRequestMsg) getheader() RequestMsgHeader {
	return o.RequestMsgHeader
}

//LocationInfo 位置信息
type LocationInfo struct {
	X       CDATA `xml:"Location_X"` ////	X坐标信息
	Y       CDATA `xml:"Location_Y"` //	Y坐标信息
	Scale   CDATA //精度，可理解为精度或者比例尺、越精细的话 scale越高
	Label   CDATA //	地理位置的字符串信息
	Poiname CDATA //	POI的名字，可能为空
}

//MenuRequestMsg 菜单点击事件
type MenuRequestMsg struct {
	RequestMsgHeader
}

func (o MenuRequestMsg) getheader() RequestMsgHeader {
	return o.RequestMsgHeader
}

//ScancodeRequestMsg 扫码推事件的事件推送
type ScancodeRequestMsg struct {
	RequestMsgHeader
	ScanType   CDATA `XML:"ScanCodeInfo>ScanType"`
	ScanResult CDATA `XML:"ScanCodeInfo>ScanResult"`
}

func (o ScancodeRequestMsg) getheader() RequestMsgHeader {
	return o.RequestMsgHeader
}

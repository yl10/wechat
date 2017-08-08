package server

import (
	"encoding/xml"
	"fmt"

	. "github.com/yl10/wechat/util"
)

//HanderAction 最终动作定义
type HanderAction string

func (h HanderAction) String() string {
	return string(h)
}

//RequestMsg RequestMsg 接口
type RequestMsg interface {
	getheader() RequestMsgHeader
}

//RequestMsgHeader 信息公共字段
type RequestMsgHeader struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATA    //	企业微信CorpID
	FromUserName CDATA    //	成员UserID
	CreateTime   int64    //消息创建时间（整型）
	MsgType      CDATA    //消息类型，此时固定为：event
	Event        CDATA    //事件类型，subscribe(订阅)、unsubscribe(取消订阅)
	EventKey     CDATA    //事件KEY值
	AgentID      int64    //	企业应用的id，整型。可在应用的设置页面查看
}

//MsgHandleFunc 对消息进行处理的函数
type MsgHandleFunc func(requestbody interface{}) error

/*PareXMLMsg 解析明文消息
 *@param xmldata []byte 类型的明文消息
 *@param @onlyHeader 是否只检查消息头部
 */
func PareXMLMsg(xmldata []byte, onlyHeader bool) (action HanderAction, header RequestMsgHeader, result RequestMsg, err error) {
	//解析消息头部

	var param struct {
		RequestMsgHeader
		ChangeType string //通讯录的时候用得上
		JobType    CDATA  //
	}

	err = xml.Unmarshal(xmldata, &param)
	if err != nil {
		err = fmt.Errorf("请求解析失败：%v", err)
		return
	}
	header = param.RequestMsgHeader
	if onlyHeader {

		return
	}
	msgtype := param.MsgType.String()
	eventtype := HanderAction(param.Event.String())
	jobtype := param.JobType.String()
	switch msgtype {
	case "event": //事件
		switch eventtype {
		case "change_contact": //通讯录事件
			action = HanderAction(param.ChangeType)
			switch action {
			case ActionCreateUser, ActionUpdateUser, ActionDeleteUser: //新增成员,update_user 更新成员 delete_user 删除成员
				result = &UserRequestMsg{}
			case ActionCreateParty, ActionUpdateParty, ActionDeleteParty: //create_party 新增部 //update_party 更新部门//delete_party 删除部门
				result = &PartyRequestMsg{}
			case ActionUpdateTag: //update_tag 变更标签
				result = &TagRequestMsg{}
			default:
				err = fmt.Errorf("尚未定义的通讯录事件：%s", eventtype)
				return
			}
		case ActionClick, ActionView: //点击菜单拉取消息，点才菜单跳转
			action = eventtype
			result = &MenuRequestMsg{}

		case ActionScancodePush, ActionScancodeWaitMsg: //扫描推送 //scancode_waitmsg 扫码推事件且弹出“消息接收中”提示框的事件推送
			action = eventtype
			result = &ScancodeRequestMsg{}

		case ActionPicPhotoOrAlbum, ActionPicSysphoto, ActionPicWeixin: //pic_sysphoto 弹出系统拍照发图的事件推送 //pic_photo_or_album 弹出拍照或者相册发图的事件推送 //pic_weixin 弹出微信相册发图器的事件推送
			action = eventtype
			result = &PicWeixinRequestMsg{}

		case ActionLocationSelect: //location_select 弹出地理位置选择器的事件推送
			action = eventtype
			result = &LocationSelectRequestMsg{}

		case ActionSubscribe, ActionUnSubscribe: //成员关注
			action = eventtype
			result = &SubscribeRequestMsg{}

		case ActionEnterAgent: //事件类型：enter_agent
			action = ActionEnterAgent
			result = &EnterAgentRequestMsg{}
		case ActionLocation: //上报地理应用
			action = ActionLocation
			result = &LocationRequestMsg{}
		case "batch_job_result": //batch_job_result
			action = HanderAction(jobtype)
			result = &BatchJobResultRequestMsg{}
		}

	case "text":
		action = ActionTextMsg
		result = &TextMsg{}
	case "voice":
		action = ActionVoiceMsg
		result = &VoiceMsg{}
	case "video":
		action = ActionVideoMsg
		result = &VideoMsg{}
	case "location":
		action = ActionLocationMsg
		result = &LocationMsg{}
	case "link":
		action = ActionLinkMsg
		result = &LinkMsg{}
	default:
		err = fmt.Errorf("未定义的MsgType:%s", msgtype)
		return
	}
	err = xml.Unmarshal(xmldata, &result)
	if err != nil {
		err = fmt.Errorf("二次解析失败：%v", err)
		return
	}

	return
}

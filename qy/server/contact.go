package server

import (
	. "github.com/yl10/wechat/util"
)

const (
	ActionCreateUser  HanderAction = "create_user" //CreateUser df
	ActionUpdateUser               = "update_user"
	ActionDeleteUser               = "delete_user"
	ActionCreateParty              = "create_party" //create_party 新增部
	ActionUpdateParty              = "update_party" //update_party 更新部门
	ActionDeleteParty              = "delete_party" //delete_party 删除部门
	ActionUpdateTag                = "update_tag"
)

//FnCreateUser 增加用户
type FnCreateUser func(u User) error

//FnUpdateUser 更新用户
type FnUpdateUser func(u User) error

//FnDeleteUser 删除用户
type FnDeleteUser func(userid string) error

/*
通讯录变更事件

当企业通过通讯录助手开通通讯录权限后，成员的变更会通知给企业。变更的事件，将推送到企业微信管理端通讯录助手中的‘接收事件服务器’。由通讯录同步助手调用接口触发的变更事件不回调通讯录同步助手本身。企业微信的成员在客户端变更自己的个人信息将推送给通讯录同步助手。第三方通讯录变更事件参见第三方回调协议

*/

//UserRequestMsg 通讯录变更事件
type UserRequestMsg struct {
	RequestMsgHeader
	ChangeType string //create_user|update_user|delete_user
	User
}

func (o UserRequestMsg) getheader() RequestMsgHeader {
	return o.RequestMsgHeader
}

//User 用户信息
type User struct {
	UserID      CDATA      //	成员UserID
	Name        CDATA      //成员名称
	Department  CDATA      //成员部门列表
	Mobile      CDATA      //手机号码
	Position    CDATA      //职位信息。长度为0~64个字节
	Gender      int        //	性别，1表示男性，2表示女性
	Email       CDATA      //邮箱
	Status      int        //激活状态：1=已激活 2=已禁用
	Avatar      CDATA      //头像
	EnglishName CDATA      //英文名
	IsLeader    int        //上级字段，标识是否为上级。0表示普通成员，1表示上级
	Telephone   CDATA      //座机
	ExtAttr     []UserAttr //扩展属性
}

//UserAttr 用户扩展属性
type UserAttr struct {
	Name  CDATA
	Value CDATA
}

//PartyRequestMsg 部门事件请求
type PartyRequestMsg struct {
	RequestMsgHeader
	ChangeType CDATA //
	Party
}

func (o PartyRequestMsg) getheader() RequestMsgHeader {
	return o.RequestMsgHeader
}

//TagRequestMsg 标签请求
type TagRequestMsg struct {
	RequestMsgHeader
	ChangeType    string
	TagID         int   `xml:"TagId"` //标签Id
	AddUserItems  CDATA // 标签中新增的成员userid列表，用逗号分隔
	DelUserItems  CDATA //标签中删除的成员userid列表，用逗号分隔
	AddPartyItems CDATA //标签中新增的部门id列表，用逗号分隔
	DelPartyItems CDATA //标签中删除的部门id列表，用逗号分隔
}

func (o TagRequestMsg) getheader() RequestMsgHeader {
	return o.RequestMsgHeader
}

//Party 部门
type Party struct {
	ID       int
	Name     CDATA
	ParentID CDATA `xml:"ParentId"`
	Order    int
}

package qy

import (
	"github.com/yl10/wechat/qy/addr"
)

//GetUserInfoByUserID 通过userID获取用户信息
func (wx *Wx) GetUserInfoByUserID(userID string) (addr.User, error) {
	return addr.GetUserinfo(wx.client, userID)
}

//GetDepartmentlist 获取部门列表
func (wx *Wx) GetDepartmentlist(id ...int) ([]addr.Department, error) {
	return addr.GetDepartmentlist(wx.client, id...)
}

//CreateDepartment 创建部门，传入部门名称，上级部门id，顺序号，返回部门id和错误
func (wx *Wx) CreateDepartment(name, parentid, order string) (id int, err error) {
	return addr.CreateDepartment(wx.client, name, parentid, order)
}

//UpdateDepartment 更新部门
func (wx *Wx) UpdateDepartment(id int, name, parentid, order string) error {
	return addr.UpdateDepartment(wx.client, id, name, parentid, order)
}

//DeleteDepartment 根据部门ID删除部门
func (wx *Wx) DeleteDepartment(id int) error {
	return addr.DeleteDepartment(wx.client, id)
}

//CreateUser 创建用户
func (wx *Wx) CreateUser(user addr.User) error {
	return addr.CreateUser(wx.client, user)
}

//CreateUserFull 创建用户，全信息
func (wx *Wx) CreateUserFull(userid, name string, department []int, position, mobile, email, weixinid string, attrs []addr.UserAttr) error {
	return addr.CreateUserFull(wx.client, userid, name, department, position, mobile, email, weixinid, attrs)
}

//DeleteUser 删除单个用户
func (wx *Wx) DeleteUser(userid string) error {
	return addr.DeleteUser(wx.client, userid)
}

//BatchDeleteUser 批量删除用户
func (wx *Wx) BatchDeleteUser(useridlist []string) error {
	return addr.BatchDeleteUser(wx.client, useridlist)
}

//GetUserListByDept 获取部门成员
//deptID 部门ID
//details 是否获取详情，非详情只有userid和name
//fetch_child 是否递归，不传就是false
func (wx *Wx) GetUserListByDept(deptID string, details bool, fetchChild ...bool) ([]addr.User, error) {
	return addr.GetUserListByDept(wx.client, deptID, details, fetchChild...)
}

//CreateTag 创建标签
//name 标签名称
//id 标签id
func (wx *Wx) CreateTag(name string, id ...int) (int, error) {
	return addr.CreateTag(wx.client, name, id...)
}

//UpdateTag 更新tag
func (wx *Wx) UpdateTag(name string, id int) error {
	return addr.UpdateTag(wx.client, name, id)
}

//DeleteTag 删除标签
func (wx *Wx) DeleteTag(tagid int) error {
	return addr.DeleteTag(wx.client, tagid)
}

/*GetUsersByTag 获取标签下成员
userlist:map[userid]username
*/
func (wx *Wx) GetUsersByTag(tagid int) (userlist map[string]string, partylist []int, err error) {
	return addr.GetUsersByTag(wx.client, tagid)
}

//AddUserToTag 给标签添加成员
func (wx *Wx) AddUserToTag(tagid int, userlist []string, partylist []int) (invalidlist []string, invalidparty []int, err error) {
	return addr.AddUserToTag(wx.client, tagid, userlist, partylist)
}

//DeleteUserFromTag 从标签中移除成员
func (wx *Wx) DeleteUserFromTag(tagid int, userlist []string, partylist []int) (invalidlist []string, invalidparty []int, err error) {
	return addr.DeleteUserFromTag(wx.client, tagid, userlist, partylist)
}


package addr

import (
	"encoding/json"
	"fmt"

	"github.com/yl10/wechat/qy/client"
	"github.com/yl10/wechat/qy/model"
)

//Department 部门
type Department struct {
	ID       int `json:"id"`
	Name     string
	ParentID int `json:"parentid"`
	Order    int
	Child    []Department
}

//User 用户
type User struct {
	UserID        string `json:"userid"`
	Name          string
	Department    []int
	Position      string
	Mobile        string
	Email         string
	Gender        string
	WeixinID      string `json:"weixinid"`
	Enable        bool
	AvatarMediaID string `json:"avatar_mediaid"`
	Status        int

	Extattr struct {
		Attrs []UserAttr
	}
}

//UserAttr 用户属性
type UserAttr struct {
	Name  string
	value string
}

//CreateDepartment 创建部门，传入部门名称，上级部门id，顺序号，返回部门id和错误
func CreateDepartment(c *client.Client, name, parentid, order string) (id int, err error) {

	var p struct {
		Name     string
		Parentid string
		Order    string
	}
	p.Name = name
	p.Parentid = parentid
	p.Order = order

	data, err := c.PostJSON("https://qyapi.weixin.qq.com/cgi-bin/department/create?access_token=%s", p)
	if err != nil {
		return 0, err
	}

	var v struct {
		ID int `josn:"Id"`
	}
	err = json.Unmarshal(data, &v)
	if err != nil {
		return 0, err
	}
	return v.ID, nil
}

//UpdateDepartment 更新部门
func UpdateDepartment(c *client.Client, id int, name, parentid, order string) error {
	var p struct {
		ID       int `json:"id"`
		Name     string
		Parentid string
		Order    string
	}
	p.ID = id
	p.Name = name
	p.Parentid = parentid
	p.Order = order

	data, err := c.PostJSON("https://qyapi.weixin.qq.com/cgi-bin/department/update?access_token=%s", p)
	if err != nil {
		return err
	}

	var v struct {
		ID int `json:"id"`
	}
	err = json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	return nil
}

//DeleteDepartment 根据部门ID删除部门
func DeleteDepartment(c *client.Client, id int) error {
	reqstr := "https://qyapi.weixin.qq.com/cgi-bin/department/list?access_token=%s&id=" + fmt.Sprint(id)
	_, err := c.SendGetRequest(reqstr)
	return err
}

//GetDepartmentlist 获取部门列表
func GetDepartmentlist(c *client.Client) ([]Department, error) {
	data, err := c.SendGetRequest("https://qyapi.weixin.qq.com/cgi-bin/department/list?access_token=%s")
	var ds struct {
		Department []Department
	}
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &ds)
	if err != nil {
		return nil, err
	}
	return ds.Department, nil

}

//CreateUser 创建用户
func CreateUser(c *client.Client, user model.User) error {
	return c.CreateUserFull(user.Userid, user.Name, user.Department, user.Position, user.Mobile, user.Email, user.Weixinid, user.Extattr.Attrs)
}

//CreateUserFull 创建用户，全信息
func CreateUserFull(c *client.Client, userid, name string, department []int, position, mobile, email, weixinid string, attrs []UserAttr) error {
	var user struct {
		Userid     string
		Name       string
		Department []int
		Position   string
		Mobile     string
		Email      string
		Weixinid   string
		Extattr    struct {
			Attrs []UserAttr
		}
	}
	user.Userid = userid
	user.Name = name
	user.Department = department
	user.Position = position
	user.Mobile = mobile
	user.Email = email
	user.Weixinid = weixinid
	user.Extattr.Attrs = attrs

	_, err := c.PostJSON("https://qyapi.weixin.qq.com/cgi-bin/user/create?access_token=%s", user)
	return err
}

//DeleteUser 删除单个用户
func DeleteUser(c *client.Client, userid string) error {

	urlstr := "https://qyapi.weixin.qq.com/cgi-bin/user/delete?access_token=%s&userid=" + userid
	_, err := c.SendGetRequest(urlstr)
	return err
}

//BatchDeleteUser 批量删除用户
func BatchDeleteUser(c *client.Client, useridlist []string) error {
	var list struct {
		Useridlist []string
	}
	list.Useridlist = useridlist
	_, err := c.PostJSON("https://qyapi.weixin.qq.com/cgi-bin/user/batchdelete?access_token=%s", list)
	return err
}

//GetUserinfo 获取用户信息
func GetUserinfo(c *client.Client, userid string) (User, error) {
	var user User
	data, err := c.SendGetRequest("https://qyapi.weixin.qq.com/cgi-bin/user/get?access_token=%s&userid=" + userid)
	if err != nil {
		return user, err
	}
	err = json.Unmarshal(data, &user)

	return user, err
}

//获取部门用户

//更新用户信息，还没写
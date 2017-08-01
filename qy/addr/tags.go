package addr

import (
	"encoding/json"
	"fmt"

	"github.com/yl10/wechat/qy/client"
)

const (
	//ADD 增加
	ADD Action = "add"
	//DELETE 删除
	DELETE = "delete"
)

//Action 动作
type Action string

//Tag 标签
type Tag struct {
	Name string `json:"tagname"`
	ID   int    `json:"tagid"`
}

//CreateTag 创建标签
//name 标签名称
//id 标签id
func CreateTag(c *client.Client, name string, id ...int) (int, error) {
	if len(name) >= 32 {
		return 0, fmt.Errorf("更新标签失败：%s", "标签的长度不能超过32")
	}
	uri := "https://qyapi.weixin.qq.com/cgi-bin/tag/create?"
	tag := Tag{Name: name}
	if len(id) > 0 {
		tag.ID = id[0]
	}
	resdata, err := c.PostJSON(uri, tag)
	if err != nil {
		return 0, err
	}
	var result struct {
		Tagid int
	}

	err = json.Unmarshal(resdata, &result)
	return result.Tagid, err
}

//UpdateTag 更新tag
func UpdateTag(c *client.Client, name string, id int) error {
	uri := "https://qyapi.weixin.qq.com/cgi-bin/tag/update?"
	tag := Tag{Name: name, ID: id}

	if len(name) >= 32 {
		return fmt.Errorf("更新标签失败：%s", "标签的长度不能超过32")
	}
	_, err := c.PostJSON(uri, tag)
	return err

}

//DeleteTag 删除标签
func DeleteTag(c *client.Client, tagid int) error {
	uri := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/tag/delete?tagid=%d", tagid)

	_, err := c.SendGetRequest(uri)
	return err
}

/*GetUsersByTag 获取标签下成员
userlist:map[userid]username
*/
func GetUsersByTag(c *client.Client, tagid int) (userlist map[string]string, partylist []int, err error) {
	uri := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/tag/get?tagid=%d", tagid)
	resdata, err := c.SendGetRequest(uri)

	var result struct {
		Userlist []struct {
			Userid string
			Name   string
		}
		Partylist []int
	}

	err = json.Unmarshal(resdata, &result)
	if err != nil {
		return nil, nil, err
	}

	for _, v := range result.Userlist {
		userlist[v.Userid] = v.Name
	}
	partylist = append(partylist, result.Partylist...)
	return

}

//AddUserToTag 给标签添加成员
func AddUserToTag(c *client.Client, tagid int, userlist []string, partylist []int) (invalidlist []string, invalidparty []int, err error) {

	return patchuserbytag(c, ADD, tagid, userlist, partylist)
}

//DeleteUserFromTag 从标签中移除成员
func DeleteUserFromTag(c *client.Client, tagid int, userlist []string, partylist []int) (invalidlist []string, invalidparty []int, err error) {
	return patchuserbytag(c, DELETE, tagid, userlist, partylist)
}

func patchuserbytag(c *client.Client, action Action, tagid int, userlist []string, partylist []int) (invalidlist []string, invalidparty []int, err error) {

	if len(userlist) == 0 && len(partylist) == 0 {
		return nil, nil, fmt.Errorf("标签操作失败：动作(%s),成员和组织不能同时为控", action)
	}

	uri := ""
	switch action {
	case ADD:
		uri = "https://qyapi.weixin.qq.com/cgi-bin/tag/addtagusers?"
	case DELETE:
		uri = "https://qyapi.weixin.qq.com/cgi-bin/tag/deltagusers?"
	default:
		return nil, nil, fmt.Errorf("标签操作失败：未定义的动作类型，请使用ADD|DELETE")
	}
	resdata, err := c.PostJSON(uri, struct {
		Tagid     int
		Userlist  []string
		Partylist []int
	}{Tagid: tagid, Userlist: userlist, Partylist: partylist})
	if err != nil {
		return nil, nil, err
	}

	var invalid struct {
		Invalidlist  []string
		Invalidparty []int
	}

	err = json.Unmarshal(resdata, &invalid)
	return invalid.Invalidlist, invalid.Invalidparty, err

}

//GetTagList 获取标签列表
func GetTagList(c *client.Client) (map[int]string, error) {
	uri := "https://qyapi.weixin.qq.com/cgi-bin/tag/list?"

	var result struct {
		Taglist []Tag
	}

	resdata, err := c.SendGetRequest(uri)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resdata, &result)
	if err != nil {
		return nil, fmt.Errorf("解析json失败：%v", err)
	}
	m := make(map[int]string, 0)
	for _, v := range result.Taglist {
		m[v.ID] = v.Name
	}
	return m, nil

}

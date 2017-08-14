package agent

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yl10/wechat/qy/client"
)

const (
	//ERRNoClient 没有tokenserver，请使用NewAgent进行定义
	ERRNoClient = "没有tokenserver，请使用NewAgent进行定义。"
)

//Agent 应用
type Agent struct {
	corpID string
	secret string

	ID   string `json:"agentid"`
	Name string `json:"name"`
	//企业应用方形头像
	SquareLogoURL string `json:"square_logo_url"`
	//企业应用详情
	Description string `json:"description"`

	Users          []string `json:","`
	Partys         []int    `json:","`
	Tags           []int    `json:","`
	Closed         bool     `json:"close"`
	RedirectDomain string   `json:"redirect_domain"`
	ReportLocation bool     `json:"report_location_flag"`
	ReportEnter    bool     `json:"isreportenter"`
	HomeURL        string   `json:"home_url"`

	client *client.Client
}

//UnmarshalJSON 实现Unmarshaler接口
func (a *Agent) UnmarshalJSON(data []byte) (err error) {

	agent := make(map[string]interface{})
	err = json.Unmarshal(data, &agent)
	if err != nil {
		return
	}
	defer func() {
		if err2 := recover(); err2 != nil {
			err = fmt.Errorf("%v", err2)
		}
	}()
	a.Name = agent["name"].(string)
	a.SquareLogoURL = agent["square_logo_url"].(string)
	a.Description = agent["description"].(string)
	a.RedirectDomain = agent["redirect_domain"].(string)
	a.HomeURL = agent["home_url"].(string)
	a.Closed = agent["close"].(float64) == 1
	a.ReportLocation = agent["report_location_flag"].(float64) == 1
	a.ReportEnter = agent["isreportenter"].(float64) == 1

	fmt.Println(agent)

	userinfo, _ := agent["allow_userinfos"].(map[string][]string)

	a.Users = userinfo["user"]
	if agent["allow_userinfos"] != nil {

		usermap := agent["allow_userinfos"].(map[string]interface{})["user"].([]interface{})
		for _, v := range usermap {
			a.Users = append(a.Users, v.(map[string]interface{})["userid"].(string))
		}

	}

	if agent["allow_partys"] != nil {
		partymap := agent["allow_partys"].(map[string]interface{})["partyid"].([]interface{})
		a.Partys = make([]int, 0)
		for _, v := range partymap {
			fmt.Printf("shzi:%T     %v", v, v)
			pid := int(v.(float64))
			fmt.Printf("pid: %d ", pid)
			a.Partys = append(a.Partys, pid)

		}
	}

	if agent["allow_tags"] != nil {
		tagmap := agent["allow_tags"].(map[string]interface{})["tagid"].([]interface{})
		a.Tags = make([]int, 0)
		for _, v := range tagmap {
			fmt.Printf("shzi:%T     %v", v, v)
			tid := int(v.(float64))

			a.Tags = append(a.Tags, tid)

		}
	}

	// 	type at struct {
	// 		Agent
	// 	}
	// 	var anget Agent
	// 	// var anget struct {
	// 	// 	Agent
	// 	// 	//AllowUserinfos map[string][]map[string]string `json:"allow_userinfos"`
	// 	// 	// AllowUserinfos struct {
	// 	// 	// 	User []struct {
	// 	// 	// 		Userid string
	// 	// 	// 	}
	// 	// 	// }
	// 	// 	// AllowPartys struct {
	// 	// 	// 	Partyid []int
	// 	// 	// } `json:"allow_partys"`
	// 	// 	// AllowTags struct {
	// 	// 	// 	Tagid []int
	// 	// 	// } `json:"allow_tags"`
	// 	// }

	// 	if err := json.Unmarshal(data, &anget); err != nil {
	// 		return err
	// 	}

	// 	a.ID = anget.ID
	// 	a.Name = anget.Name
	// 	a.SquareLogoURL = anget.SquareLogoURL
	// 	a.Description = anget.Description
	// 	a.Closed = anget.Closed
	// 	a.RedirectDomain = anget.RedirectDomain
	// 	a.ReportEnter = anget.ReportEnter
	// 	a.ReportLocation = anget.ReportLocation
	// 	a.HomeURL = anget.HomeURL

	// 	// a.Users = make([]string, 0)
	// 	// for _, v := range anget.AllowUserinfos.User {
	// 	// 	a.Users = append(a.Users, v.Userid)
	// 	// }

	// 	// a.Partys = make([]int, 0)

	// 	// a.Partys = append(a.Partys, anget.AllowPartys.Partyid...)

	// 	// a.Tags = make([]int, 0)
	// 	// a.Tags = append(a.Tags, anget.AllowTags.Tagid...)
	return nil
}

//GetInfoFromTencent 从微信服务器获取应用向前；第三方仅可获取被授权的应用。
func (a *Agent) GetInfoFromTencent() error {
	strurl := "https://qyapi.weixin.qq.com/cgi-bin/agent/get?access_token=%s&agentid=" + a.ID
	if a.client == nil {
		return fmt.Errorf(ERRNoClient)
	}
	result, err := a.client.SendGetRequest(strurl)
	if err != nil {
		return err
	}
	return json.Unmarshal(result, a)
}

//NewAgent 创建一个应用
func NewAgent(id string, corpid string, secret string, srv client.AccesstokenServer, htc *http.Client) (*Agent, error) {
	c, err := client.NewClient(corpid, secret, srv, htc)
	if err != nil {
		return nil, err
	}
	agent := &Agent{
		ID:     id,
		corpID: corpid,
		secret: secret,
		client: c,
	}
	return agent, err

}

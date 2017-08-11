package server

import (
	"encoding/json"

	"github.com/yl10/wechat/qy/client"
)

//GetCallbackIP 获取企业微信服务器的ip段
func GetCallbackIP(c *client.Client) ([]string, error) {
	data, err := c.SendGetRequest("https://qyapi.weixin.qq.com/cgi-bin/getcallbackip?access_token=%s")
	if err != nil {
		return nil, err
	}

	var result struct {
		IpList []string `json:"ip_list"`
	}
	err = json.Unmarshal(data, &result)
	return result.IpList, err
}

package oauth

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/yl10/tencent/wechat/qy/client"
)

//User 授权用户
type User struct {
	IsQy     bool
	UserID   string `json:"UserId"`
	OpenID   string `json:"OpenId"`
	DeviceID string `json:"DeviceId"`
}

//GetOAuthURI 构造获取用户信息用的URL
//appid:企业的CorpID
//redirectURI:授权后重定向的回调链接地址，无需使用urlencode对链接进行处理
//state:重定向后会带上state参数，企业可以填写a-zA-Z0-9的参数值，长度不可超过128个字节
func GetOAuthURI(appid, redirectURI string, state ...string) string {
	str := ""
	redirectURI = url.QueryEscape(redirectURI)

	if len(state) > 0 {

		str = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=%s&state=%s#wechat_redirect"

		return fmt.Sprintf(str, appid, redirectURI, "snsapi_base", state[0])

	}

	str = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=%s#wechat_redirect"

	return fmt.Sprintf(str, appid, redirectURI, "snsapi_base")

}

//GetUserInfo 还
func GetUserInfo(c *client.Client, code string) (*User, error) {

	uri := "https://qyapi.weixin.qq.com/cgi-bin/user/getuserinfo?access_token=%s&code=" + code
	resp, err := c.SendGetRequest(uri)
	if err != nil {
		return nil, err
	}
	var user User
	err = json.Unmarshal(resp, &user)

	return &user, err

}

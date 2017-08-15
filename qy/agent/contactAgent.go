package agent

import (
	"github.com/yl10/wechat/qy/client"
	"github.com/yl10/wechat/qy/server"
)

//ContactAgent 通讯录应用
type ContactAgent struct {
	client *client.Client
	server *server.Server
}

// //NewContactAgent 创建一个通讯录应用
// func NewContactAgent(corpid, secret string, token, encodingAesKey string) (*ContactAgent, error) {
// return n
// }

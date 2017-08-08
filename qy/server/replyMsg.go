package server

import (
	"fmt"
	"net/http"
	"time"
)

func (s Server) ReplyText(req RequestMsg, w http.ResponseWriter, text string) error {
	reqheader := req.getheader()
	xmlstr := `<xml>
   <ToUserName><![CDATA[%s]]></ToUserName>
   <FromUserName><![CDATA[%s]]></FromUserName> 
   <CreateTime>%d</CreateTime>
   <MsgType><![CDATA[text]]></MsgType>
   <Content><![CDATA[%s]]></Content>
</xml>`

	timestamp := time.Now().UnixNano()
	encryMsg, err := s.encryptHelper.EncryptMsg(fmt.Sprintf(xmlstr, reqheader.FromUserName, timestamp, text), fmt.Sprintf("%d", timestamp), "")
	if err != nil {
		return err
	}
	w.Write([]byte(encryMsg))
	return nil
}

package server

import (
	"encoding/xml"
	"fmt"
	"testing"
)

func TestMsg(t *testing.T) {
	pweixin := PicWeixinRequestMsg{
	// ToUserName:   "toUser",
	// FromUserName: "fromuser",
	// CreateTime:   1408090816,
	// MsgType:      "event",
	// Event:        "pic_weixin",
	// EventKey:     "6",
	}
	pweixin.ToUserName = "touser"
	pweixin.FromUserName = "Fromuser"
	pweixin.SendPicsInfo.Count = 1
	pweixin.SendPicsInfo.PicList = []PicItem{
		PicItem{PicMd5Sum: "5a75aaca956d97be686719218f275c6b"},
	}

	data, err := xml.MarshalIndent(pweixin, "", "  ")
	fmt.Print(err)
	fmt.Printf("%s", data)

	xmlstr := `<xml><ToUserName><![CDATA[toUser]]></ToUserName>
<FromUserName><![CDATA[FromUser]]></FromUserName>
<CreateTime>1408090816</CreateTime>
<MsgType><![CDATA[event]]></MsgType>
<Event><![CDATA[pic_weixin]]></Event>
<EventKey><![CDATA[6]]></EventKey>
<SendPicsInfo><Count>1</Count>
<PicList><item><PicMd5Sum><![CDATA[5a75aaca956d97be686719218f275c6b]]></PicMd5Sum>
</item>
</PicList>
</SendPicsInfo>
<AgentID>1</AgentID>
</xml>`

	msg2 := PicWeixinRequestMsg{}
	xml.Unmarshal([]byte(xmlstr), &msg2)

	//fmt.Println(msg2)

}

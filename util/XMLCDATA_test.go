package util

import (
	"encoding/xml"
	"fmt"
	"testing"
)

type msg struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATA
	FromUserName CDATA
}

func TestCdata(t *testing.T) {
	m := msg{ToUserName: "a<>aaa", FromUserName: "bbb"}
	//data, err := xml.Marshal(m)
	data, err := xml.MarshalIndent(m, "", " ")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%s\r\n", data)

	m2 := msg{}
	xml.Unmarshal(data, &m2)
	fmt.Println(m2)
}

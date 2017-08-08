package util

import (
	"encoding/xml"
)

//CDATA CDATA类型，用于转XML的时候带CDATA
type CDATA string

//MarshalXML MarshalXML
func (c CDATA) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(struct {
		string `xml:",cdata"`
	}{string(c)}, start)
}

func (c CDATA) String() string {
	return string(c)
}

package message

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/yl10/wechat/qy/client"
)

const (
	//SendMesseageURL 主动发送消息使用的url
	SendMesseageURL = "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s"
)

const (
	//MsgTypeText 文本
	MsgTypeText = "text"
	//MsgTypeImage 图片
	MsgTypeImage = "image"
	//MsgTypeVoice 语音
	MsgTypeVoice = "voice"
	//MsgTypeVideo 视频
	MsgTypeVideo = "video"
	//MsgTypeFile 文件
	MsgTypeFile = "file"
	//MsgTypeNews news消息
	MsgTypeNews = "news"
	//MsgTypeMPNews mpnews消息 mpnews消息与news消息类似，不同的是图文消息内容存储在微信后台，并且支持保密选项。每个应用每天最多可以发送100次。
	MsgTypeMPNews = "mpnews"
)

var (
	//ErrUnexpectedMesseage 异常：未定义的消息
	ErrUnexpectedMesseage = func(o interface{}) error {
		return fmt.Errorf("传入了没有定义的消息：%v", o)
	}
)

//Message 消息
type Message struct {
	Header
	Content interface{}
}

//Contenter 消息体接口
type Contenter interface{}

//Header 消息头
type Header struct {
	ToUser  string `json:"touser,omitempty"`  // 非必须; 员工ID列表(消息接收者, 多个接收者用‘|’分隔, 最多支持1000个). 特殊情况: 指定为@all, 则向关注该企业应用的全部成员发送
	ToParty string `json:"toparty,omitempty"` // 非必须; 部门ID列表, 多个接收者用‘|’分隔, 最多支持100个. 当touser为@all时忽略本参数
	ToTag   string `json:"totag,omitempty"`   // 非必须; 标签ID列表, 多个接收者用‘|’分隔. 当touser为@all时忽略本参数

	MsgType string `json:"msgtype"`        // 必须; 消息类型
	AgentID int64  `json:"agentid"`        // 必须; 企业应用的id, 整型
	Safe    *int   `json:"safe,omitempty"` // 非必须; 表示是否是保密消息, 0表示否, 1表示是, 默认0
}

//Text 文本消息
type Text struct {
	Content string `json:"content"`
}

//Image 图片消息
type Image struct {
	MediaID string `json:"media_id"` // 图片媒体文件id, 可以调用上传媒体文件接口获取

}

//Voice 语音消息
type Voice struct {
	MediaID string `json:"media_id"` // 语音文件id, 可以调用上传媒体文件接口获取

}

//Video 视频消息
type Video struct {
	MediaID     string `json:"media_id"`              // 视频媒体文件id, 可以调用上传媒体文件接口获取
	Title       string `json:"title,omitempty"`       // 视频消息的标题
	Description string `json:"description,omitempty"` // 视频消息的描述

}

//File 文件消息
type File struct {
	MediaID string `json:"media_id"` // 媒体文件id, 可以调用上传媒体文件接口获取

}

//Result 发送消息的返回结果
type Result struct {
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
	InvaliUser  string `json:"invaliduser"`
	InvaliParty string `json:"invalidparty"`
	InvalidTag  string `json:"invalidtag"`
}

//NewsArticle 图文消息
type NewsArticle struct {
	Title       string `json:"title,omitempty"`       // 图文消息标题
	Description string `json:"description,omitempty"` // 图文消息描述
	URL         string `json:"url,omitempty"`         // 点击后跳转的链接.
	PicURL      string `json:"picurl,omitempty"`      // 图文消息的图片链接, 支持JPG, PNG格式, 较好的效果为大图640*320, 小图80*80. 如不填, 在客户端不显示图片
}

//NewsArticleCountLimit 图文消息最多10个
const NewsArticleCountLimit = 10

// News 消息, 注意沒有 Safe 字段.
type News struct {
	Articles []NewsArticle `json:"articles,omitempty"` // 图文消息, 一个图文消息支持1到10条图文

}

//MPNewsArticle 图文消息
type MPNewsArticle struct {
	ThumbMediaID     string `json:"thumb_media_id"`               // 图文消息缩略图的media_id, 可以在上传多媒体文件接口中获得. 此处thumb_media_id即上传接口返回的media_id
	Title            string `json:"title"`                        // 图文消息的标题
	Author           string `json:"author,omitempty"`             // 图文消息的作者
	ContentSourceURL string `json:"content_source_url,omitempty"` // 图文消息点击"阅读原文"之后的页面链接
	Content          string `json:"content"`                      // 图文消息的内容, 支持html标签
	Digest           string `json:"digest,omitempty"`             // 图文消息的描述
	ShowCoverPic     int    `json:"show_cover_pic"`               // 是否显示封面, 1为显示, 0为不显示
}

//SetShowCoverPic 设置是否显示封面
func (article *MPNewsArticle) SetShowCoverPic(b bool) {
	if b {
		article.ShowCoverPic = 1
	} else {
		article.ShowCoverPic = 0
	}
}

// MPNews 消息与 News 消息类似, 不同的是图文消息内容存储在微信后台, 并且支持保密选项.
type MPNews struct {
	Articles []MPNewsArticle `json:"articles,omitempty"` // 多条图文消息信息, 默认第一个item为大图, 注意, 如果图文数超过10, 则将会无响应

}

//MarshalJSON 实现接口自定义MarshalJSON,并且对图文消息进行了检查
func (m Message) MarshalJSON() ([]byte, error) {
	result := make(map[string]interface{})
	if strings.Trim(m.ToUser, " ") != "" {
		result["touser"] = m.ToUser
	}
	if strings.Trim(m.ToParty, " ") != "" {
		result["toparty"] = m.ToParty
	}
	if strings.Trim(m.ToTag, " ") != "" {
		result["totag"] = m.ToTag
	}

	result["msgtype"] = m.MsgType
	result["agentid"] = m.AgentID
	result["safe"] = *m.Safe

	switch msg := m.Content.(type) {
	case Text:
		result[MsgTypeText] = m.Content
	case File:
		result[MsgTypeFile] = m.Content
	case Image:
		result[MsgTypeImage] = m.Content
	case Video:
		result[MsgTypeVideo] = m.Content
	case Voice:
		result[MsgTypeVoice] = m.Content
	case News:
		n := len(msg.Articles)
		if n <= 0 {
			return nil, errors.New("没有有效的图文消息")
		}
		if n > NewsArticleCountLimit {
			return nil, fmt.Errorf("图文消息的文章个数不能超过 %d, 现在为 %d", NewsArticleCountLimit, n)
		}
		result[MsgTypeNews] = m.Content
		delete(result, "safe")
	case MPNews:
		n := len(msg.Articles)
		if n <= 0 {
			return nil, errors.New("没有有效的图文消息")
		}
		if n > NewsArticleCountLimit {
			return nil, fmt.Errorf("图文消息的文章个数不能超过 %d, 现在为 %d", NewsArticleCountLimit, n)
		}
		result[MsgTypeMPNews] = m.Content
	default:
		return nil, ErrUnexpectedMesseage(m)
	}

	return json.Marshal(result)

}

//UnmarshalJSON 实现接口
func (m *Message) UnmarshalJSON(data []byte) error {
	var d []byte
	copy(d, data)

	var msg struct {
		Header
		Text   Text
		File   File
		Image  Image
		News   News
		MPNews MPNews
		Voice  Voice
		Video  Video
	}

	err := json.Unmarshal(d, &msg)
	if err != nil {
		return err
	}

	m.Header = msg.Header

	switch m.MsgType {
	case MsgTypeText:
		m.Content = msg.Text
	case MsgTypeFile:
		m.Content = msg.File
	case MsgTypeImage:
		m.Content = msg.Image
	case MsgTypeMPNews:
		m.Content = msg.MPNews
	case MsgTypeNews:
		m.Content = msg.News
	case MsgTypeVideo:
		m.Content = msg.Video
	case MsgTypeVoice:
		m.Content = msg.Voice
	default:
		return ErrUnexpectedMesseage(string(d))
	}

	return nil
}

//NewText 实例化一个文本消息
func NewText(agenid int64, toAll, issafe bool, user, party, tag []string, textcontent string) *Message {
	msg, _ := NewMessage(agenid, toAll, issafe, user, party, tag, Text{Content: textcontent})
	return msg
}

//NewMessage 实例化一个微信消息
func NewMessage(agenid int64, toAll, issafe bool, user, party, tag []string, msg Contenter) (*Message, error) {

	var touser, toparty, totag, msgtype string
	safe := 0
	switch {
	case toAll:
		touser = "@all"
	default:

		touser = strings.TrimRight(strings.Join(user, "|"), "|")
		toparty = strings.TrimRight(strings.Join(party, "|"), "|")
		totag = strings.TrimRight(strings.Join(tag, "|"), "|")

	}
	if issafe {
		safe = 1
	}
	msgtype = func() string {
		switch msg.(type) {
		case Text:
			return MsgTypeText
		case File:
			return MsgTypeFile
		case Image:
			return MsgTypeImage
		case Video:
			return MsgTypeVideo
		case Voice:
			return MsgTypeVoice
		case News:
			return MsgTypeNews
		case MPNews:
			return MsgTypeMPNews
		default:
			return ""
		}

	}()
	msgheader := Header{
		ToUser:  touser,
		ToParty: toparty,
		ToTag:   totag,
		MsgType: msgtype,
		AgentID: agenid,
		Safe:    &safe,
	}
	if msgheader.MsgType == "" {
		return nil, ErrUnexpectedMesseage(msg)
	}

	return &Message{Header: msgheader, Content: msg}, nil

}

//SendMessage 发送消息,暂时没有封装数量操作的情况
func SendMessage(c *client.Client, msg Message) (*Result, error) {

	resp, err := c.PostJSON("https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s", msg)
	if err != nil {

		return nil, err
	}

	var res Result
	err = json.Unmarshal(resp, &res)
	if err != nil {
		return nil, err
	}
	if res.ErrCode != 0 {
		fmt.Println(res)
		return &res, fmt.Errorf("返回状态不为0，请处理。")
	}
	return &res, nil
}

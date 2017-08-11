package msg

// import (
// 	"fmt"
// 	"testing"
// )

// //测试加密
// func TestEncryptmsg(t *testing.T) {
// 	wxbiz := WXBizMsgCrypt{}
// 	wxbiz.token = "myjs"
// 	wxbiz.encodingAESKey = "A270O4fF5jGa7IMbJ9M2Ql3HHfy5pWzCq8MHgApqtGa"
// 	wxbiz.appid = "wxf9e88950a29b0ca3"

// 	//signature := "a772a5cb65d14851a7b0d25bd8ea3863dfc43f5b"
// 	//timestamp = 1415350442
// 	//nonce = "875469987"
// 	//encrypt_type = "aes"
// 	//msg_signature = "04f265be17ce41c62a689f623b58cfcb17e40ddb"

// 	sReqMsgSig := "04f265be17ce41c62a689f623b58cfcb17e40ddb"
// 	sReqTimeStamp := "1415350442"
// 	sReqNonce := "875469987"

// 	//sRespData := "<xml><ToUserName><![CDATA[mycreate]]></ToUserName><FromUserName><![CDATA[wx582测试一下中文的情况，消息长度是按字节来算的396d3bd56c7]]></FromUserName><CreateTime>1348831860</CreateTime><MsgType><![CDATA[text]]></MsgType><Content><![CDATA[this is a test]]></Content><MsgId>1234567890123456</MsgId></xml>"
// 	sReqData := `<xml>
//     <ToUserName><![CDATA[gh_ecd1786c7c6d]]></ToUserName>
//     <FromUserName><![CDATA[oVient4B1pjm5YWnP5KkeS68NoFM]]></FromUserName>
//     <CreateTime>1415350442</CreateTime>
//     <MsgType><![CDATA[text]]></MsgType>
//     <Content><![CDATA[ooooo]]></Content>
//     <MsgId>6078883860970254946</MsgId>
//     <Encrypt><![CDATA[tvOUUCfWWlCWNbMeVen/7EWkW9TVMm2xWtjWY6nG0xBgQq7Kv9D59OD6SxTzQgGfLLzNAZ5fll5MDBrahiOtRndfz7Zd/MiAlEzkUWqivudE1Xtiovk+xJ8yFL1qVZQ+3lIk8TGGv/v9SwWkcl7ILIslrJrWCzpWonbtbQ1N+euGczGAzCZRPAmm8mfpJyi6PkHq3ShuyRXuJg0Uv3sKFzf2QnFR9njbXIps6KhpaZbKUOX2DONnU4raZU06UVXHHm3SnTrqcDFOSi8+Bv6AdAYsvth4tkM3/97vdtV58ZvYSxAJNzn9ac8DUmCrLa53qRc5sD5+3kBANBQsDItieMKv8gkDZ2/fYu5vP58MTyB8Q57DOItxmV/JfD5VSE9i7wpIhrD4zaNV36IlADVzZezkqxDAdgVFr358RFVtvRA=]]></Encrypt>
// </xml>`

// 	//encryptxml, err := wxbiz.EncryptMsg(sRespData, sReqTimeStamp, sReqNonce)
// 	//if err != nil {
// 	//	fmt.Println(err)
// 	//} else {
// 	//	fmt.Println(encryptxml)
// 	//}

// 	ori, err := wxbiz.DecryptMsg(sReqMsgSig, sReqTimeStamp, sReqNonce, sReqData)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(ori)
// 	}

// }

// //func Testdecryptmsg(t *testing.T) {
// //	encodingAesKey := "jWmYm7qr5nMoAUwZRjGtBxmz3KA1tkAj3ykkR6q2B2C"
// //	token := "QDG6eK"

// //	appId := "wx5823bf96d3bd56c7"

// //	wxbiz := WXBizMsgCrypt{}
// //	wxbiz.Token = token
// //	wxbiz.EncodingAESKey = encodingAesKey
// //	wxbiz.Appid = appId

// //	//	sReqMsgSig := "477715d11cdb4164915debcba66cb864d751f3e6"
// //	sReqTimeStamp := "1409659813"
// //	sReqNonce := "1372623149"
// //	sReqData := "<xml><ToUserName><![CDATA[wx5823bf96d3bd56c7]]></ToUserName><Encrypt><![CDATA[RypEvHKD8QQKFhvQ6QleEB4J58tiPdvo+rtK1I9qca6aM/wvqnLSV5zEPeusUiX5L5X/0lWfrf0QADHHhGd3QczcdCUpj911L3vg3W/sYYvuJTs3TUUkSUXxaccAS0qhxchrRYt66wiSpGLYL42aM6A8dTT+6k4aSknmPj48kzJs8qLjvd4Xgpue06DOdnLxAUHzM6+kDZ+HMZfJYuR+LtwGc2hgf5gsijff0ekUNXZiqATP7PF5mZxZ3Izoun1s4zG4LUMnvw2r+KqCKIw+3IQH03v+BCA9nMELNqbSf6tiWSrXJB3LAVGUcallcrw8V2t9EL4EhzJWrQUax5wLVMNS0+rUPA3k22Ncx4XXZS9o0MBH27Bo6BpNelZpS+/uh9KsNlY6bHCmJU9p8g7m3fVKn28H3KDYA5Pl/T8Z1ptDAVe0lXdQ2YoyyH2uyPIGHBZZIs2pDBS8R07+qN+E7Q==]]></Encrypt></xml>"
// //	//sMsg := "" //解析之后的明文
// //	fmt.Println("======测试解密=====")
// //	smsg, err := wxbiz.EncryptMsg(sReqData, sReqTimeStamp, sReqNonce)

// //	if err != nil {
// //		fmt.Println(err)
// //	} else {
// //		fmt.Println(smsg)
// //	}

// //}

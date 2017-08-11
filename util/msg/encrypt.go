package msg

// import (
// 	"encoding/base64"
// 	"errors"
// 	"fmt"
// 	"strconv"
// 	"time"

// 	"github.com/yl10/wechat/qy/model"
// 	"github.com/yl10/wechat/util"
// )

// //WXBizMsgCrypt WXBizMsgCrypt
// type WXBizMsgCrypt struct {
// 	token          string
// 	encodingAESKey string
// 	appid          string
// }

// //NewWXBizMsgCrypt NewWXBizMsgCrypt
// func NewWXBizMsgCrypt(token, appid, encodingaeskey string) *WXBizMsgCrypt {
// 	biz := &WXBizMsgCrypt{}
// 	//if token == "" || appid == "" {
// 	//	return biz, errors.New("token or appid error!")
// 	//}
// 	//if len(encodingaeskey) != 43 {
// 	//	return biz, errors.New("encodingaeskey is wrong!")
// 	//}
// 	biz.token = token
// 	biz.appid = appid
// 	biz.encodingAESKey = encodingaeskey
// 	return biz
// }

// //EncryptMsg 加密
// //将公众号回复用户的消息加密打包
// // @param sReplyMsg:公众号待回复用户的消息，xml格式的字符串
// // @param sTimeStamp: 时间戳，可以自己生成，也可以用URL参数的timestamp
// // @param sNonce: 随机串，可以自己生成，也可以用URL参数的nonce
// // @param sEncryptMsg: 加密后的可以直接回复用户的密文，包括msg_signature, timestamp, nonce, encrypt的xml格式的字符串,当return返回0时有效
// // return：成功0，失败返回对应的错误码
// func (msg *WXBizMsgCrypt) EncryptMsg(sReplyMsg, sTimeStamp, sNonce string) (string, error) {
// 	//加密待发送的消息
// 	encrypt, err := msg.encrypt(sReplyMsg)
// 	if err != nil {
// 		return "", err
// 	}
// 	//生成安全签名
// 	if sTimeStamp == "" {
// 		sTimeStamp = strconv.FormatInt(time.Now().Unix(), 10)
// 	}
// 	signature := getsignture(this.token, sTimeStamp, sNonce, encrypt)
// 	// fmt.Println("加密的时候的签名是:%v", signature)
// 	// fmt.Println(this.token)
// 	// fmt.Println(sTimeStamp)
// 	// fmt.Println(sNonce)
// 	// fmt.Println(encrypt)
// 	return model.EncryptMsgGenerate(encrypt, signature, sTimeStamp, sNonce), nil
// }

// /**
//  * 对明文进行加密.
//  *
//  * @param text 需要加密的明文
//  * @return 加密后base64编码的字符串
//  * @throws AesException aes加密失败
//  */
// func (this *WXBizMsgCrypt) encrypt(text string) (string, error) {
// 	//其中，msg_encrypt = Base64_Encode( AES_Encrypt[ random(16B) + msg_len(4B) + msg + $AppId] )
// 	//16位随机数
// 	//
// 	random := []byte(util.RandomAlphanumeric(16))
// 	msglen := []byte(getNetworkBytesOrder(len(text)))
// 	msg := []byte(text)
// 	appid := []byte(this.appid)
// 	prebyte := []byte{}
// 	prebyte = append(prebyte, random...)
// 	prebyte = append(prebyte, msglen...)
// 	prebyte = append(prebyte, msg...)
// 	prebyte = append(prebyte, appid...)

// 	aeskey, err := AESKeyDecode(this.encodingAESKey)
// 	if err != nil {
// 		return "", err
// 	}
// 	v := []byte(aeskey)
// 	s, err := AesEncrypt(prebyte, v[:32])
// 	if err != nil {
// 		return "", err
// 	}
// 	return s, nil
// }

// /**
//  * 检验消息的真实性，并且获取解密后的明文.
//  * <ol>
//  * 	<li>利用收到的密文生成安全签名，进行签名验证</li>
//  * 	<li>若验证通过，则提取xml中的加密消息</li>
//  * 	<li>对消息进行解密</li>
//  * </ol>
//  *
//  * @param msgSignature 签名串，对应URL参数的msg_signature
//  * @param timeStamp 时间戳，对应URL参数的timestamp
//  * @param nonce 随机串，对应URL参数的nonce
//  * @param postData 密文，对应POST请求的数据
//  *
//  * @return 解密后的原文
//  * @throws AesException 执行失败，请查看该异常的错误码和具体的错误信息
//  */
// func (this *WXBizMsgCrypt) DecryptMsg(msgSignature, timeStamp, nonce, postData string) (string, error) {

// 	encrypt, err := model.EncryptMsgExtract(postData)
// 	if err != nil {
// 		return "", err
// 	}
// 	fmt.Println(encrypt)
// 	fmt.Println("提取加密的字符串成功")
// 	if msgSignature != getsignture(this.token, timeStamp, nonce, encrypt) {
// 		fmt.Println("msgSignature:=%s,计算的签名:%s", msgSignature, getsignture(this.token, timeStamp, nonce, encrypt))
// 		return "", errors.New("qianmingbudui")
// 	}
// 	fmt.Println("检验签名成功")
// 	result, err := this.decrypt(encrypt)
// 	if err != nil {
// 		return "", err
// 	}

// 	return string(result), nil
// }

// /**
//  * 对密文进行解密.
//  *
//  * @param text 需要解密的密文
//  * @return 解密得到的明文
//  * @throws AesException aes解密失败
//  */
// func (this *WXBizMsgCrypt) decrypt(text string) ([]byte, error) {
// 	//解密
// 	xmlContent := []byte{}
// 	aeskey, err := AESKeyDecode(this.encodingAESKey)
// 	if err != nil {
// 		return xmlContent, err
// 	}
// 	v := []byte(aeskey)
// 	original, err := AesDecrypt(text, string(v[:32]))
// 	if err != nil {
// 		return xmlContent, err
// 	}
// 	fmt.Println(len(original))
// 	// 分离16位随机字符串,网络字节序和AppId
// 	networkorder := original[16:20]
// 	fmt.Println(networkorder)
// 	xmlLength := recoverNetworkBytesOrder(networkorder)
// 	fmt.Println(xmlLength)
// 	xmlContent = original[20 : 20+xmlLength]
// 	from_appid := original[20+xmlLength:]

// 	if this.appid != string(from_appid) {

// 		fmt.Println(this.appid)
// 		fmt.Println(string(from_appid))
// 		return xmlContent, errors.New("appid不对")
// 	}

// 	return xmlContent, nil

// }
// func getNetworkBytesOrder(n int) []byte {
// 	orderBytes := make([]byte, 4)
// 	orderBytes[0] = byte(n >> 24)
// 	orderBytes[1] = byte(n >> 16)
// 	orderBytes[2] = byte(n >> 8)
// 	orderBytes[3] = byte(n)
// 	return orderBytes
// }

// func recoverNetworkBytesOrder(orderBytes []byte) int {

// 	if len(orderBytes) != 4 {
// 		panic("the length of orderBytes must be equal to 4")
// 	}
// 	n := int(orderBytes[0])<<24 |
// 		int(orderBytes[1])<<16 |
// 		int(orderBytes[2])<<8 |
// 		int(orderBytes[3])
// 	return n

// }

// // 把整数 n 格式化成 4 字节的网络字节序
// func encodeNetworkBytesOrder(n int, orderBytes []byte) {
// 	if len(orderBytes) != 4 {
// 		panic("the length of orderBytes must be equal to 4")
// 	}
// 	orderBytes[0] = byte(n >> 24)
// 	orderBytes[1] = byte(n >> 16)
// 	orderBytes[2] = byte(n >> 8)
// 	orderBytes[3] = byte(n)
// }

// // 从 4 字节的网络字节序里解析出整数
// func decodeNetworkBytesOrder(orderBytes []byte) (n int) {
// 	if len(orderBytes) != 4 {
// 		panic("the length of orderBytes must be equal to 4")
// 	}
// 	n = int(orderBytes[0])<<24 |
// 		int(orderBytes[1])<<16 |
// 		int(orderBytes[2])<<8 |
// 		int(orderBytes[3])
// 	return
// }

// func AESKeyDecode(encodingAESKey string) (AESKey []byte, err error) {
// 	if len(encodingAESKey) != 43 {
// 		err = errors.New("the length of encodedAESKey is" + string(len(encodingAESKey)) + ",not equal to 43")
// 		return
// 	}
// 	return base64.StdEncoding.DecodeString(encodingAESKey + "=")
// }

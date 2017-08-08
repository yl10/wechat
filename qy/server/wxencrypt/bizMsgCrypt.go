package wxencrypt

import (
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"math/rand"
	"time"

	. "github.com/yl10/wechat/util"
)

//WXBizMsgCrypt 企业微信加解密对象
type WXBizMsgCrypt struct {
	token string
	//encodingAesKey string //EncodingAESKey用于消息体的加密，长度固定为43个字符，从a-z, A-Z, 0-9共62个字符中选取，是AESKey的Base64编码。解码后即为32字节长的AESKey
	corpID string
	aesKey string
}

//NewWxBIzMsgCrypt new
func NewWxBIzMsgCrypt(token, encodingAesKey, corpID string) (*WXBizMsgCrypt, error) {
	aesKey, err := AESKeyDecode(encodingAesKey)
	if err != nil {
		return nil, err
	}
	return &WXBizMsgCrypt{
		token:  token,
		aesKey: string(aesKey),
		corpID: corpID,
	}, nil
}

//getNetworkBytesOrder 生成4个字节的网络字节序
func getNetworkBytesOrder(n int) []byte {
	orderBytes := make([]byte, 4)
	orderBytes[0] = byte(n >> 24)
	orderBytes[1] = byte(n >> 16)
	orderBytes[2] = byte(n >> 8)
	orderBytes[3] = byte(n)
	return orderBytes
}

//recoverNetworkBytesOrder  还原4个字节的网络字节序
func recoverNetworkBytesOrder(orderBytes []byte) int {

	if len(orderBytes) != 4 {
		panic("the length of orderBytes must be equal to 4")
	}
	n := int(orderBytes[0])<<24 |
		int(orderBytes[1])<<16 |
		int(orderBytes[2])<<8 |
		int(orderBytes[3])
	return n

}

//AESKeyDecode 对encodingAESKey进行解密，得到AESKey
/*AESKey=Base64_Decode(EncodingAESKey + “=”)，是AES算法的密钥，长度为32字节。AES采用CBC模式，数据采用PKCS#7填充至32字节的倍数；IV初始向量大小为16字节，取AESKey前16字节。具体详见：http://tools.ietf.org/html/rfc2315
 */
func AESKeyDecode(encodingAESKey string) (AESKey []byte, err error) {
	if len(encodingAESKey) != 43 {
		return nil, GenAesErr(IllegalAesKey)
	}
	keyData, err := base64.StdEncoding.DecodeString(encodingAESKey + "=")
	if err != nil {
		return nil, GenAesErr(DecryptAESError)
	}
	return keyData, nil
}

//getRandomStr 随机生成16位字符串
func getRandomStr() string {
	base := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	stringBuffer := ""
	for i := 0; i < 16; i++ {
		stringBuffer += string(base[random.Intn(len(base))])
	}
	return stringBuffer
}

/*DecryptMsg 检验消息的真实性，并且获取解密后的明文.
 * <ol>
 * 	<li>利用收到的密文生成安全签名，进行签名验证</li>
 * 	<li>若验证通过，则提取xml中的加密消息</li>
 * 	<li>对消息进行解密</li>
 * </ol>
 *
 * @param msgSignature 签名串，对应URL参数的msg_signature
 * @param timeStamp 时间戳，对应URL参数的timestamp
 * @param nonce 随机串，对应URL参数的nonce
 * @param postData 密文，对应POST请求的数据
 *
 * @return 解密后的原文
 * @throws AesException 执行失败，请查看该异常的错误码和具体的错误信息
 */
func (w *WXBizMsgCrypt) DecryptMsg(msgSignature, timeStamp, nonce string, postData []byte) ([]byte, error) {
	//提取密文消息
	reqmsq := EncryptedRequestMsq{}
	err := xml.Unmarshal(postData, &reqmsq)
	if err != nil {
		return nil, GenAesErr(ParseXmlError)
	}

	encryptedXMLMsg := reqmsq.Encrypt.String() //加密的xml消息

	msgByte, err := w.decrypt(encryptedXMLMsg)

	return msgByte, err
}

/*VerifyURL  验证URL
 * @param msgSignature 签名串，对应URL参数的msg_signature
 * @param timeStamp 时间戳，对应URL参数的timestamp
 * @param nonce 随机串，对应URL参数的nonce
 * @param echoStr 随机串，对应URL参数的echostr
 *
 * @return 解密之后的echostr
 * @throws AesException 执行失败，请查看该异常的错误码和具体的错误信息
 */
func (w *WXBizMsgCrypt) VerifyURL(msgSignature, timeStamp, nonce, echoStr string) (string, error) {
	signture := GetSHA1(w.token, timeStamp, nonce, echoStr)
	if signture != msgSignature {
		return "", GenAesErr(ValidateSignatureError)
	}
	result, err := w.decrypt(echoStr)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

/**encrypt 对明文进行加密.
 *
 *
 * @param text 需要加密的明文
 * @return 加密后base64编码的字符串
 * @throws AesException aes加密失败
 */

func (w *WXBizMsgCrypt) encrypt(random, text string) (string, error) {
	//其中，msg_encrypt = Base64_Encode( AES_Encrypt[ random(16B) + msg_len(4B) + msg + $AppId] )
	//16位随机数

	msglen := []byte(getNetworkBytesOrder(len(text)))
	msg := []byte(text)
	corpid := []byte(w.corpID)
	prebyte := []byte{}
	prebyte = append(prebyte, random...)
	prebyte = append(prebyte, msglen...)
	prebyte = append(prebyte, msg...)
	prebyte = append(prebyte, corpid...)

	aesencrypt, err := AesEncrypt(prebyte, []byte(w.aesKey))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(aesencrypt), nil
}

/**decrypt对密文进行解密.
 *
 * @param text 需要解密的密文
 * @return 解密得到的明文
 * @throws AesException aes解密失败
 *对密文BASE64解码：aes_msg=Base64_Decode(msg_encrypt)
 *使用AESKey做AES解密：rand_msg=AES_Decrypt(aes_msg)
 *验证解密后$CorpID、msg_len，当为第三方套件回调事件时，CorpID的内容为suiteid
 *去掉rand_msg头部的16个随机字节，4个字节的msg_len,和尾部的$CorpID即为最终的消息体原文msg
 */
func (w *WXBizMsgCrypt) decrypt(msgEncrypt string) ([]byte, error) {

	//对密文BASE64解码：aes_msg=Base64_Decode(msg_encrypt)
	aesMsg, err := base64.StdEncoding.DecodeString(msgEncrypt)
	if err != nil {
		return nil, GenAesErr(DecodeBase64Error)
	}
	//使用AESKey做AES解密：rand_msg=AES_Decrypt(aes_msg)
	randMsg, err := AesDecrypt(string(aesMsg), w.aesKey)
	if err != nil {
		return nil, GenAesErr(DecryptAESError)
	}

	networkorder := randMsg[16:20]
	xmlLength := recoverNetworkBytesOrder(networkorder)
	xmlContent := randMsg[20 : 20+xmlLength]
	fromCorpID := randMsg[20+xmlLength:]

	if string(fromCorpID) != w.corpID {
		return nil, GenAesErr(ValidateCorpidError)
	}
	return xmlContent, nil
}

/*EncryptMsg 将企业微信回复用户的消息加密打包.
 *
 * <ol>
 * 	<li>对要发送的消息进行AES-CBC加密</li>
 * 	<li>生成安全签名</li>
 * 	<li>将消息密文和安全签名打包成xml格式</li>
 * </ol>
 *
 * @param replyMsg 企业微信待回复用户的消息，xml格式的字符串
 * @param timeStamp 时间戳，可以自己生成，也可以用URL参数的timestamp
 * @param nonce 随机串，可以自己生成，也可以用URL参数的nonce
 *
 * @return 加密后的可以直接回复用户的密文，包括msg_signature, timestamp, nonce, encrypt的xml格式的字符串
 * @throws AesException 执行失败，请查看该异常的错误码和具体的错误信息
 */
func (w *WXBizMsgCrypt) EncryptMsg(replyMsg, timeStamp, nonce string) (string, error) {

	//消息加密
	encrypt, err := w.encrypt(getRandomStr(), replyMsg)
	if err != nil {
		return "", err
	}
	//生成安全签名

	if timeStamp == "" {
		timeStamp = fmt.Sprintf("%d", time.Now().UnixNano())
	}

	signature := GetSHA1(w.token, timeStamp, nonce, encrypt)

	format := `<xml>
	<Encrypt><![CDATA[%s]]></Encrypt>
	<MsgSignature><![CDATA[%s]]></MsgSignature>
	<TimeStamp>%s</TimeStamp>
	<Nonce><![CDATA[%s]]></Nonce>
</xml>`

	return fmt.Sprintf(format, encrypt, signature, timeStamp, nonce), nil

}

//EncryptedRequestMsq 微信POST过来的加密过的消息
type EncryptedRequestMsq struct {
	ToUserName CDATA
	Encrypt    CDATA
	AgentID    CDATA
}

func extract(xmlByte []byte) (EncryptedRequestMsq, error) {
	reqmsq := EncryptedRequestMsq{}
	err := xml.Unmarshal(xmlByte, &reqmsq)
	if err != nil {
		return reqmsq, GenAesErr(ParseXmlError)
	}
	return reqmsq, nil
}

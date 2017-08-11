package wxencrypt

import (
	"crypto/sha1"
	"fmt"
	"sort"
)

//GetSHA1 计算企业微信的消息签名
//@param token 票据
//@param tiemstamp 时间戳
//@param nonce 随机字符串
//@param encrypt 密文消息
func GetSHA1(token, timestamp, nonce, encrypt string) string {
	strs := sort.StringSlice{token, timestamp, nonce, encrypt}
	sort.Strings(strs)
	var str string
	for _, s := range strs {
		str += s
	}
	h := sha1.New()
	h.Write([]byte(str))
	return fmt.Sprintf("%x", h.Sum(nil))
}

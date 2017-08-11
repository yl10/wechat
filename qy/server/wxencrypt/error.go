package wxencrypt

import (
	"fmt"
)

const (
	ValidateSignatureError EncryptError = -40001
	ParseXmlError                       = -40002
	ComputeSignatureError               = -40003
	IllegalAesKey                       = -40004
	ValidateCorpidError                 = -40005
	EncryptAESError                     = -40006
	DecryptAESError                     = -40007
	IllegalBuffer                       = -40008
	EncodeBase64Error                   = -40009
	DecodeBase64Error                   = -40010
	GenReturnXmlError                   = -40011
)

type EncryptError int

var errMap = map[EncryptError]string{
	ValidateSignatureError: "签名验证错误",
	ParseXmlError:          "xml解析失败",
	ComputeSignatureError:  "sha加密生成签名失败",
	IllegalAesKey:          "encodingAesKey 非法",
	ValidateCorpidError:    "corpid 校验错误",
	EncryptAESError:        "aes 加密失败",
	DecryptAESError:        "aes 解密失败",
	IllegalBuffer:          "解密后得到的buffer非法",
	EncodeBase64Error:      "base64加密失败",
	DecodeBase64Error:      "base64解密失败",
	GenReturnXmlError:      "生成xml失败",
}

func GenAesErr(e EncryptError) error {
	return fmt.Errorf("code:%v,msg:%s", e, errMap[e])
}

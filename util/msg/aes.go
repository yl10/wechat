package msg

// import (
// 	"bytes"
// 	"crypto/aes"
// 	"crypto/cipher"
// 	"crypto/sha1"
// 	"encoding/base64"
// 	"fmt"
// 	"sort"
// )

// func AesEncrypt(origData, key []byte) (string, error) {

// 	block, err := aes.NewCipher(key)
// 	if err != nil {
// 		return "", err
// 	}
// 	blockSize := block.BlockSize()
// 	//origData = PKCS5Padding(origData, blockSize)
// 	origData = ZeroPadding(origData, blockSize)
// 	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
// 	crypted := make([]byte, len(origData))
// 	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
// 	// crypted := origData
// 	blockMode.CryptBlocks(crypted, origData)
// 	return base64.StdEncoding.EncodeToString(crypted), nil
// }

// func AesDecrypt(strcrypted, strkey string) ([]byte, error) {

// 	basedcrypted, err1 := base64.StdEncoding.DecodeString(strcrypted)
// 	if err1 != nil {
// 		return []byte{}, err1
// 	}
// 	crypted := []byte(basedcrypted)
// 	key := []byte(strkey)

// 	block, err := aes.NewCipher([]byte(key))
// 	if err != nil {
// 		return []byte{}, err
// 	}

// 	blockSize := block.BlockSize()
// 	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
// 	origData := make([]byte, len(crypted))
// 	//origData := crypted
// 	blockMode.CryptBlocks(origData, crypted)
// 	//origData = PKCS5UnPadding(origData)
// 	origData = ZeroUnPadding(origData)

// 	return origData, nil
// }

// func ZeroPadding(ciphertext []byte, blockSize int) []byte {
// 	padding := blockSize - len(ciphertext)%blockSize
// 	padtext := bytes.Repeat([]byte{0}, padding)
// 	return append(ciphertext, padtext...)
// }

// func ZeroUnPadding(origData []byte) []byte {
// 	length := len(origData)
// 	unpadding := int(origData[length-1])
// 	return origData[:(length - unpadding)]
// }

// func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
// 	padding := blockSize - len(ciphertext)%blockSize
// 	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
// 	return append(ciphertext, padtext...)
// }

// func PKCS5UnPadding(origData []byte) []byte {
// 	length := len(origData)
// 	// 去掉最后一个字节 unpadding 次
// 	unpadding := int(origData[length-1])
// 	return origData[:(length - unpadding)]
// }

// func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
// 	padding := blockSize - len(ciphertext)%blockSize
// 	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
// 	return append(ciphertext, padtext...)
// }

// func PKCS7UnPadding(origData []byte) []byte {
// 	length := len(origData)
// 	// 去掉最后一个字节 unpadding 次
// 	unpadding := int(origData[length-1])
// 	return origData[:(length - unpadding)]
// }

// func getsignture(token, timeStamp, nonce, encrypt string) string {
// 	strs := sort.StringSlice{token, timeStamp, nonce, encrypt}
// 	sort.Strings(strs)
// 	var str string
// 	for _, s := range strs {
// 		str += s
// 	}
// 	h := sha1.New()
// 	h.Write([]byte(str))
// 	return fmt.Sprintf("%x", h.Sum(nil))
// }

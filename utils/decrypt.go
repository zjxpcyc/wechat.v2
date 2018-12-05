package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

// Base64Decode Base64 解密
func Base64Decode(dt string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(dt)
}

// AESP7Decrypt  AES-128-CBC 数据(PKCS#7)解密
// 参考: https://github.com/medivhzhan/weapp/blob/master/util/crypto.go
func AESP7Decrypt(dt, key []byte, iv ...[]byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	dtLen := len(dt)
	if dtLen < aes.BlockSize {
		return nil, errors.New("解密数据太小")
	}

	var theIV []byte
	if iv != nil && len(iv) > 0 {
		theIV = iv[0][:aes.BlockSize]
	} else {
		theIV = dt[:aes.BlockSize]
	}

	mode := cipher.NewCBCDecrypter(block, theIV)

	dist := make([]byte, dtLen)
	mode.CryptBlocks(dist, dt)

	return pkcs7UnPadding(dist), nil
}

func pkcs7UnPadding(dt []byte) []byte {
	length := len(dt)
	unpadding := int(dt[length-1])
	return dt[:(length - unpadding)]
}

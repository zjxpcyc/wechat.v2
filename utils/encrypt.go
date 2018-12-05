package utils

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io"
	"strings"
)

// MD5 md5 加密
func MD5(target string, salt ...string) string {
	h := md5.New()
	io.WriteString(h, target)
	io.WriteString(h, strings.Join(salt, ""))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// HmacSHA256 HmacSha256 加密
func HmacSHA256(target, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	io.WriteString(h, target)
	return fmt.Sprintf("%x", h.Sum(nil))
}

//Sha1  sha1 加密
func Sha1(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func pkcs7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

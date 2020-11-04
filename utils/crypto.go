package utils

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"math/rand"
	"unicode/utf8"
)

//随机生成字符串
func RandomString(len int) string {
	if len < 1 {
		return ""
	}
	token := make([]byte, (len*6)/8+1)
	rand.Read(token)
	s := base64.RawStdEncoding.EncodeToString(token)
	if utf8.RuneCountInString(s) > len {
		return s[0:len]
	}
	return s
}

func Sha1s(input string) string {
	payload := sha1.Sum([]byte(input))
	return base64.RawStdEncoding.EncodeToString(payload[:])
}

func Sha256s(input string) string {
	payload := sha256.Sum256([]byte(input))
	return base64.RawStdEncoding.EncodeToString(payload[:])
}

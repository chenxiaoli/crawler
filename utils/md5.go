package utils

import (
	"crypto/md5"
	"encoding/hex"
)

/*
ToMd5string 把一个字符串MD5成32位字符串
*/
func ToMd5String(s string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(s))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

/*
StringToHash 给字符串hash
*/
func StringToHash(str string) string {
	srcData := []byte(str)
	return BytesToHash(srcData)
}

/*
BytesToHash 给[]byte hash
*/
func BytesToHash(srcData []byte) string {
	hash := md5.New()
	hash.Write(srcData)
	cipherText2 := hash.Sum(nil)
	hexText := make([]byte, 32)
	hex.Encode(hexText, cipherText2)

	return string(hexText)

}

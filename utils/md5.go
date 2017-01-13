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

package util

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

//low
func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

//Upper
func MD5Encode(data string) string {
	return strings.ToUpper(Md5Encode(data))
}

//Valid the password
func ValidatePasswd(plainpwd, salt, passwd string) bool {
	return Md5Encode(plainpwd+salt) == passwd
}

//make md5 password
func MakePasswd(plainpwd, salt string) string {
	return Md5Encode(plainpwd + salt)
}

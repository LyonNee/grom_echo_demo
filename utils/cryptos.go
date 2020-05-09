package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMD5HashCode(message string)string{
	h := md5.New()
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}
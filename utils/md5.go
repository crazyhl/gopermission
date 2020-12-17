package utils

import (
	"crypto/md5"
	"fmt"
)

func GetStrMd5(str string) string {
	if len(str) == 0 {
		return ""
	}
	strBytes := []byte(str)
	md5Bytes := md5.Sum(strBytes)
	return fmt.Sprintf("%x", md5Bytes)
}

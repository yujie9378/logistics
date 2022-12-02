package util

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func Encode(data string) string {
	md := md5.New()
	md.Write([]byte(data))
	return hex.EncodeToString(md.Sum(nil))
}

func Check(content, encrypted string) bool {
	//不区分大小写比较的
	return strings.EqualFold(Encode(content), encrypted)
}

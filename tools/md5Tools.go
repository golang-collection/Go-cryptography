package tools

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

/**
* @Author: super
* @Date: 2020-08-19 10:58
* @Description:
**/

func Check(content, encrypted string) bool {
	return strings.EqualFold(Md5Encode(content), encrypted)
}
func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

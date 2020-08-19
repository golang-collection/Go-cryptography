package tools

import (
	"crypto/md5"
	"encoding/hex"
)

/**
* @Author: super
* @Date: 2020-08-19 10:58
* @Description:
**/

func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

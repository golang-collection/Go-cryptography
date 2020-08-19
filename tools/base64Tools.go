package tools

import (
	"encoding/base64"
	"fmt"
)

/**
* @Author: super
* @Date: 2020-08-19 11:03
* @Description:
**/

const (
	base64Table = "IJjkKLMNO567PQX12RVW3YZaDEFGbcdefghiABCHlSTUmnopqrxyz04stuvw89+/"

)

var coder = base64.NewEncoding(base64Table)

func Base64Encode(str string) string {
	src := []byte(str)
	result := fmt.Sprintf("%s", []byte(coder.EncodeToString(src)))
	return result
}

func Base64Decode(str string) (string, error) {
	bytes, err := coder.DecodeString(str)
	if err != nil{
		return "", err
	}
	result := fmt.Sprintf("%s", bytes)
	return result, nil
}

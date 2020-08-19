package tools

import "testing"

/**
* @Author: super
* @Date: 2020-08-19 15:00
* @Description:
**/

func TestMd5Encode(t *testing.T) {
	var md5Tests = []struct{
		orig string
		result string
	}{
		{"hello world",  "5eb63bbbe01eeed093cb22bb8f5acdc3"},
		{"",  "d41d8cd98f00b204e9800998ecf8427e"},
		{" ",  "7215ee9c7d9dc229d2921a40e899ec5f"},
	}
	for _, tt := range md5Tests{
		actual := Md5Encode(tt.orig)
		if actual != tt.result{
			t.Errorf("orig: %s, expect: %s, actual: %s", tt.orig, tt.result, actual)
		}
	}
}
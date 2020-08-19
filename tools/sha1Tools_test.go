package tools

import "testing"

/**
* @Author: super
* @Date: 2020-08-19 15:03
* @Description:
**/

func TestSha1Encode(t *testing.T) {
	var sha1Tests = []struct{
		orig string
		result string
	}{
		{"hello world",  "2aae6c35c94fcfb415dbe95f408b9ce91ee846ed"},
	}
	for _, tt := range sha1Tests{
		actual := Sha1Encode(tt.orig)
		if actual != tt.result{
			t.Errorf("orig: %s, expect: %s, actual: %s", tt.orig, tt.result, actual)
		}
	}
}
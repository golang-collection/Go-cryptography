package tools

import "testing"

/**
* @Author: super
* @Date: 2020-08-19 14:24
* @Description:
**/

func TestBase64Encode(t *testing.T) {
	var base64Tests = []struct{
		orig string
		result string
	}{
		{"hello world",  "FMYmGM8fc49xGM2="},
	}
	for _, tt := range base64Tests{
		actual := Base64Encode(tt.orig)
		if actual != tt.result{
			t.Errorf("orig: %s, expect: %s, actual: %s", tt.orig, tt.result, actual)
		}
	}
}

func TestBase64Decode(t *testing.T) {
	var base64Tests = []struct{
		orig string
		result string
		err error
	}{
		{"hello world",  "FMYmGM8fc49xGM2=", nil},
		{"", "", nil},
	}
	for _, tt := range base64Tests{
		actual, err := Base64Decode(tt.result)
		if actual != tt.orig || err != tt.err{
			t.Errorf("orig: %s, expect: %s, actual: %s, err: %v", tt.orig, tt.result, actual, err)
		}
	}
}
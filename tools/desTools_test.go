package tools

import (
	"crypto/des"
	"testing"
)

/**
* @Author: super
* @Date: 2020-08-19 14:54
* @Description:
**/

func TestDesEncrypt(t *testing.T) {
	var desTests = []struct {
		orig   string
		key    string
		result string
		err    error
	}{
		{"hello world", "sd45c1a2", "9e2a856cfb8f0868aa9baa407de2503c", nil},
		{"", "1", "", des.KeySizeError(1)},
		{"", "", "", des.KeySizeError(0)},
		{"hello", "", "", des.KeySizeError(0)},
	}
	for _, tt := range desTests {
		s, err := DesEncrypt(tt.orig, tt.key)
		if s != tt.result || err != tt.err {
			t.Errorf("orig: %s, key: %s, expect: %s, actual: %s, err: %v", tt.orig, tt.key, tt.result, s, err)
		}
	}
}

func TestDesDecrypt(t *testing.T) {
	var desTests = []struct {
		orig   string
		key    string
		result string
		err    error
	}{
		{"hello world", "sd45c1a2", "9e2a856cfb8f0868aa9baa407de2503c", nil},
	}
	for _, tt := range desTests {
		s, err := DesDecrypt(tt.result, tt.key)
		if s != tt.orig || err != tt.err {
			t.Errorf("orig: %s, key: %s, expect: %s, actual: %s", tt.result, tt.key, tt.orig, s)
		}
	}
}

package tools

import (
	"crypto/aes"
	"testing"
)

/**
* @Author: super
* @Date: 2020-08-19 14:03
* @Description:
**/


func TestAesEncrypt(t *testing.T) {
	var aesTests = []struct{
		orig string
		key string
		result string
		err error
	}{
		{"hello world", "123456781234567812345678", "cOYnULidg5pVZlS3bxTLpQ==", nil},
		{"", "1", "", aes.KeySizeError(1)},
		{"", "", "", aes.KeySizeError(0)},
		{"hello", "", "", aes.KeySizeError(0)},
	}
	for _, tt := range aesTests{
		s, err := AesEncrypt(tt.orig, tt.key)
		if s != tt.result || err != tt.err{
			t.Errorf("orig: %s, key: %s, expect: %s, actual: %s, err: %v", tt.orig, tt.key, tt.result, s, err)
		}
	}
}

func TestAesDecrypt(t *testing.T) {
	var aesTests = []struct{
		orig string
		key string
		result string
	}{
		{"hello world", "123456781234567812345678", "cOYnULidg5pVZlS3bxTLpQ=="},
	}
	for _, tt := range aesTests{
		s := AesDecrypt(tt.result, tt.key)
		if s != tt.orig{
			t.Errorf("orig: %s, key: %s, expect: %s, actual: %s", tt.result, tt.key, tt.orig, s)
		}
	}
}


func BenchmarkAesEncrypt(b *testing.B) {
	for i := 0; i<b.N; i++{
		_, _ = AesEncrypt("hello world", "123456781234567812345678")
	}
}
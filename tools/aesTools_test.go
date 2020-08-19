package tools

import (
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
	}{
		{"hello world", "123456781234567812345678", "cOYnULidg5pVZlS3bxTLpQ=="},
		{"hello world", "123456781234567812345678", "cOYnULidg5pVZlS3bxTLpQ=="},
		{"hello world", "123456781234567812345678", "cOYnULidg5pVZlS3bxTLpQ=="},
	}
	for _, tt := range aesTests{
		s := AesEncrypt(tt.orig, tt.key)
		if s != tt.result{
			t.Errorf("orig: %s, key: %s, expect: %s, actual: %s", tt.orig, tt.key, tt.result, s)
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
		{"hello world", "123456781234567812345678", "cOYnULidg5pVZlS3bxTLpQ=="},
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
		_ = AesEncrypt("hello world", "123456781234567812345678")
	}
}
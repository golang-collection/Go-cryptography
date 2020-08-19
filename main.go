package main

import (
	"Go-cryptography/tools"
	"fmt"
	"log"
)

/**
* @Author: super
* @Date: 2020-08-19 10:49
* @Description:
**/

func main() {
	fmt.Println("AES~~~~~~~~~~~~~~~~~~~")

	orig := "hello world"
	key := "123456781234567812345678"

	encryptCode := tools.AesEncrypt(orig, key)
	fmt.Println("密文：", encryptCode)

	decryptCode := tools.AesDecrypt(encryptCode, key)
	fmt.Println("解密结果：", decryptCode)

	fmt.Println("DES~~~~~~~~~~~~~~~~~~~")

	newkey := []byte("2fa6c1e9")
	str :="I love this beautiful world!"
	strEncrypted, err := tools.DesEncrypt(str, newkey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("密文:", strEncrypted)
	strDecrypted, err := tools.DesDecrypt(strEncrypted, newkey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("解密结果:", strDecrypted)

	fmt.Println("MD5~~~~~~~~~~~~~~~~~~~")

	strTest := "I love this beautiful world!"
	fmt.Println(tools.Md5Encode(strTest))

	fmt.Println("SHA1~~~~~~~~~~~~~~~~~~~")

	strTest = "I love this beautiful world!"
	tools.Sha1Encode(strTest)

	fmt.Println("BASE64~~~~~~~~~~~~~~~~~~~")

	strTest = "I love this beautiful world!"
	strEncr := tools.Base64Encode(strTest)
	fmt.Println(strEncr)
	fmt.Println(tools.Base64Decode(strEncr))
}

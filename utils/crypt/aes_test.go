package crypt

import (
	"crypto/aes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestAesEncodeDecode(t *testing.T) {
	key, _ := hex.DecodeString("6368616e676520746869732070617373")
	plaintext := []byte("hello 您的")

	c := make([]byte, aes.BlockSize+len(plaintext))
	iv := c[:aes.BlockSize]

	//加密
	ciphertext, err := AesEncrypt(plaintext, key, iv)
	if err != nil {
		panic(err)
	}

	//打印加密base64后密码
	fmt.Println(base64.StdEncoding.EncodeToString(ciphertext))

	//解密
	plaintext, err = AesDecrypt(ciphertext, key, iv)
	if err != nil {
		panic(err)
	}

	//打印解密明文
	fmt.Println(string(plaintext))
}

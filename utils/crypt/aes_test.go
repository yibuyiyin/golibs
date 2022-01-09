package crypt

import (
	"fmt"
	"testing"
)

func TestAesEncodeDecode(t *testing.T) {
	// key的长度必须是16、24或者32字节，分别用于选择AES-128, AES-192, or AES-256
	var aeskey = []byte("12345678abcdefgh")
	pass := []byte("vdncloud123456")
	pass64, err := AesEncrypt(pass, aeskey)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("加密后:%v\n", pass64)
	tpass, err := AesDecrypt(pass64, aeskey)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("解密后:%s\n", tpass)
}

package utils

import (
	"encoding/hex"
	"fmt"
	"gitee.com/itsos/golibs/v2/utils/crypt"
	"strings"
	"testing"
)

func Test_Crypto(t *testing.T) {
	// https://heroims.github.io/2018/11/21/AES-128-CBC%20Base64%E5%8A%A0%E5%AF%86%E2%80%94%E2%80%94OC,Java,Golang%E8%81%94%E8%B0%83/
	key := []byte("846acca7d02f6ba83460f0b626f37468")
	str := []byte("{\"userInfo\":{\"openStatus\":\"0\",\"gender\":\"2\",\"raind\":\"0\",\"oldUserId\":\"1502356914142\",\"certiType\":\"14\",\"updateTime\":\"2017-10-25 16:58:07\",\"mobileNo\":\"13950385870\",\"sysSrc\":\"001001001000\",\"userName\":\"林雨星\",\"birthDate\":\"2005-02-27\",\"userId\":1000000000066363,\"userCode\":\"M1502356914142\",\"policyLevel\":\"T\",\"checkStatus\":\"2\",\"userLevel\":\"1\",\"userPsw\":\"\",\"validStatus\":\"1\",\"createTime\":\"2017-08-10 00:00:00\",\"integral\":0.0,\"certiNo\":\"350123200\",\"userType\":\"0\"},\"respCode\":\"01\",\"errorMsg\":\"成功\"}\n")
	// 加密
	res, err := AesEncryptECB(str, key)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(crypt.Sha1(strings.ToUpper(hex.EncodeToString(res)))) // 输出：668c826342b8703d86e8bbf404610499

	// 解密
	de, err := AesDecryptECB(res, key)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(de))
}

func Test_Cryptos(t *testing.T) {
	encodingString := PswEncrypt("test123")
	decodingString := PswDecrypt(encodingString)
	fmt.Printf("AES-128-CBC\n加密：%s\n解密：%s\n", encodingString, decodingString)
}

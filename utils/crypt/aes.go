package crypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
)

// PKCS5Padding 填充明文
func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS5UnPadding 去除填充数据
func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

//AesEncrypt AES加密
// key的长度必须是16、24或者32字节，分别用于选择AES-128, AES-192, or AES-256
func AesEncrypt(origData, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	//AES分组长度为128位，所以blockSize=16，单位字节
	blockSize := block.BlockSize()
	origData = pkcs5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize]) //初始向量的长度必须等于块block的长度16字节
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	bas64Crypted := base64.StdEncoding.EncodeToString(crypted)
	return bas64Crypted, nil
}

func AesEncryptCBC(origData []byte, key []byte) (encrypted []byte) {
	// NewCipher该函数限制了输入k的长度必须为16, 24或者32
	key, _ = AesSha1prng(key, 128) // 比示例一多出这一步
	block, _ := aes.NewCipher(generateKey(key))
	blockSize := block.BlockSize()                              // 获取秘钥块的长度
	origData = pkcs5Padding(origData, blockSize)                // 补全码
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize]) // 加密模式
	encrypted = make([]byte, len(origData))                     // 创建数组
	blockMode.CryptBlocks(encrypted, origData)                  // 加密
	return encrypted
}
func AesDecryptCBC(encrypted []byte, key []byte) (decrypted []byte) {
	key, _ = AesSha1prng(key, 128)                              // 比示例一多出这一步
	block, _ := aes.NewCipher(generateKey(key))                 // 分组秘钥
	blockSize := block.BlockSize()                              // 获取秘钥块的长度
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize]) // 加密模式
	decrypted = make([]byte, len(encrypted))                    // 创建数组
	blockMode.CryptBlocks(decrypted, encrypted)                 // 解密
	decrypted = pkcs5UnPadding(decrypted)                       // 去除补全码
	return decrypted
}

// AesDecrypt AES解密
// key的长度必须是16、24或者32字节，分别用于选择AES-128, AES-192, or AES-256
func AesDecrypt(crypted64 string, key []byte) ([]byte, error) {
	crypted, err := base64.StdEncoding.DecodeString(crypted64)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	//AES分组长度为128位，所以blockSize=16，单位字节
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize]) //初始向量的长度必须等于块block的长度16字节
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = pkcs5UnPadding(origData)
	return origData, nil
}

// AesDecryptECB 适用于Java的解密
func AesDecryptECB(encrypted []byte, key []byte) ([]byte, error) {
	key, err := AesSha1prng(key, 128) // 比示例一多出这一步
	if err != nil {
		return nil, err
	}
	cipher, _ := aes.NewCipher(generateKey(key))
	decrypted := make([]byte, len(encrypted))
	//
	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}
	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}
	return decrypted[:trim], nil
}

// AesEncryptECB 适用于Java的加密
func AesEncryptECB(src []byte, key []byte) ([]byte, error) {
	key, err := AesSha1prng(key, 128) // 比示例一多出这一步
	if err != nil {
		return nil, err
	}

	cipher, _ := aes.NewCipher(generateKey(key))
	length := (len(src) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, src)
	pad := byte(len(plain) - len(src))
	for i := len(src); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted := make([]byte, len(plain))
	// 分组分块加密
	for bs, be := 0, cipher.BlockSize(); bs <= len(src); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}
	fmt.Println("2", Sha1(string(encrypted)))
	return encrypted, nil
}

// AesSha1prng 模拟 java SHA1PRNG 处理
func AesSha1prng(keyBytes []byte, encryptLength int) ([]byte, error) {
	hashs := Sha1Byte(Sha1Byte(keyBytes))
	maxLen := len(hashs)
	realLen := encryptLength / 8
	if realLen > maxLen {
		return nil, errors.New("invalid length!")
	}
	return hashs[0:realLen], nil
}

func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	fmt.Println("k2", Sha1(string(genKey)))
	return genKey
}

func AesKeySecureRandom(keyword string) (key []byte) {
	key, _ = AesSha1prng(key, 128)
	return generateKey(key)
}

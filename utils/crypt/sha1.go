package crypt

import (
	"crypto/sha1"
	"encoding/hex"
)

func Sha1(s string) string {
	o := sha1.New()
	o.Write([]byte(s))
	return hex.EncodeToString(o.Sum(nil))
}

func Sha1Byte(data []byte) []byte {
	h := sha1.New()
	h.Write(data)
	return h.Sum(nil)
}

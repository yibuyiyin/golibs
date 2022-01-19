package utils

import (
	"math/rand"
	"time"
)

const (
	RandDigit  = "01"
	RandLetter = "02"
	RandMix    = "03"
)

func Rand(length int, t string) string {
	var s string
	switch t {
	case RandMix:
		s = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	case RandDigit:
		s = "0123456789"
	case RandLetter:
		s = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	}
	l := len(s)
	rand.Seed(time.Now().UnixNano())
	str := ""
	for i := 0; i < length; i++ {
		rands := rand.Intn(l)
		str += string(s[rands])
	}
	return str
}

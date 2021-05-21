package utils

import (
	"math/rand"
	"time"
)

func Rand(lenght int) string {
	s := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	l := len(s)
	rand.Seed(time.Now().UnixNano())
	str := ""
	for i := 0; i < lenght; i++ {
		rands := rand.Intn(l)
		str += string(s[rands])
	}
	return str
}

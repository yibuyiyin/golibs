package aes

import (
	"testing"
)

func TestAesECB(t *testing.T) {
	key := []byte("12345678")
	str := "{\"hello\", \"world\"}"
	en, err := JavaEncryptECB([]byte(str), key)
	if err != nil {
		t.Log(err)
	}
	t.Log(string(en))
	de, err := JavaDecryptECB(en, key)
	if err != nil {
		t.Log(err)
	}
	t.Log(string(de))
}

func TestAesCBC(t *testing.T) {
	key := []byte("12345678")
	str := "{\"hello\", \"world\"}"
	en, err := JavaEncryptCBC([]byte(str), key)
	if err != nil {
		t.Log(err)
	}
	t.Log(string(en))
	de, err := JavaDecryptCBC(en, key)
	if err != nil {
		t.Log(err)
	}
	t.Log(string(de))
}

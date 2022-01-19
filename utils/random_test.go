package utils

import (
	"testing"
)

func TestRand(t *testing.T) {
	t.Log(Rand(32, RandMix))
	t.Log(Rand(32, RandDigit))
	t.Log(Rand(32, RandLetter))
}

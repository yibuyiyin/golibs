package utils

import (
	"testing"
)

func TestRand(t *testing.T) {
	t.Run("常规测试", func(t *testing.T) {
		t.Parallel()
		s := Rand(32, RandMix)
		if len(s) != 32 {
			t.Error("预期不符")
		}
	})
}

func BenchmarkRand(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var s, s1 string
		for pb.Next() {
			s = Rand(16, RandMix)
			if s == s1 {
				b.Error("预期不符")
			}
			s1 = s
		}
	})
}

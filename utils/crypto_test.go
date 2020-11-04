package utils

import (
	"testing"
	"unicode/utf8"
)

func TestRandomString(t *testing.T) {
	for i := 0; i < 32; i += 1 {
		s := RandomString(i)
		t.Logf("len(%d), %s, %d", i, s, utf8.RuneCountInString(s))
		if i != utf8.RuneCountInString(s) {
			t.Error("error len")
		}
	}
}

func TestSha256s(t *testing.T) {
	s := Sha256s(RandomString(500))
	t.Logf("%d", utf8.RuneCountInString(s))
	if utf8.RuneCountInString(s) >= 100 {
		t.Error("to long")
	}
}

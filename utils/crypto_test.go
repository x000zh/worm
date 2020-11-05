package utils

import (
	"math/rand"
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

func TestEncryptDecryptInt64(t *testing.T) {
	var i int64
	//key := []byte("bac0cd32ad54d0e10665ca8593ec9e2e")
	key := []byte("bac0cd32ad54d0e1")
	for i = -5; i < 10; i += 1 {
		s, err := EncryptInt64(i, key)
		if nil != err {
			t.Error(err)
			break
		}
		t.Logf("in %d %d, out %s, len(%d)", i, uint64(i), s, len(s))
		o, err := DecryptInt64(s, key)
		if nil != err {
			t.Error(err)
			break
		}
		if o != i {
			t.Errorf("in %d out: %s, back %d", i, s, o)
			break
		}
	}
}

func TestEncodeDecodeInt64(t *testing.T) {
	//key := []byte("bac0cd32ad54d0e10665ca8593ec9e2e")
	key := uint64(0x6ca3bb584523124c)
	for j := int64(-5); j < 10; j += 1 {
		i := j
		if j > 5 {
			i = rand.Int63()
		}
		s := EncodeInt64(i, key)
		t.Logf("in %d %d, out %s, len(%d)", i, uint64(i), s, len(s))
		o, err := DecodeInt64(s, key)
		if nil != err {
			t.Error(err)
			break
		}
		if o != i {
			t.Errorf("in %d out: %s, back %d", i, s, o)
			break
		}
	}
}

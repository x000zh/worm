package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base32"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"unicode/utf8"
)

var ivMask, _ = hex.DecodeString("217c4c2faeca17a10af2bce633b5e573298a953bb0b765a9b2b3052635fc2d51")
var ivMaskLen = len(ivMask)
var base32Encoder = "kfqpl6r935vnae1bs8cwxyi7gdthj42u"
var base32EncoderPadding = '0'

//随机生成字符串
func RandomString(len int) string {
	if len < 1 {
		return ""
	}
	token := make([]byte, (len*6)/8+1)
	rand.Read(token)
	s := base64.RawStdEncoding.EncodeToString(token)
	if utf8.RuneCountInString(s) > len {
		return s[0:len]
	}
	return s
}

func Sha1s(input string) string {
	payload := sha1.Sum([]byte(input))
	return base64.RawStdEncoding.EncodeToString(payload[:])
}

func Sha256s(input string) string {
	payload := sha256.Sum256([]byte(input))
	return base64.RawStdEncoding.EncodeToString(payload[:])
}

//key []byte 32/24/16 bytes
func EncryptInt64(i int64, key []byte) (string, error) {
	c, err := aes.NewCipher(key)
	if nil != err {
		return "", err
	}
	gcm, err := cipher.NewGCM(c)
	if nil != err {
		return "", err
	}
	nonce := makeNonce(key, gcm.NonceSize())
	val := make([]byte, 8)
	binary.LittleEndian.PutUint64(val, uint64(i))
	result := gcm.Seal(nil, nonce, val, nil)
	base32Encoding := base32.NewEncoding(base32Encoder).WithPadding(base32EncoderPadding)
	return base32Encoding.EncodeToString(result), nil
}

func DecryptInt64(i string, key []byte) (int64, error) {

	base32Encoding := base32.NewEncoding(base32Encoder).WithPadding(base32EncoderPadding)
	data, err := base32Encoding.DecodeString(i)
	if nil != err {
		return 0, err
	}

	c, err := aes.NewCipher(key)
	if nil != err {
		return 0, err
	}
	gcm, err := cipher.NewGCM(c)
	if nil != err {
		return 0, err
	}

	nonce := makeNonce(key, gcm.NonceSize())
	result, err := gcm.Open(nil, nonce, data, nil)
	if nil != err {
		return 0, err
	}
	//判断是否是64位
	if len(result) != 8 {
		return 0, fmt.Errorf("invalid input:%s", i)
	}
	ui := binary.LittleEndian.Uint64(result)
	return int64(ui), nil
}

var bytesOrder = []int{6, 5, 8, 7, 4}

//var bytesOrder = []int{6, 8, 7, 5}

//简单加密
func EncodeInt64(i int64, key uint64) string {
	a := uint64(i) ^ key
	val := make([]byte, 10)
	binary.LittleEndian.PutUint64(val, uint64(a))
	cs := Crc16(val[:8])
	binary.LittleEndian.PutUint16(val[8:], cs)

	base32Encoding := base32.NewEncoding(base32Encoder).WithPadding(base32EncoderPadding)
	for i, j := range bytesOrder {
		temp := val[i]
		val[i] = val[j]
		val[j] = temp
	}
	return base32Encoding.EncodeToString(val)
}

func DecodeInt64(i string, key uint64) (int64, error) {
	base32Encoding := base32.NewEncoding(base32Encoder).WithPadding(base32EncoderPadding)
	val, err := base32Encoding.DecodeString(i)
	if nil != err {
		return 0, err
	}
	for i, j := range bytesOrder {
		temp := val[i]
		val[i] = val[j]
		val[j] = temp
	}
	result := binary.LittleEndian.Uint64(val[:8])
	result = uint64(result) ^ key
	cs := Crc16(val[:8])
	if val[9] != byte(cs>>8) || val[8] != byte(0x00ff&cs) {
		return 0, fmt.Errorf("validate failed")
	}
	return int64(result), nil
}

func makeNonce(key []byte, size int) []byte {
	nonce := make([]byte, size)
	for i := 0; i < size; i += 1 {
		mod := i % ivMaskLen
		if len(key) > i {
			nonce[i] = ivMask[mod] ^ key[i]
		} else {
			nonce[i] = ivMask[mod]
		}
	}
	return nonce
}

var crc16tab = [256]uint16{
	0x0000, 0x1021, 0x2042, 0x3063, 0x4084, 0x50a5, 0x60c6, 0x70e7,
	0x8108, 0x9129, 0xa14a, 0xb16b, 0xc18c, 0xd1ad, 0xe1ce, 0xf1ef,
	0x1231, 0x0210, 0x3273, 0x2252, 0x52b5, 0x4294, 0x72f7, 0x62d6,
	0x9339, 0x8318, 0xb37b, 0xa35a, 0xd3bd, 0xc39c, 0xf3ff, 0xe3de,
	0x2462, 0x3443, 0x0420, 0x1401, 0x64e6, 0x74c7, 0x44a4, 0x5485,
	0xa56a, 0xb54b, 0x8528, 0x9509, 0xe5ee, 0xf5cf, 0xc5ac, 0xd58d,
	0x3653, 0x2672, 0x1611, 0x0630, 0x76d7, 0x66f6, 0x5695, 0x46b4,
	0xb75b, 0xa77a, 0x9719, 0x8738, 0xf7df, 0xe7fe, 0xd79d, 0xc7bc,
	0x48c4, 0x58e5, 0x6886, 0x78a7, 0x0840, 0x1861, 0x2802, 0x3823,
	0xc9cc, 0xd9ed, 0xe98e, 0xf9af, 0x8948, 0x9969, 0xa90a, 0xb92b,
	0x5af5, 0x4ad4, 0x7ab7, 0x6a96, 0x1a71, 0x0a50, 0x3a33, 0x2a12,
	0xdbfd, 0xcbdc, 0xfbbf, 0xeb9e, 0x9b79, 0x8b58, 0xbb3b, 0xab1a,
	0x6ca6, 0x7c87, 0x4ce4, 0x5cc5, 0x2c22, 0x3c03, 0x0c60, 0x1c41,
	0xedae, 0xfd8f, 0xcdec, 0xddcd, 0xad2a, 0xbd0b, 0x8d68, 0x9d49,
	0x7e97, 0x6eb6, 0x5ed5, 0x4ef4, 0x3e13, 0x2e32, 0x1e51, 0x0e70,
	0xff9f, 0xefbe, 0xdfdd, 0xcffc, 0xbf1b, 0xaf3a, 0x9f59, 0x8f78,
	0x9188, 0x81a9, 0xb1ca, 0xa1eb, 0xd10c, 0xc12d, 0xf14e, 0xe16f,
	0x1080, 0x00a1, 0x30c2, 0x20e3, 0x5004, 0x4025, 0x7046, 0x6067,
	0x83b9, 0x9398, 0xa3fb, 0xb3da, 0xc33d, 0xd31c, 0xe37f, 0xf35e,
	0x02b1, 0x1290, 0x22f3, 0x32d2, 0x4235, 0x5214, 0x6277, 0x7256,
	0xb5ea, 0xa5cb, 0x95a8, 0x8589, 0xf56e, 0xe54f, 0xd52c, 0xc50d,
	0x34e2, 0x24c3, 0x14a0, 0x0481, 0x7466, 0x6447, 0x5424, 0x4405,
	0xa7db, 0xb7fa, 0x8799, 0x97b8, 0xe75f, 0xf77e, 0xc71d, 0xd73c,
	0x26d3, 0x36f2, 0x0691, 0x16b0, 0x6657, 0x7676, 0x4615, 0x5634,
	0xd94c, 0xc96d, 0xf90e, 0xe92f, 0x99c8, 0x89e9, 0xb98a, 0xa9ab,
	0x5844, 0x4865, 0x7806, 0x6827, 0x18c0, 0x08e1, 0x3882, 0x28a3,
	0xcb7d, 0xdb5c, 0xeb3f, 0xfb1e, 0x8bf9, 0x9bd8, 0xabbb, 0xbb9a,
	0x4a75, 0x5a54, 0x6a37, 0x7a16, 0x0af1, 0x1ad0, 0x2ab3, 0x3a92,
	0xfd2e, 0xed0f, 0xdd6c, 0xcd4d, 0xbdaa, 0xad8b, 0x9de8, 0x8dc9,
	0x7c26, 0x6c07, 0x5c64, 0x4c45, 0x3ca2, 0x2c83, 0x1ce0, 0x0cc1,
	0xef1f, 0xff3e, 0xcf5d, 0xdf7c, 0xaf9b, 0xbfba, 0x8fd9, 0x9ff8,
	0x6e17, 0x7e36, 0x4e55, 0x5e74, 0x2e93, 0x3eb2, 0x0ed1, 0x1ef0}

func Crc16(bs []byte) (crc uint16) {
	l := len(bs)
	for i := 0; i < l; i++ {
		crc = ((crc << 8) & 0xff00) ^ crc16tab[((crc>>8)&0xff)^uint16(bs[i])]
	}

	return
}


//
func SubCipher(chars []rune, n int) [][]rune {
	ret := make([][]rune, n)
	cLen := len(chars)
	for i:=0; i<n; i+=1 {
		ret[i] = make([]rune, cLen)
		for j:=0; j<cLen; j+=1 {
			ret[i][j] = chars[j]
		}
		rand.Shuffle(cLen, func(j, k int) {
			t := ret[i][k]
			ret[i][k] = ret[i][j]
			ret[i][j] = t
		})
	}

	return ret
}

func SubCipherEncode(cipher [][]rune, input int, min int) string {
	pattern := fmt.Sprintf("%%0%dd", min)
	s := fmt.Sprintf(pattern, input)
	l := len(s)
	
	ret := make([]string, l)
	cipherLen := len(cipher)
	j := 0
	for i:=l-1; i>=0; i-=1 {
		if i < cipherLen {
			ret[j] = fmt.Sprintf("%c", cipher[j][s[i] % '0'])
		} else {
			ret[j] = fmt.Sprintf("%c", s[i])
		}
		j += 1
	}

	out, _ := JoinSlice(ret, "")
	return out
}


func SubCipherDecode(cipher [][]rune, s string, min int) int {
	l := len(s)
	if l < min {
		l = min
	}
	cipherLen := len(cipher)
	ret := make([]rune, l)
	j := l-1
	for i:=0; i<l; i+=1 {
		idx := 0
		for ii, r := range cipher[i] {
			if r == rune(s[i]) {
				idx = ii
			}
		}
		if i < cipherLen {
			ret[j] = rune(idx) + '0'
		} else {
			ret[j] = rune(s[i])
		}
		j -= 1
	}

	out, _ := strconv.Atoi(JoinRunes(ret, ""))
	return out
}



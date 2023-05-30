package wutils

import "unsafe"


func BytesToString(ss []byte) string {
	return *(*string)(unsafe.Pointer(&ss))
}


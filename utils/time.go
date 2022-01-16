package utils

import (
	"time"
)


func GetCurrentTimeStamp() int64 {
	t := time.Now()
	return t.Unix()
}

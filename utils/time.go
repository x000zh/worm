package utils

import (
	"strconv"
	"time"
)


func GetCurrentTimeStamp() int64 {
	t := time.Now()
	return t.Unix()
}


//当前时间20060102
func GetCurrentDT() int32 {
	t := time.Now()
	dt := t.Format("20060102")
	i, _ := strconv.Atoi(dt)
	return int32(i)
}


func GetDT(t time.Time) int32 {
	dt := t.Format("20060102")
	i, _ := strconv.Atoi(dt)
	return int32(i)
}

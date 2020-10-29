package utils

import (
	"testing"
)

type testGetValuesOfSliceStruct struct {
	A string
	B int64
}

func TestGetValuesOfSlice(t *testing.T) {
	s1 := testGetValuesOfSliceStruct{
		A: "a1",
		B: 1,
	}

	s2 := testGetValuesOfSliceStruct{
		A: "a2",
		B: 2,
	}
	slice := make([]testGetValuesOfSliceStruct, 2)
	slice[0] = s1
	slice[1] = s2

	i, err := GetValuesOfSlice(slice, "A")
	if nil != err {
		t.Error(err)
		return
	}
	a := i.([]string)
	i, err = GetValuesOfSlice(slice, "B")
	if nil != err {
		t.Error(err)
		return
	}
	b := i.([]int64)
	if len(a) != 2 {
		t.Error("run failed", slice, a)
		return
	}
	if len(b) != 2 {
		t.Error("run failed", slice, b)
		return
	}
	if a[0] != "a1" {
		t.Error("run failed", slice, a)
		return
	}
	if a[1] != "a2" {
		t.Error("run failed", slice, a)
		return
	}
	if b[1] != 2 {
		t.Error("run failed", slice, b)
		return
	}

	slice2 := make([]testGetValuesOfSliceStruct, 0)
	i, err = GetValuesOfSlice(slice2, "B")
	if nil != err {
		t.Error(err)
		return
	}
	c := i.([]int64)
	if len(c) != 0 {
		t.Error("error", c)
		return
	}
}

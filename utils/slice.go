package utils

import (
	"fmt"
	"strconv"
	"strings"

	//"math/rand"
	//"net"
	"reflect"

	"github.com/pkg/errors"
	//"time"
)

//IntefaceSlice - 通过反射
func IntefaceSlice(v interface{}) []interface{} {
	items := reflect.ValueOf(v)
	if items.Kind() == reflect.Slice {
		ret := make([]interface{}, 0, items.Len())
		length := items.Len()
		for i := 0; i < length; i++ {
			ret = append(ret, items.Index(i).Interface())
		}
		return ret
	}
	ret := make([]interface{}, 1)
	ret[0] = v
	return ret
}


//GetValuesOfSlice - 将对象数组的值塞到另一个数组里
func GetValuesOfSlice(v interface{}, key string) (interface{}, error) {
	items := reflect.ValueOf(v)
	if items.Kind() == reflect.Slice {
		length := items.Len()
		//sliceType := toReflect.Type().Elem()
		//ret := reflect.Value{}
		elemType := reflect.TypeOf(v).Elem()
		value, ok := elemType.FieldByName(key)
		if !ok {
			return nil, fmt.Errorf("invalid field")
		}
		ret := reflect.MakeSlice(reflect.SliceOf(value.Type), 0, length)
		for i := 0; i < length; i++ {
			v := items.Index(i)
			if !v.CanAddr() {
				return nil, fmt.Errorf("can't not addr slice")
			}
			//item := v.Interface()
			value := v.FieldByName(key)
			ret = reflect.Append(ret, value)
		}
		return ret.Interface(), nil
	}
	return nil, fmt.Errorf("invald type")
}


func JoinSlice(arr interface{}, sep string) (ret string, err error) {
	r := reflect.ValueOf(arr)
	if r.Kind() == reflect.Slice {
		length := r.Len()
		elemType := reflect.TypeOf(arr).Elem()
		sArr := make([]string, length)
		for i:= 0; i < length; i++ {
			v := r.Index(i)
			switch elemType.Kind() {
			case reflect.Int16,
				 reflect.Int32,
				 reflect.Int64,
			 	 reflect.Uint16,
				 reflect.Uint32,
				 reflect.Uint64,
				 reflect.Int,
				 reflect.Uint:
					sArr[i] = fmt.Sprintf("%d", v.Interface())
			case reflect.String:
					sArr[i] = v.String()
			default:
				sArr[i] = fmt.Sprintf("%v", v.Interface())
			}
		}
		ret = strings.Join(sArr, sep)
		return
	} else {
		err = errors.Errorf("invalid type %s, expect slice", r.Kind().String())
	}
	return
}


func JoinRunes(arr []rune, sep string) string {
	l := len(arr)
	ret := make([]string, l)
	for i:=0; i<l; i+=1 {
		ret[i] = fmt.Sprintf("%c", arr[i])
	}
	return strings.Join(ret, sep)
}


func IntSlice(arr []string) []int {
	ret := make([]int, len(arr))
	if (len(ret) == 0) {
		return ret
	}

	for i, s := range arr {
		ret[i], _ = strconv.Atoi(s)
	}
	return ret
}

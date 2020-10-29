package utils

import (
	"fmt"
	//"math/rand"
	//"net"
	"reflect"
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

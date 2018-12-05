package utils

import (
	"reflect"
)

// SliceIndex  查找 slice 中元素
// 返回 元素在 slice 中位置, -1 为未找到, -2 为发生错误
func SliceIndex(arr interface{}, it interface{}) int {
	if arr == nil || it == nil {
		return -1
	}

	v := reflect.ValueOf(arr)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Slice {
		return -2
	}

	arrLen := v.Len()
	for i := 0; i < arrLen; i++ {
		theV := v.Index(i).Interface()

		if reflect.DeepEqual(theV, it) {
			return i
		}
	}

	return -1
}

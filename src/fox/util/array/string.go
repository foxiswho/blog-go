package array

import (
	"fmt"
	"fox/util"
)
//slice翻转
func StringReverse(src []string) error {
	if src == nil {
		panic(fmt.Errorf("the src can't be empty!"))
		return &util.Error{Msg:"数组不能为空"}
	}
	count := len(src)
	if count < 1 {
		return &util.Error{Msg:"数组不能为空"}
	}
	mid := count / 2
	for i := 0; i < mid; i++ {
		tmp := src[i]
		src[i] = src[count - 1]
		src[count - 1] = tmp
		count--
	}
	return nil
}
//判断是否包含
func SliceContains(src []string, value string) bool {
	isContain := false
	for _, srcValue := range src {
		if (srcValue == value) {
			isContain = true
			break
		}
	}
	return isContain
}

//判断key是否存在
func MapContains(src map[string]int, key string) bool {
	if _, ok := src[key]; ok {
		return true
	}
	return false
}
package array

import "fmt"
//slice翻转
func StringReverse(src []string){
	if src == nil {
		panic(fmt.Errorf("the src can't be empty!"))
	}
	count := len(src)
	mid := count/2
	for i := 0;i < mid; i++{
		tmp := src[i]
		src[i] = src[count-1]
		src[count-1] = tmp
		count--
	}
}
//判断是否包含
func SliceContains(src []string,value string)bool{
	isContain := false
	for _,srcValue := range src  {
		if(srcValue == value){
			isContain = true
			break
		}
	}
	return isContain
}

//判断key是否存在
func MapContains(src map[string]int ,key string) bool{
	if _, ok := src[key]; ok {
		return true
	}
	return false
}
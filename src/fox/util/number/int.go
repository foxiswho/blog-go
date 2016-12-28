package number

import "strconv"

//此处只记录 转换方式，并不使用
//
//数字变成字符串
func str(i int) string {
	return strconv.Itoa(i)
}
//数字变成字符串
func strFormInt64(i int64) string {
	return strconv.FormatInt(i,10)
}
//数字变成字符串
func strFormInterface(i interface{}) string {
	return strconv.Itoa(i.(int))
}
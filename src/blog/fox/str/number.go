package str

import "strconv"

//字符串转换成  int 没有错误返回
func Int(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}
//字符串转换成  int64 没有错误返回
func Int64(str string) int64 {
	i, _ := strconv.ParseInt(str,10,64)
	return i
}
//字符串转换成  float64 没有错误返回
func Float64(str string) float64 {
	i, _ := strconv.ParseFloat(str,64)
	return i
}
//字符串转换成  float64 没有错误返回
func Float64FormInterface(str interface{}) float64 {
	i, _ := strconv.ParseFloat(str.(string),64)
	return i
}
//字符串转换成  int 没有错误返回
func IntFormInterface(str interface{}) int {
	i, _ := strconv.Atoi(str.(string))
	return i
}
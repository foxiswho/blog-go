package str

import "fmt"

//截取字符串 start 起点下标 end 终点下标(不包括)
func Substr(str string, start int, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		fmt.Println("Substr error: start is wrong")
		return str
		//panic("start is wrong")
	}

	if end < 0 || end > length {
		fmt.Println("Substr error: end is wrong")
		return str
		//panic("end is wrong")
	}

	return string(rs[start:end])
}
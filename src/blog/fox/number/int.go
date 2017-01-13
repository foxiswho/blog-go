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
	return strconv.FormatInt(i, 10)
}
//数字变成字符串
func strFormInterface(i interface{}) string {
	return strconv.Itoa(i.(int))
}
//onj变成数字
func ObjToInt(i interface{}) (int, error) {
	n := 0
	switch i.(type) {
	case int:
		n = i.(int)
	case int32:
		n = int(i.(int32))
	case int64:
		n = int(i.(int64))
	case float32:
		n = int(i.(float32))
	case float64:
		n = int(i.(float64))
	case string:
		var err error
		n, err = strconv.Atoi(i.(string))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}
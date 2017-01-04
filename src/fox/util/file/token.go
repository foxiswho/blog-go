package file

import "fox/util/str"
//令牌生成
func TokeMake(maps map[string]interface{}) string {
	return str.JsonEnCode(maps)
}
//func TokenDeCode(str string)  {
//	return
//}
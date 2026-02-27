package strPg2

// ArrayToJsonExpr
//
//	@Description:  数组转为 json 搜索条件
//	@param str
//	@return string
func ArrayToJsonExpr(str []string) string {
	if len(str) == 0 {
		return "[]"
	}
	val := ""
	for _, v := range str {
		val += `"` + v + `",`
	}
	val = val[:len(val)-1]
	return "[" + val + "]"
}

// StrToArrayJsonExpr
//
//	@Description:  字符转 数组json 搜索条件
//	@param str
//	@return string
func StrToArrayJsonExpr(str string) string {
	return `["` + str + `"]`
}

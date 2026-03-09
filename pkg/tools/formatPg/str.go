package formatPg

import "regexp"

// 预编译正则表达式（全局只编译一次，提升性能）
var validStrRegex = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]*$`)

// ValidateString 校验字符串是否符合规则：
// 1. 仅包含大小写字母、下划线、数字
// 2. 首字符必须是英文字母
func ValidateString(s string) bool {
	// 先判断空字符串（空字符串不符合规则）
	if s == "" {
		return false
	}
	// 使用正则匹配
	return validStrRegex.MatchString(s)
}

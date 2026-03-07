package timePg

import "time"

// IsValidYearMonthFlexible 兼容 "YYYY-M" 或 "YYYY-MM" 格式
func IsValidYearMonthFlexible(input string) (time.Time, bool) {
	// 先尝试解析严格格式
	layout := "2006-01"
	t, err := time.Parse(layout, input)
	if err == nil {
		return t, true
	}

	// 再尝试解析宽松格式（月份无前置零）
	layoutFlex := "2006-1"
	t, err = time.Parse(layoutFlex, input)
	if err != nil {
		return t, false
	}
	return t, true
}

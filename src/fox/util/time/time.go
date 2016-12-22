package time

import (
	"time"
	"fmt"
)

const (
	Y_M = "2006-01"
	Y_M_D = "2006-01-02"
	Y_M_D_2 = "2006年01月02日"
	Y_M_D_H_I_S = "2006-01-02 15:04:05"
	Y_M_D_H_I_S_2 = "2006年01月02日 15:04:05"
	H_I_S = "15:04:05"
)

type DateTime time.Time

//func (t DateTime) MarshalJSON() ([]byte, error) {
//	//do your serializing here
//	stamp := fmt.Sprintf("\"%s\"", t.DateTime())
//	return []byte(stamp), nil
//}
////日期时间
//func (t DateTime) DateTime() string {
//	return t.Format(Y_M_D_H_I_S)
//}
////日期
//func (t DateTime) Date() string {
//	return t.Format(Y_M_D)
//}
////时间
//func (t DateTime) Time() string {
//	return t.Format(H_I_S)
//}
//func (t DateTime) Format(layout string) string {
//	return time.Now().Format(layout)
//}
//格式
func Format(str interface{}, layout string) string {
	var date time.Time
	var err error
	//判断变量类型
	switch str.(type) {
	case string:
		//如果是字符串则转换成 标准日期时间格式
		fmt.Println(str)
		date, err = time.Parse(layout, str.(string))
		if err != nil {
			return ""
		}
	case time.Time:
		date = str
	}

	return date.Format(layout)
}
//当前日期时间
func Now() string {
	return Format(time.Now(), Y_M_D_H_I_S)
}
//当前日期时间
func DateTime() string {
	return Format(time.Now(), Y_M_D_H_I_S)
}
//当前日期
func Date() string {
	return Format(time.Now(), Y_M_D)
}
//当前时间
func Time() string {
	return Format(time.Now(), H_I_S)
}
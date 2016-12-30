package main

import (
	"fmt"
	"reflect"
	"strings"
	"encoding/json"
)

func main() {
	xx := []string{"ss", "sd", "dasd"}
	fmt.Println("数组长度", len(xx))
	fmt.Println("数组长度", cap(xx))
	where := make(map[string]interface{})
	where["id"] = 199
	where["id=?"] = 15
	where["id>=?"] = "ddd"
	where["id in (?)"] = []int{1, 2, 3}
	where["id in (?)"] = []string{"ss", "sd", "dasd"}
	where["id between ? and ? "] = []int{1, 2, 4}
	fmt.Println("i     v ")
	str := ""
	for k, v := range where {
		fmt.Println(k, v, reflect.TypeOf(v))
		fmt.Println("?号个数为", strings.Count(k, "?"))
		QuestionMarkCount := strings.Count(k, "?")
		isEmpty := false
		isMap := false
		switch v.(type) {
		case string:
			//是字符时做的事情
			isEmpty = v == ""
		case int:
		//是整数时做的事情
		case []string :
			isMap = true
			isEmpty = len(v) == 0
		case []int :
			isMap = true
			isEmpty = len(v) == 0
		}
		if QuestionMarkCount == 0 && isEmpty {
			str += " AND " + k + " = '' "
		} else if QuestionMarkCount == 0 && !isEmpty {
			//是数组
			if (isMap) {
				str += " AND " + k + " = " + JsonEnCode(v)
			} else {
				//不是数组
				str += " AND " + k + " = " + v
			}
		} else if QuestionMarkCount == 1 && isEmpty {
			//值为空字符串,不是数组
			str += " AND " + k + " = " + v
		} else if QuestionMarkCount == 1 && !isEmpty {
			//是数组
			if isMap {
				str += " AND " + k + " = " + JsonEnCode(v)
			} else {
				//不是数组
				//不是数组，有值
				str += " AND " + k + " = '' "
			}
		} else if QuestionMarkCount > 1 && isEmpty {
			//不是数组，空值
			str += " AND " + k + " = ''"
		} else if QuestionMarkCount > 1 && !isEmpty & isMap {
			count := len(v)
			//问号 与  数组相同时
			if QuestionMarkCount == count {
				//不是数组，空值
				str += " AND " + k + " = ''"
			}else{
				//问号 与  数组不同时
				str += " AND " + k + " = ''"
			}
		}else {
			fmt.Println("其他还没有收录")
		}
	}
}
func JsonEnCode(v interface{}) string {
	str, err := json.Marshal(v)
	if err != nil {
		fmt.Println("序列化失败:", err)
	}
	return string(str)
}
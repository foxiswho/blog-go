package array

import (
	"encoding/json"
	"blog/util"
)


// 函　数：Obj2map
// 概　要：
// 参　数：
//      obj: 传入Obj
// 返回值：
//      mapObj: map对象
//      err: 错误
func ObjToMap(obj interface{}) (mapObj map[string]interface{}, err error) {
	// 结构体转json
	b, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(b, &result); err != nil {
		return nil, err
	}
	return result, nil
}
//JSON格式数据转换为map
func StrToMap(str string) (mapObj map[string]interface{}, err error) {
	// 结构体转json
	if str == "" {
		return nil, &util.Error{Msg:"字符串为空不能转换"}
	}
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(str), &result); err != nil {
		return nil, err
	}
	return result, nil
}
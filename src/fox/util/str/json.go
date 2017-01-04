package str

import (
	"encoding/json"
	"fmt"
	"fox/util"
)

func JsonEnCode(v interface{}) (string,error) {
	str, err := json.Marshal(v)
	if err != nil {
		fmt.Println("序列化失败:", err)
		return "",&util.Error{Msg:"序列化失败:"+err.Error()}
	}
	return string(str),nil
}
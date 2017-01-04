package str

import (
	"encoding/json"
	"fmt"
)

func JsonEnCode(v interface{}) string {
	str, err := json.Marshal(v)
	if err != nil {
		fmt.Println("序列化失败:", err)
		return ""
	}
	return string(str)
}
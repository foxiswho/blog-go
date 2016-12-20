package Response

import (
	"fmt"
	"encoding/json"
	"net/http"
)
const (
	WARNING = "warning"
	SUCCESS = "success"
	ALERT   = "alert"
	INFO    = "info"
)

type Response struct {
	Code int
	Data   interface{}
	Err    Error
}
type Error struct {
	Level string
	Msg   string
}
type Success struct {
	Level string
	Msg   string
}

func NewResponse() *Response {
	return &Response{Code: 1}
}
func (resp *Response) Tips(msg,level string) {
	resp.Err = Error{level, msg}
}
func (resp *Response) WriteJson(w http.ResponseWriter) {
	b, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("helper.go line:55", err)
		w.Write([]byte(`{Status:-1,Err:Error{Level:"alert",Msg:"code=-1|序列化失败！"}}`))
	} else {
		w.Write(b)
	}
}
func (resp *Response) Success() {
	resp.Code = 1
	resp.Data = Success{Level: SUCCESS, Msg: "恭喜(●'◡'●)|操作成功。"}
}
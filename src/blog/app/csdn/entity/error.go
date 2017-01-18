package entity
//错误
type Error struct {
	Request string `json:"request"`
	ErrorCode string `json:"error_code"`
	Error string `json:"error"`
}
//初始化
func NewError()*Error{
	return  new(Error)
}
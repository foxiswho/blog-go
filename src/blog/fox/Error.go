package fox
//此处 为以后 更换框架做准备
//错误基类
type Error struct {
	Msg string
}

func (e *Error) Error() string {
	return e.Msg
}
func NewError(msg string) *Error {
	e := new(Error)
	e.Msg = msg
	return e
}
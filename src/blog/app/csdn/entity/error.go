package entity

type Error struct {
	Request string `json:"request"`
	ErrorCode string `json:"error_code"`
	Error string `json:"error"`
}
func NewError()*Error{
	return  new(Error)
}
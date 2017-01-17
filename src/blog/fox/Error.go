package fox

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
package fox


type Error struct {
	Msg string
}

func (e *Error) Error() string {
	return e.Msg
}
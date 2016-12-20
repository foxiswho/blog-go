package Error


type Error struct {
	Mesage string
}

func (e *Error) Error() string {
	return e.Mesage
}
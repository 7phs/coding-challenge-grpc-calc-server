package calc

// An error uses in divide by zero.
type ErrZeroDiv struct {
	msg string
}

func NewErrZeroDiv(msg string) *ErrZeroDiv {
	return &ErrZeroDiv{
		msg: msg,
	}
}
func (e *ErrZeroDiv) Error() string {
	return e.msg
}
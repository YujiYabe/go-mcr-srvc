package controller

// Error ...
type Error struct {
	Message string
}

// NewError ...
func NewError(err error) *Error {
	return &Error{
		Message: err.Error(),
	}
}

package errors

type BusinessError struct {
	Code    string
	Message string
}

func (e *BusinessError) Error() string {
	return e.Message
}

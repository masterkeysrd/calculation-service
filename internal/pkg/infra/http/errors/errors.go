package errors

type HTTPError struct {
	code    int
	message string
}

func New(code int, message string) HTTPError {
	return HTTPError{
		code:    code,
		message: message,
	}
}

func (e HTTPError) StatusCode() int {
	return e.code
}

func (e HTTPError) Error() string {
	return e.message
}

package errors

import "net/http"

var (
	ErrUnauthorized = Unauthorized("unauthorized")
)

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

func BadRequest(message string) HTTPError {
	return New(http.StatusBadRequest, message)
}

func NotFound(message string) HTTPError {
	return New(http.StatusNotFound, message)
}

func InternalServerError(message string) HTTPError {
	return New(http.StatusInternalServerError, message)
}

func Conflict(message string) HTTPError {
	return New(http.StatusConflict, message)
}

func Forbidden(message string) HTTPError {
	return New(http.StatusForbidden, message)
}

func Unauthorized(message string) HTTPError {
	return New(http.StatusUnauthorized, message)
}

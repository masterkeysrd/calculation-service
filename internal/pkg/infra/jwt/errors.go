package jwt

import "errors"

var (
	ErrTokenIsNotProvided      = errors.New("bearer token is not provided")
	ErrUnexpectedSigningMethod = errors.New("unexpected signing method")
	ErrorInvalidToken          = errors.New("invalid token")
)

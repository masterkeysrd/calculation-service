package user

import "errors"

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrUserNameRequired   = errors.New("username is required")
	ErrInvalidCredentials = errors.New("username or password is invalid")
)

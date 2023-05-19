package auth

import "errors"

var (
	ErrInvalidUserName     = errors.New("invalid username")
	ErrInvalidCredentials  = errors.New("invalid credentials")
	ErrInvalidRefreshToken = errors.New("invalid refresh token")
)

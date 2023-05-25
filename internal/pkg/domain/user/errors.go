package user

import (
	"net/http"

	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/http/errors"
)

var (
	ErrUserNotFound        = errors.New(http.StatusNotFound, "user not found")
	ErrUserAlreadyExists   = errors.New(http.StatusConflict, "user already exists")
	ErrUserNameRequired    = errors.New(http.StatusBadRequest, "user name is required")
	ErrInvalidCredentials  = errors.New(http.StatusUnauthorized, "invalid credentials")
	ErrUserBalanceNotFound = errors.New(http.StatusNotFound, "user balance not found")
)

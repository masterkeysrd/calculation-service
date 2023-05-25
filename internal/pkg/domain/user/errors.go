package user

import (
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/http/errors"
)

var (
	ErrUserNotFound        = errors.NotFound("user not found")
	ErrUserAlreadyExists   = errors.Conflict("user already exists")
	ErrUserIDRequired      = errors.BadRequest("user id required")
	ErrInvalidCredentials  = errors.Unauthorized("invalid credentials")
	ErrUserBalanceNotFound = errors.NotFound("user balance not found")
)

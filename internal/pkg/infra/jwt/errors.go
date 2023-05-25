package jwt

import "github.com/masterkeysrd/calculation-service/internal/pkg/infra/http/errors"

var (
	ErrInvalidUserID           = errors.BadRequest("invalid user id")
	ErrorInvalidToken          = errors.Unauthorized("invalid token")
	ErrTokenIsNotProvided      = errors.Unauthorized("token is not provided")
	ErrUnexpectedSigningMethod = errors.InternalServerError("unexpected signing method")
)

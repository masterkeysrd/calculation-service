package calculation

import "github.com/masterkeysrd/calculation-service/internal/pkg/infra/http/errors"

var (
	ErrInvalidNumberOfArguments = errors.BadRequest("invalid number of arguments")
	ErrOperationNotSupported    = errors.BadRequest("operation not supported")
	ErrRandomDataEmpty          = errors.InternalServerError("generated random string data is empty")
)

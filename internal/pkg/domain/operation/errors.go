package operation

import (
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/http/errors"
)

var (
	ErrOperationNotFound = errors.NotFound("operation not found")
)

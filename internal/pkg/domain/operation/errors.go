package operation

import (
	"net/http"

	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/http/errors"
)

var (
	ErrOperationNotFound = errors.New(http.StatusNotFound, "Operation not found")
)

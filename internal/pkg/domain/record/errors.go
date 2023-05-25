package record

import (
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/http/errors"
)

var (
	ErrRecordNotFound = errors.NotFound("record not found")
)

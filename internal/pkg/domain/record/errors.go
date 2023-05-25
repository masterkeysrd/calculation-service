package record

import (
	"net/http"

	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/http/errors"
)

var (
	ErrRecordNotFound = errors.New(http.StatusNotFound, "Record not found")
)

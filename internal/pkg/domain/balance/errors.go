package balance

import "github.com/masterkeysrd/calculation-service/internal/pkg/infra/http/errors"

var (
	ErrInsufficientFunds = errors.Forbidden("insufficient funds")
)

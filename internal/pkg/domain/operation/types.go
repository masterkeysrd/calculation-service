package operation

import (
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/common/pagination"
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/common/search"
)

type ListRequest struct {
	search.Searchable
	pagination.Pageable
}

type OperationResponse struct {
	ID   uint          `json:"id"`
	Type OperationType `json:"type"`
	Cost float64       `json:"cost"`
}

package record

import (
	"time"

	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/common/pagination"
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/common/search"
)

type GetRecordRequest struct {
	ID     uint
	UserID uint
}

type ListRecordsRequest struct {
	UserID uint `json:"userId"`
	search.Searchable
	pagination.Pageable
}

type RecordResponse struct {
	ID            uint      `json:"id"`
	UserID        uint      `json:"userId"`
	OperationID   uint      `json:"operationId"`
	OperationType string    `json:"operationType"`
	Amount        float64   `json:"amount"`
	UserBalance   float64   `json:"userBalance"`
	Result        string    `json:"result"`
	CreatedAt     time.Time `json:"createdAt"`
}

type CreateRecordRequest struct {
	UserID      uint    `json:"userId"`
	OperationID uint    `json:"operationId"`
	Amount      float64 `json:"amount"`
	UserBalance float64 `json:"userBalance"`
	Result      string  `json:"result"`
}

type DeleteRecordRequest struct {
	ID     uint
	UserID uint
}

package record

import "time"

type Record struct {
	ID          uint64    `json:"id"`
	UserID      uint64    `json:"user_id"`
	OperationId uint64    `json:"operation_id"`
	Amount      float64   `json:"amount"`
	UserBalance float64   `json:"user_balance"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

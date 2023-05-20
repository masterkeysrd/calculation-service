package record

import "time"

type Record struct {
	ID          uint64    `json:"id"`
	UserID      uint64    `json:"user_id"`
	OperationID uint64    `json:"operation_id"`
	Amount      float64   `json:"amount"`
	UserBalance float64   `json:"user_balance"`
	CreatedAt   time.Time `json:"created_at"`
}

type NewRecordInput struct {
	UserID      uint64  `json:"user_id"`
	OperationID uint64  `json:"operation_id"`
	Amount      float64 `json:"amount"`
	UserBalance float64 `json:"user_balance"`
}

func NewRecord(input NewRecordInput) *Record {
	return &Record{
		UserID:      input.UserID,
		OperationID: input.OperationID,
		Amount:      input.Amount,
		UserBalance: input.UserBalance,
		CreatedAt:   time.Now(),
	}
}

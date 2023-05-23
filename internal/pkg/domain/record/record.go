package record

import "time"

type Record struct {
	ID          uint
	UserID      uint
	Operation   RecordOperation
	Amount      float64
	UserBalance float64
	Result      string
	CreatedAt   time.Time
}

type RecordOperation struct {
	ID   uint
	Type string
}

type NewRecordInput struct {
	UserID      uint
	OperationID uint
	Amount      float64
	UserBalance float64
	Result      string
}

func NewRecord(input NewRecordInput) *Record {
	return &Record{
		UserID: input.UserID,
		Operation: RecordOperation{
			ID: input.OperationID,
		},
		Amount:      input.Amount,
		UserBalance: input.UserBalance,
		Result:      input.Result,
		CreatedAt:   time.Now(),
	}
}

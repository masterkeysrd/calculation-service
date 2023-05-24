package record

func mapRecordToResponse(record Record) RecordResponse {
	return RecordResponse{
		ID:            record.ID,
		UserID:        record.UserID,
		OperationID:   record.Operation.ID,
		OperationType: record.Operation.Type,
		Amount:        record.Amount,
		UserBalance:   record.UserBalance,
		Result:        record.Result,
		CreatedAt:     record.CreatedAt,
	}
}

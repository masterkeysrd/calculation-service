package record

import (
	"errors"
	"time"

	"go.uber.org/dig"
)

type Service interface {
	Get(GetRecordRequest) (*RecordResponse, error)
	List(ListRecordRequest) (*ListRecordResponse, error)
	Create(CreateRecordRequest) (*RecordResponse, error)
	Delete(DeleteRecordRequest) error
}

type ListRecordRequest struct {
	UserID uint
}

type ListRecordResponse struct {
	Data []*RecordResponse `json:"data"`
}

type GetRecordRequest struct {
	UserID uint
	ID     uint
}

type RecordResponse struct {
	ID            uint      `json:"id"`
	UserID        uint      `json:"userId"`
	OperationID   uint      `json:"operationId"`
	OperationType string    `json:"operationType"`
	Amount        float64   `json:"amount"`
	UserBalance   float64   `json:"userBalance"`
	Result        string    `json:"result"`
	CreatedAt     time.Time `json:"date"`
}

type CreateRecordRequest struct {
	UserID      uint    `json:"userId"`
	OperationID uint    `json:"operationId"`
	Amount      float64 `json:"amount"`
	UserBalance float64 `json:"userBalance"`
	Result      string  `json:"result"`
}

type DeleteRecordRequest struct {
	UserID uint
	ID     uint
}

type recordService struct {
	repository Repository
}

type RecordServiceParams struct {
	dig.In
	Repository Repository
}

func NewService(params RecordServiceParams) Service {
	return &recordService{
		repository: params.Repository,
	}
}

func (s *recordService) Get(request GetRecordRequest) (*RecordResponse, error) {
	record, err := s.repository.GetWithUserID(request.UserID, request.ID)
	if err != nil {
		return nil, err
	}

	return mapRecordToResponse(record), nil
}

func (s *recordService) List(request ListRecordRequest) (*ListRecordResponse, error) {
	records, err := s.repository.ListWithUserID(request.UserID)
	if err != nil {
		return nil, err
	}

	result := []*RecordResponse{}
	for _, record := range records {
		result = append(result, mapRecordToResponse(record))
	}

	return &ListRecordResponse{
		Data: result,
	}, nil
}

func (s *recordService) Create(request CreateRecordRequest) (*RecordResponse, error) {
	record := NewRecord(
		NewRecordInput(request),
	)

	err := s.repository.Create(record)
	if err != nil {
		return nil, err
	}

	return mapRecordToResponse(record), nil
}

func (s *recordService) Delete(request DeleteRecordRequest) error {
	record, err := s.repository.GetWithUserID(request.UserID, request.ID)
	if err != nil {
		return err
	}

	if record == nil {
		return errors.New("record not found")
	}

	return s.repository.Delete(record)
}

func mapRecordToResponse(record *Record) *RecordResponse {
	return &RecordResponse{
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

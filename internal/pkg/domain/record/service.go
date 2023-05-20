package record

import (
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
	UserID uint64
}

type ListRecordResponse struct {
	Data []RecordResponse `json:"data"`
}

type GetRecordRequest struct {
	UserID uint64
	ID     uint64
}

type RecordResponse struct {
	ID          uint64    `json:"id"`
	UserID      uint64    `json:"userId"`
	OperationID uint64    `json:"operationId"`
	Amount      float64   `json:"amount"`
	UserBalance float64   `json:"userBalance"`
	CreatedAt   time.Time `json:"date"`
}

type CreateRecordRequest struct {
	UserID      uint64  `json:"userId"`
	OperationID uint64  `json:"operationId"`
	Amount      float64 `json:"amount"`
	UserBalance float64 `json:"userBalance"`
}

type DeleteRecordRequest struct {
	UserID uint64
	ID     uint64
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
	record, err := s.repository.FindByUserIDAndID(request.UserID, request.ID)
	if err != nil {
		return nil, err
	}

	return &RecordResponse{
		ID:          record.ID,
		UserID:      record.UserID,
		OperationID: record.OperationID,
		Amount:      record.Amount,
		UserBalance: record.UserBalance,
		CreatedAt:   record.CreatedAt,
	}, nil
}

func (s *recordService) List(request ListRecordRequest) (*ListRecordResponse, error) {
	records, err := s.repository.FindByUserID(request.UserID)
	if err != nil {
		return nil, err
	}

	result := []RecordResponse{}
	for _, record := range *records {
		result = append(result, RecordResponse(record))
	}

	return &ListRecordResponse{
		Data: result,
	}, nil
}

func (s *recordService) Create(request CreateRecordRequest) (*RecordResponse, error) {
	record := NewRecord(NewRecordInput{
		UserID:      request.UserID,
		OperationID: request.OperationID,
		Amount:      request.Amount,
		UserBalance: request.UserBalance,
	})

	err := s.repository.Create(record)
	if err != nil {
		return nil, err
	}

	return &RecordResponse{
		ID:          record.ID,
		UserID:      record.UserID,
		OperationID: record.OperationID,
		Amount:      record.Amount,
		UserBalance: record.UserBalance,
		CreatedAt:   record.CreatedAt,
	}, nil
}

func (s *recordService) Delete(request DeleteRecordRequest) error {
	record, err := s.repository.FindByUserIDAndID(request.UserID, request.ID)
	if err != nil {
		return err
	}

	return s.repository.Delete(record)
}

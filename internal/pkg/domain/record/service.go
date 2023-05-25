package record

import (
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/common/pagination"
	"go.uber.org/dig"
)

type Service interface {
	Get(GetRecordRequest) (RecordResponse, error)
	List(ListRecordsRequest) (pagination.Page[RecordResponse], error)
	Create(CreateRecordRequest) (RecordResponse, error)
	Delete(DeleteRecordRequest) error
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

func (s *recordService) Get(request GetRecordRequest) (RecordResponse, error) {
	record, err := s.repository.GetWithUserID(request.UserID, request.ID)
	if err != nil {
		return RecordResponse{}, err
	}

	return mapRecordToResponse(*record), nil
}

func (s *recordService) List(request ListRecordsRequest) (pagination.Page[RecordResponse], error) {
	page, err := s.repository.List(request)
	if err != nil {
		return nil, err
	}

	return pagination.MapPage(page, mapRecordToResponse), nil
}

func (s *recordService) Create(request CreateRecordRequest) (RecordResponse, error) {
	record := NewRecord(
		NewRecordInput(request),
	)

	err := s.repository.Create(record)
	if err != nil {
		return RecordResponse{}, err
	}

	return mapRecordToResponse(*record), nil
}

func (s *recordService) Delete(request DeleteRecordRequest) error {
	record, err := s.repository.GetWithUserID(request.UserID, request.ID)
	if err != nil {
		return err
	}

	if record == nil {
		return ErrRecordNotFound
	}

	return s.repository.Delete(record)
}

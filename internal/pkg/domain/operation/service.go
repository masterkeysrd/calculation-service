package operation

import (
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/common/pagination"
	"go.uber.org/dig"
)

type Service interface {
	Get(id uint) (*OperationResponse, error)
	List(ListRequest) (pagination.Page[OperationResponse], error)
}

type OperationServiceParams struct {
	dig.In
	Repository Repository
}

type service struct {
	repository Repository
}

func NewOperationService(params OperationServiceParams) Service {
	return &service{
		repository: params.Repository,
	}
}

func (s *service) Get(id uint) (*OperationResponse, error) {
	operation, err := s.repository.Get(id)
	if err != nil {
		return nil, err
	}

	return &OperationResponse{
		ID:   operation.ID,
		Type: operation.Type,
		Cost: operation.Cost,
	}, nil
}

func (s *service) List(request ListRequest) (pagination.Page[OperationResponse], error) {
	operations, err := s.repository.List(request)
	if err != nil {
		return nil, err
	}

	return pagination.MapPage(operations, mapOperationToResponse), nil
}

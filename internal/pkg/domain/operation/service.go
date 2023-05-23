package operation

import "go.uber.org/dig"

type OperationResponse struct {
	ID   uint          `json:"id"`
	Type OperationType `json:"type"`
	Cost float64       `json:"cost"`
}
type ListOperationsResponse struct {
	Data []*OperationResponse `json:"data"`
}

type Service interface {
	Get(id uint) (*OperationResponse, error)
	List() (*ListOperationsResponse, error)
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

func (s *service) List() (*ListOperationsResponse, error) {
	operations, err := s.repository.List()
	if err != nil {
		return nil, err
	}

	var operationsResponse []*OperationResponse
	for _, operation := range operations {
		op := OperationResponse(*operation)
		operationsResponse = append(operationsResponse, &op)
	}

	return &ListOperationsResponse{
		Data: operationsResponse,
	}, nil
}

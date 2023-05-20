package operation

import "go.uber.org/dig"

type OperationResponse struct {
	ID   uint64        `json:"id"`
	Type OperationType `json:"type"`
	Cost float64       `json:"cost"`
}
type ListOperationsResponse struct {
	Data []OperationResponse `json:"data"`
}

type Service interface {
	Get(id uint64) (*OperationResponse, error)
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

func (s *service) Get(id uint64) (*OperationResponse, error) {
	operation, err := s.repository.FindByID(id)
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
	operations, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	var operationsResponse []OperationResponse
	for _, operation := range operations {
		operationsResponse = append(operationsResponse, OperationResponse(operation))
	}

	return &ListOperationsResponse{
		Data: operationsResponse,
	}, nil
}

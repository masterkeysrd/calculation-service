package balance

import (
	"go.uber.org/dig"
)

type Service interface {
	FindByUserID(userID uint) (*BalanceGetResponse, error)
	Reserve(request BalanceTransactionRequest) (*BalanceGetResponse, error)
	Release(request BalanceTransactionRequest) (*BalanceGetResponse, error)
	Commit(request BalanceTransactionRequest) (*BalanceGetResponse, error)
	Rollback(request BalanceTransactionRequest) (*BalanceGetResponse, error)
}

type BalanceGetResponse struct {
	Amount         float64 `json:"amount"`
	InFlightAmount float64 `json:"inFlightAmount"`
}

type CreateBalanceRequest struct {
	UserID uint    `json:"userId"`
	Amount float64 `json:"amount"`
}

type BalanceTransactionRequest struct {
	UserID uint    `json:"userId"`
	Amount float64 `json:"amount"`
}

type balanceService struct {
	repository Repository
}

type ServiceParams struct {
	dig.In
	Repository Repository
}

func NewService(params ServiceParams) Service {
	return &balanceService{
		repository: params.Repository,
	}
}

func (s *balanceService) FindByUserID(userID uint) (*BalanceGetResponse, error) {
	balance, err := s.repository.GetWithUserID(userID)
	if err != nil {
		return nil, err
	}

	return mapToResponse(balance), nil
}

func (s *balanceService) Reserve(request BalanceTransactionRequest) (*BalanceGetResponse, error) {
	balance, err := s.repository.PerformTransaction(request.UserID, func(balance *Balance) error {
		return balance.Reserve(request.Amount)
	})

	if err != nil {
		return nil, err
	}

	return mapToResponse(balance), nil
}

func (s *balanceService) Release(request BalanceTransactionRequest) (*BalanceGetResponse, error) {
	balance, err := s.repository.PerformTransaction(request.UserID, func(balance *Balance) error {
		return balance.Release(request.Amount)
	})

	if err != nil {
		return nil, err
	}

	return mapToResponse(balance), nil
}

func (s *balanceService) Commit(request BalanceTransactionRequest) (*BalanceGetResponse, error) {
	balance, err := s.repository.PerformTransaction(request.UserID, func(balance *Balance) error {
		return balance.Confirm(request.Amount)
	})

	if err != nil {
		return nil, err
	}

	return mapToResponse(balance), nil
}

func (s *balanceService) Rollback(request BalanceTransactionRequest) (*BalanceGetResponse, error) {
	balance, err := s.repository.PerformTransaction(request.UserID, func(balance *Balance) error {
		return balance.Rollback(request.Amount)
	})

	if err != nil {
		return nil, err
	}

	return mapToResponse(balance), nil
}

func mapToResponse(balance *Balance) *BalanceGetResponse {
	return &BalanceGetResponse{
		Amount:         balance.Amount,
		InFlightAmount: balance.InFlightAmount,
	}
}

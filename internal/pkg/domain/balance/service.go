package balance

import "go.uber.org/dig"

type Service interface {
	FindByUserID(userID uint64) (*BalanceGetResponse, error)
	Create(request CreateBalanceRequest) error
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
	UserID uint64  `json:"userId"`
	Amount float64 `json:"amount"`
}

type BalanceTransactionRequest struct {
	UserID uint64  `json:"userId"`
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

func (s *balanceService) FindByUserID(userID uint64) (*BalanceGetResponse, error) {
	balance, err := s.repository.FindByUserID(userID)
	if err != nil {
		return nil, err
	}

	return &BalanceGetResponse{
		Amount:         balance.Amount,
		InFlightAmount: balance.InFlight,
	}, nil
}

func (s *balanceService) Create(request CreateBalanceRequest) error {
	balance := NewBalance(NewBalanceInput(request))
	return s.repository.Create(balance)
}

func (s *balanceService) Reserve(request BalanceTransactionRequest) (*BalanceGetResponse, error) {
	balance, err := s.repository.FindByUserID(request.UserID)
	if err != nil {
		return nil, err
	}

	err = balance.Reserve(request.Amount)
	if err != nil {
		return nil, err
	}

	err = s.repository.Update(balance)
	if err != nil {
		return nil, err
	}

	return &BalanceGetResponse{
		Amount:         balance.Amount,
		InFlightAmount: balance.InFlight,
	}, nil
}

func (s *balanceService) Release(request BalanceTransactionRequest) (*BalanceGetResponse, error) {
	balance, err := s.repository.FindByUserID(request.UserID)
	if err != nil {
		return nil, err
	}

	err = balance.Release(request.Amount)
	if err != nil {
		return nil, err
	}

	err = s.repository.Update(balance)
	if err != nil {
		return nil, err
	}

	return &BalanceGetResponse{
		Amount:         balance.Amount,
		InFlightAmount: balance.InFlight,
	}, nil
}

func (s *balanceService) Commit(request BalanceTransactionRequest) (*BalanceGetResponse, error) {
	balance, err := s.repository.FindByUserID(request.UserID)
	if err != nil {
		return nil, err
	}

	err = balance.Confirm(request.Amount)

	if err != nil {
		return nil, err
	}

	err = s.repository.Update(balance)
	if err != nil {
		return nil, err
	}

	return &BalanceGetResponse{
		Amount:         balance.Amount,
		InFlightAmount: balance.InFlight,
	}, nil
}

func (s *balanceService) Rollback(request BalanceTransactionRequest) (*BalanceGetResponse, error) {
	balance, err := s.repository.FindByUserID(request.UserID)
	if err != nil {
		return nil, err
	}

	err = balance.Rollback(request.Amount)

	if err != nil {
		return nil, err
	}

	err = s.repository.Update(balance)
	if err != nil {
		return nil, err
	}

	return &BalanceGetResponse{
		Amount:         balance.Amount,
		InFlightAmount: balance.InFlight,
	}, nil
}

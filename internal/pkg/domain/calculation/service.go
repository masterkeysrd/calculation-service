package calculation

import (
	"errors"
	"time"

	"github.com/masterkeysrd/calculation-service/internal/pkg/domain/balance"
	"github.com/masterkeysrd/calculation-service/internal/pkg/domain/operation"
	"github.com/masterkeysrd/calculation-service/internal/pkg/domain/user"
	"go.uber.org/dig"
)

type Service interface {
	Calculate(request CalculateRequest) (*CalculateResponse, error)
}

type CalculateRequest struct {
	UserID      uint64   `json:"userId"`
	OperationId uint64   `json:"operationId" binding:"required" validate:"required"`
	Arguments   []string `json:"arguments" binding:"required" validate:"required"`
}

type CalculateResponse struct {
	RecordID    uint64    `json:"recordId"`
	Amount      float64   `json:"amount"`
	UserBalance float64   `json:"userBalance"`
	Result      string    `json:"result"`
	Date        time.Time `json:"date"`
}

type ServiceParams struct {
	dig.In
	OperationService operation.Service
	UserService      user.Service
	BalanceService   balance.Service
}

type service struct {
	operationService operation.Service
	userService      user.Service
	balanceService   balance.Service
}

func NewService(params ServiceParams) Service {
	return &service{
		userService:      params.UserService,
		operationService: params.OperationService,
		balanceService:   params.BalanceService,
	}
}

func (s *service) Calculate(request CalculateRequest) (*CalculateResponse, error) {
	operation, err := s.operationService.Get(request.OperationId)

	if err != nil {
		return nil, err
	}

	transaction := balance.BalanceTransactionRequest{
		UserID: request.UserID,
		Amount: operation.Cost,
	}

	_, err = s.balanceService.Reserve(transaction)
	if err != nil {
		return nil, err
	}

	result, err := performOperation(operation.Type, request.Arguments)
	if err != nil {
		s.balanceService.Release(transaction)
		return nil, err
	}

	balance, err := s.balanceService.Commit(transaction)
	if err != nil {
		s.balanceService.Release(transaction)
		return nil, err
	}

	return &CalculateResponse{
		RecordID:    0,
		Amount:      operation.Cost,
		UserBalance: balance.Amount,
		Result:      result,
		Date:        time.Now(),
	}, nil

}

func performOperation(operationType operation.OperationType, arguments []string) (string, error) {
	switch operationType {
	case operation.OperationTypeAddition:
		return addition(arguments)
	case operation.OperationTypeSubtraction:
		return subtraction(arguments)
	case operation.OperationTypeMultiplication:
		return multiplication(arguments)
	case operation.OperationTypeDivision:
		return division(arguments)
	default:
		return "", errors.New("operation not supported")
	}
}

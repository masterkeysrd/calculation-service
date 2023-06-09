package calculation

import (
	"time"

	"github.com/masterkeysrd/calculation-service/internal/pkg/domain/balance"
	"github.com/masterkeysrd/calculation-service/internal/pkg/domain/operation"
	"github.com/masterkeysrd/calculation-service/internal/pkg/domain/record"
	"github.com/masterkeysrd/calculation-service/internal/pkg/domain/user"
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/random"
	"go.uber.org/dig"
)

type Service interface {
	Calculate(request CalculateRequest) (*CalculateResponse, error)
}

type CalculateRequest struct {
	UserID      uint     `json:"userId"`
	OperationId uint     `json:"operationId" binding:"required" validate:"required"`
	Arguments   []string `json:"arguments" binding:"required" validate:"required"`
}

type CalculateResponse struct {
	RecordID    uint      `json:"recordId"`
	Amount      float64   `json:"amount"`
	UserBalance float64   `json:"userBalance"`
	Result      string    `json:"result"`
	Date        time.Time `json:"date"`
}

type ServiceParams struct {
	dig.In
	RandomClient     random.Client
	UserService      user.Service
	RecordService    record.Service
	BalanceService   balance.Service
	OperationService operation.Service
}

type service struct {
	randomClient     random.Client
	userService      user.Service
	recordService    record.Service
	balanceService   balance.Service
	operationService operation.Service
}

func NewService(params ServiceParams) Service {
	return &service{
		userService:      params.UserService,
		recordService:    params.RecordService,
		balanceService:   params.BalanceService,
		operationService: params.OperationService,
		randomClient:     params.RandomClient,
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

	result, err := s.performOperation(operation.Type, request.Arguments)
	if err != nil {
		s.balanceService.Release(transaction)
		return nil, err
	}

	balance, err := s.balanceService.Commit(transaction)
	if err != nil {
		s.balanceService.Release(transaction)
		return nil, err
	}

	record, err := s.recordService.Create(record.CreateRecordRequest{
		UserID:      request.UserID,
		OperationID: request.OperationId,
		Amount:      operation.Cost,
		UserBalance: balance.Amount,
		Result:      result,
	})

	if err != nil {
		s.balanceService.Rollback(transaction)
		return nil, err
	}

	return &CalculateResponse{
		RecordID:    record.ID,
		Amount:      operation.Cost,
		UserBalance: balance.Amount,
		Result:      result,
		Date:        record.CreatedAt,
	}, nil
}

func (s *service) performOperation(op operation.OperationType, arguments []string) (string, error) {
	if op == operation.OperationTypeRandomString {
		return s.randomString()
	}

	return performOperation(op, arguments)
}

func (s *service) randomString() (string, error) {
	randomOperation := random.NewGenerateStringOperation(8)
	result, err := s.randomClient.GenerateRandom(randomOperation)

	if err != nil {
		return "", err
	}

	if len(result.GetData()) == 0 {
		return "", ErrRandomDataEmpty
	}

	return result.GetData()[0].(string), nil
}

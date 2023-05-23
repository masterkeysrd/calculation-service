package repositories

import (
	"github.com/masterkeysrd/calculation-service/internal/pkg/domain/operation"
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/persistence/models"
	"gorm.io/gorm"
)

type operationRepository struct {
	db *gorm.DB
}

func NewOperationRepository(db *gorm.DB) operation.Repository {
	return &operationRepository{db: db}
}

func (r *operationRepository) List() ([]*operation.Operation, error) {
	var operations []models.Operation
	if err := r.db.Find(&operations).Error; err != nil {
		return nil, err
	}

	var operationsResponse []*operation.Operation
	for _, op := range operations {
		operationsResponse = append(operationsResponse, &operation.Operation{
			ID:   op.ID,
			Type: operation.OperationType(op.OperationType),
			Cost: op.Cost,
		})
	}

	return operationsResponse, nil
}

func (r *operationRepository) Get(id uint) (*operation.Operation, error) {
	var op models.Operation
	if err := r.db.First(&op, id).Error; err != nil {
		return nil, err
	}

	return &operation.Operation{
		ID:   op.ID,
		Type: operation.OperationType(op.OperationType),
		Cost: op.Cost,
	}, nil
}

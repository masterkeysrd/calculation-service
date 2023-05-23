package loader

import (
	"github.com/masterkeysrd/calculation-service/internal/pkg/domain/operation"
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/persistence/models"
	"gorm.io/gorm"
)

func LoadDefaultOperations(db *gorm.DB) error {
	operations := []models.Operation{
		{
			OperationType: string(operation.OperationTypeAddition),
			Cost:          1,
		},
		{
			OperationType: string(operation.OperationTypeSubtraction),
			Cost:          1,
		},
		{
			OperationType: string(operation.OperationTypeMultiplication),
			Cost:          2,
		},
		{
			OperationType: string(operation.OperationTypeDivision),
			Cost:          2,
		},
		{
			OperationType: string(operation.OperationTypeSquareRoot),
			Cost:          3,
		},
		{
			OperationType: string(operation.OperationTypeRandomString),
			Cost:          5,
		},
	}

	for _, op := range operations {
		if err := db.FirstOrCreate(&op, models.Operation{OperationType: op.OperationType}).Error; err != nil {
			return err
		}
	}

	return nil
}

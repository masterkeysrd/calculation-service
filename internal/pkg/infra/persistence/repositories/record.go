package repositories

import (
	rcd "github.com/masterkeysrd/calculation-service/internal/pkg/domain/record"
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/persistence/models"
	"gorm.io/gorm"
)

type recordRepository struct {
	db *gorm.DB
}

func NewRecordRepository(db *gorm.DB) rcd.Repository {
	return &recordRepository{db: db}
}

func (r *recordRepository) GetWithUserID(userID uint, id uint) (*rcd.Record, error) {
	var record models.Record

	if err := r.db.Joins("Operation").Where("user_id = ? AND id = ?", userID, id).First(&r).Error; err != nil {
		return nil, err
	}

	return mapModelToRecord(&record), nil
}

func (r *recordRepository) ListWithUserID(userID uint) ([]*rcd.Record, error) {
	var records []models.Record

	if err := r.db.Joins("Operation").Where("user_id = ?", userID).Find(&records).Error; err != nil {
		return nil, err
	}

	var recordList []*rcd.Record
	for _, record := range records {
		recordList = append(recordList, mapModelToRecord(&record))
	}

	return recordList, nil
}

func (r *recordRepository) Create(record *rcd.Record) error {
	recordModel := &models.Record{
		UserID:      record.UserID,
		OperationID: record.Operation.ID,
		Amount:      record.Amount,
		UserBalance: record.UserBalance,
		Result:      record.Result,
	}

	if err := r.db.Create(recordModel).Error; err != nil {
		return err
	}

	record.ID = recordModel.ID
	record.CreatedAt = recordModel.CreatedAt
	record.Operation.Type = recordModel.Operation.OperationType
	return nil
}

func (r *recordRepository) Delete(record *rcd.Record) error {
	if err := r.db.Delete(&models.Record{}, record.ID).Error; err != nil {
		return err
	}

	return nil
}

func mapModelToRecord(model *models.Record) *rcd.Record {
	return &rcd.Record{
		ID:     model.ID,
		UserID: model.UserID,
		Operation: rcd.RecordOperation{
			ID:   model.OperationID,
			Type: model.Operation.OperationType,
		},
		Amount:      model.Amount,
		UserBalance: model.UserBalance,
		Result:      model.Result,
		CreatedAt:   model.CreatedAt,
	}
}

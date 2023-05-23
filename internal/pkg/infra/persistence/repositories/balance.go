package repositories

import (
	"github.com/masterkeysrd/calculation-service/internal/pkg/domain/balance"
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/persistence/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type balanceRepository struct {
	db *gorm.DB
}

func NewBalanceRepository(db *gorm.DB) balance.Repository {
	return &balanceRepository{
		db: db,
	}
}

func (r *balanceRepository) GetWithUserID(userID uint) (*balance.Balance, error) {
	var model models.Balance
	err := r.db.Where("user_id = ?", userID).First(&model).Error

	if err != nil {
		return nil, err
	}

	return &balance.Balance{
		ID:             model.ID,
		UserID:         model.UserID,
		Amount:         model.Amount,
		InFlightAmount: model.AmountInFlight,
	}, nil
}

func (r *balanceRepository) Delete(balance *balance.Balance) error {
	if err := r.db.Delete(&models.Balance{}, balance.ID).Error; err != nil {
		return err
	}

	return nil
}

func (r *balanceRepository) PerformTransaction(userId uint, fc func(*balance.Balance) error) (*balance.Balance, error) {
	var entity models.Balance
	var model *balance.Balance

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("user_id = ?", userId).First(&entity).Error; err != nil {
			return err
		}

		model = &balance.Balance{
			ID:             entity.ID,
			UserID:         entity.UserID,
			Amount:         entity.Amount,
			InFlightAmount: entity.AmountInFlight,
		}

		if err := fc(model); err != nil {
			return err
		}

		entity.Amount = model.Amount
		entity.AmountInFlight = model.InFlightAmount

		if err := tx.Save(&entity).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return model, nil
}

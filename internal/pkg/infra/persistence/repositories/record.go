package repositories

import (
	rcd "github.com/masterkeysrd/calculation-service/internal/pkg/domain/record"
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/common/pagination"
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/persistence/clauses"
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/persistence/models"
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/persistence/scopes"
	"gorm.io/gorm"
)

var (
	textSearchFields = clauses.Fields{"operation_type", "amount", "user_balance", "result"}
)

type recordRepository struct {
	db       *gorm.DB
	searcher clauses.TextSearcher
}

func NewRecordRepository(db *gorm.DB) rcd.Repository {
	return &recordRepository{
		db:       db,
		searcher: clauses.NewTextSearcher(textSearchFields),
	}
}

func (r *recordRepository) GetWithUserID(userID uint, id uint) (*rcd.Record, error) {
	var record models.Record

	if err := r.db.Joins("Operation").Where("user_id = ? AND id = ?", userID, id).First(&r).Error; err != nil {
		return nil, err
	}

	result := mapModelToRecord(&record)
	return &result, nil
}

func (r *recordRepository) List(request rcd.ListRecordsRequest) (pagination.Page[rcd.Record], error) {
	var total int64
	var records []*models.Record

	searchClause := r.searcher.Search(request)
	userIdScope := scopes.UserID(request.UserID)

	paginator := scopes.NewPaginator(
		r.db.
			Joins("Operation").
			Scopes(userIdScope).
			Clauses(searchClause),
	)

	db := r.db.Joins("Operation").
		Scopes(
			paginator.Paginate(scopes.PaginateOptions{
				Total:    &total,
				Value:    &records,
				Pageable: request,
			}),
			userIdScope,
		).
		Clauses(searchClause).
		Find(&records)

	if db.Error != nil {
		return nil, db.Error
	}

	page := pagination.NewPage(records, request, total)
	return pagination.MapPage(page, mapModelToRecord), nil
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

func mapModelToRecord(model *models.Record) rcd.Record {
	return rcd.Record{
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

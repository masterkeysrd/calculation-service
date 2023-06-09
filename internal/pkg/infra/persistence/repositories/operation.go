package repositories

import (
	"errors"

	"github.com/masterkeysrd/calculation-service/internal/pkg/domain/operation"
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/common/pagination"
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/persistence/clauses"
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/persistence/models"
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/persistence/scopes"
	"gorm.io/gorm"
)

var fields = clauses.Fields{"operation_type", "cost"}

type operationRepository struct {
	db *gorm.DB
}

func NewOperationRepository(db *gorm.DB) operation.Repository {
	return &operationRepository{db: db}
}

func (r *operationRepository) List(request operation.ListRequest) (pagination.Page[operation.Operation], error) {
	var total int64
	var operations []*models.Operation

	textSearch := clauses.NewTextSearcher(fields)
	paginator := scopes.NewPaginator(r.db)

	err := r.db.
		Scopes(
			paginator.Paginate(scopes.PaginateOptions{
				Total:    &total,
				Value:    &models.Operation{},
				Pageable: request,
			}),
		).
		Clauses(textSearch.Search(request)).
		Find(&operations).Error

	if err != nil {
		return nil, err
	}

	page := pagination.NewPage(operations, request, total)
	return pagination.MapPage(page, mapOperationToModel), nil
}

func (r *operationRepository) Get(id uint) (*operation.Operation, error) {
	var op models.Operation
	err := r.db.First(&op, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, operation.ErrOperationNotFound
	}

	if err != nil {
		return nil, err
	}

	return &operation.Operation{
		ID:   op.ID,
		Type: operation.OperationType(op.OperationType),
		Cost: op.Cost,
	}, nil
}

func mapOperationToModel(entity *models.Operation) operation.Operation {
	return operation.Operation{
		ID:   entity.ID,
		Type: operation.OperationType(entity.OperationType),
		Cost: entity.Cost,
	}
}

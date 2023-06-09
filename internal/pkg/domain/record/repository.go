package record

import (
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/common/pagination"
)

type Repository interface {
	GetWithUserID(userID uint, id uint) (*Record, error)
	List(request ListRecordsRequest) (pagination.Page[Record], error)
	Create(record *Record) error
	Delete(record *Record) error
}

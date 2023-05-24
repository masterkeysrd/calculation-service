package utils

import (
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/common/pagination"
	"gorm.io/gorm"
)

type Paginator interface {
	Paginate(value interface{}, pageable pagination.Pageable, total *int64) func(db *gorm.DB) *gorm.DB
}

type paginator struct {
	db *gorm.DB
}

func NewPaginator(db *gorm.DB) Paginator {
	return &paginator{db: db}
}

func (p *paginator) Paginate(value interface{}, pageable pagination.Pageable, total *int64) func(db *gorm.DB) *gorm.DB {
	p.db.Model(value).Count(total)
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pageable.GetOffset()).Limit(pageable.GetPageSize())
	}
}

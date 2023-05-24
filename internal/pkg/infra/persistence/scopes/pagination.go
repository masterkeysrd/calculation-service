package scopes

import (
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/common/pagination"
	"gorm.io/gorm"
)

type Paginator interface {
	Paginate(options PaginateOptions) func(db *gorm.DB) *gorm.DB
}
type PaginateOptions struct {
	Total    *int64
	Value    interface{}
	Pageable pagination.Pageable
}

type paginator struct {
	db *gorm.DB
}

func NewPaginator(db *gorm.DB) Paginator {
	return &paginator{db: db}
}

func (p *paginator) Paginate(opts PaginateOptions) func(db *gorm.DB) *gorm.DB {
	p.db.Model(opts.Value).Count(opts.Total)
	return func(db *gorm.DB) *gorm.DB {
		sort := opts.Pageable.GetSort().ToString()
		return db.Offset(opts.Pageable.GetOffset()).Limit(opts.Pageable.GetPageSize()).Order(sort)
	}
}

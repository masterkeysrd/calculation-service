package operation

import "github.com/masterkeysrd/calculation-service/internal/pkg/infra/common/pagination"

type Repository interface {
	Get(id uint) (*Operation, error)
	List(ListRequest) (pagination.Page[Operation], error)
}

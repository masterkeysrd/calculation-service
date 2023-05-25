package request

import (
	"github.com/gin-gonic/gin"
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/common/pagination"
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/common/search"
)

func PageableAndSearchable(ctx *gin.Context) (search.Searchable, pagination.Pageable, error) {
	pageable, err := Pageable(ctx)

	if err != nil {
		return nil, nil, err
	}

	searchable, err := Searchable(ctx)

	if err != nil {
		return nil, nil, err
	}

	return searchable, pageable, nil
}

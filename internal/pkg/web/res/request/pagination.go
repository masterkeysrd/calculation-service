package request

import (
	"github.com/gin-gonic/gin"
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/common/pagination"
)

func Pageable(ctx *gin.Context) (pagination.Pageable, error) {
	var pageable pagination.PageableRequest

	if err := ctx.ShouldBindQuery(&pageable); err != nil {
		return nil, err
	}

	return pagination.NewPageable(pageable), nil
}

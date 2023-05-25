package request

import (
	"github.com/gin-gonic/gin"
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/common/search"
)

func Searchable(ctx *gin.Context) (search.Searchable, error) {
	var searchable search.SearchableRequest
	if err := ctx.ShouldBindQuery(&searchable); err != nil {
		return nil, err
	}

	return search.NewSearchable(searchable), nil
}

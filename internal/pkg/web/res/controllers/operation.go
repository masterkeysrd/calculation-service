package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/masterkeysrd/calculation-service/internal/pkg/domain/operation"
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/common/pagination"
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/common/search"
	"go.uber.org/dig"
)

type OperationControllerParams struct {
	dig.In
	Service operation.Service
}

type OperationController struct {
	service operation.Service
}

func NewOperationController(params OperationControllerParams) *OperationController {
	return &OperationController{
		service: params.Service,
	}
}

func (c *OperationController) RegisterRoutes(group *gin.RouterGroup) {
	group.GET("", c.List)
	group.GET("/:id", c.Get)
}

func (c *OperationController) List(ctx *gin.Context) {
	var pageable pagination.PageableRequest

	if err := ctx.ShouldBindQuery(&pageable); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	var searchable search.SearchableRequest
	if err := ctx.ShouldBindQuery(&searchable); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	page, err := c.service.List(operation.ListRequest{
		Pageable:   pagination.NewPageable(pageable),
		Searchable: search.NewSearchable(searchable),
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, page.ToResponse())
}

func (c *OperationController) Get(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	operation, err := c.service.Get(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, operation)
}

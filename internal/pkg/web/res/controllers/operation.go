package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/masterkeysrd/calculation-service/internal/pkg/domain/operation"
	"github.com/masterkeysrd/calculation-service/internal/pkg/web/res/handlers"
	"github.com/masterkeysrd/calculation-service/internal/pkg/web/res/request"
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
	group.GET("", handlers.HandleError(c.List, handlers.DefaultResponseOptions))
	group.GET("/:id", handlers.HandleError(c.Get, handlers.DefaultResponseOptions))
}

func (c *OperationController) List(ctx *gin.Context) (interface{}, error) {
	searchable, pageable, err := request.PageableAndSearchable(ctx)

	if err != nil {
		return nil, err
	}

	page, err := c.service.List(operation.ListRequest{
		Pageable:   pageable,
		Searchable: searchable,
	})

	if err != nil {
		return nil, err
	}

	return page.ToResponse(), nil
}

func (c *OperationController) Get(ctx *gin.Context) (interface{}, error) {
	id, err := request.ParamUint(ctx, "id")

	if err != nil {
		return nil, err
	}

	return c.service.Get(uint(id))
}

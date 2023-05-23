package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/masterkeysrd/calculation-service/internal/pkg/domain/operation"
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
	operations, err := c.service.List()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, operations)
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

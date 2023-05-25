package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/masterkeysrd/calculation-service/internal/pkg/domain/calculation"
	"github.com/masterkeysrd/calculation-service/internal/pkg/web/res/handlers"
	"github.com/masterkeysrd/calculation-service/internal/pkg/web/res/request"
	"go.uber.org/dig"
)

type CalculationController struct {
	calculationService calculation.Service
}

type CalculationControllerParams struct {
	dig.In
	CalculationService calculation.Service
}

func NewCalculationController(params CalculationControllerParams) *CalculationController {
	return &CalculationController{
		calculationService: params.CalculationService,
	}
}

func (c *CalculationController) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/calculate", handlers.HandleError(c.Calculate, handlers.DefaultResponseOptions))
}

func (c *CalculationController) Calculate(ctx *gin.Context) (interface{}, error) {
	userID, err := request.UserID(ctx)
	if err != nil {
		return nil, err
	}

	var request calculation.CalculateRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return nil, err
	}

	request.UserID = userID
	return c.calculationService.Calculate(request)
}

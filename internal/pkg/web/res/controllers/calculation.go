package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/masterkeysrd/calculation-service/internal/pkg/domain/calculation"
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
	router.POST("/calculate", c.Calculate)
}

func (c *CalculationController) Calculate(ctx *gin.Context) {
	userId := ctx.GetUint("userId")
	var request calculation.CalculateRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	request.UserID = userId
	response, err := c.calculationService.Calculate(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

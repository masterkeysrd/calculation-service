package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/masterkeysrd/calculation-service/internal/pkg/domain/record"
	"go.uber.org/dig"
)

type RecordControllerParams struct {
	dig.In
	RecordService record.Service
}

type RecordController struct {
	recordService record.Service
}

func NewRecordController(params RecordControllerParams) *RecordController {
	return &RecordController{
		recordService: params.RecordService,
	}
}

func (c *RecordController) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("", c.List)
	router.GET(":id", c.Get)
	router.DELETE(":id", c.Delete)
}

func (c *RecordController) List(ctx *gin.Context) {
	userID := ctx.GetUint("userId")

	if userID == 0 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	records, err := c.recordService.List(record.ListRecordRequest{
		UserID: userID,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, records)
}

func (c *RecordController) Get(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}

func (c *RecordController) Delete(ctx *gin.Context) {
	userID := ctx.GetUint("userId")
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = c.recordService.Delete(record.DeleteRecordRequest{
		ID:     uint(id),
		UserID: userID,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

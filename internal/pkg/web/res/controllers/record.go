package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/masterkeysrd/calculation-service/internal/pkg/domain/record"
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/common/pagination"
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/common/search"
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

	var pageable pagination.PageableRequest
	if err := ctx.Bind(&pageable); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var searchable search.SearchableRequest
	if err := ctx.Bind(&searchable); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	request := record.ListRecordsInput{
		UserID:     userID,
		Searchable: search.NewSearchable(searchable),
	}

	result, err := c.recordService.List(request, pagination.NewPageable(pageable))

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, result.ToResponse())
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

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/masterkeysrd/calculation-service/internal/pkg/domain/record"
	"github.com/masterkeysrd/calculation-service/internal/pkg/web/res/handlers"
	"github.com/masterkeysrd/calculation-service/internal/pkg/web/res/request"
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
	router.GET("", handlers.HandleError(c.List, handlers.DefaultResponseOptions))
	router.DELETE(":id", handlers.HandleError(c.Delete, handlers.ResponseOptions{
		StatusCode: http.StatusNoContent,
	}))
}

func (c *RecordController) List(ctx *gin.Context) (interface{}, error) {
	userID, err := request.UserID(ctx)
	if err != nil {
		return nil, err
	}

	searchable, pageable, err := request.PageableAndSearchable(ctx)
	if err != nil {
		return nil, err
	}

	request := record.ListRecordsRequest{
		UserID:     userID,
		Searchable: searchable,
		Pageable:   pageable,
	}

	result, err := c.recordService.List(request)
	if err != nil {
		return nil, err
	}

	return result.ToResponse(), nil
}

func (c *RecordController) Delete(ctx *gin.Context) (interface{}, error) {
	userID, err := request.UserID(ctx)

	if err != nil {
		return nil, err
	}

	id, err := request.ParamUint(ctx, "id")
	if err != nil {
		return nil, err
	}

	return nil, c.recordService.Delete(record.DeleteRecordRequest{
		UserID: userID,
		ID:     uint(id),
	})
}

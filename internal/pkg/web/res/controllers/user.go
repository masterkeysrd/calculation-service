package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/masterkeysrd/calculation-service/internal/pkg/domain/user"
	"github.com/masterkeysrd/calculation-service/internal/pkg/web/res/handlers"
	"github.com/masterkeysrd/calculation-service/internal/pkg/web/res/request"
	"go.uber.org/dig"
)

type UserController struct {
	service user.Service
}

type UserControllerParams struct {
	dig.In
	Service user.Service
}

func NewUserController(options UserControllerParams) *UserController {
	return &UserController{
		service: options.Service,
	}
}

func (c *UserController) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("me", handlers.HandleError(c.FindMe, handlers.DefaultResponseOptions))
	router.DELETE("me", handlers.HandleError(c.DeleteMe, handlers.NoContentResponseOptions))
	router.GET("me/balance", handlers.HandleError(c.GetMyBalance, handlers.DefaultResponseOptions))
}

func (c *UserController) FindMe(ctx *gin.Context) (interface{}, error) {
	userId, err := request.UserID(ctx)

	if err != nil {
		return nil, err
	}

	return c.service.Get(userId)
}

func (c *UserController) DeleteMe(ctx *gin.Context) (interface{}, error) {
	userId, err := request.UserID(ctx)

	if err != nil {
		return nil, err
	}

	return nil, c.service.Delete(user.DeleteUserRequest{
		UserID: userId,
	})
}

func (c *UserController) GetMyBalance(ctx *gin.Context) (interface{}, error) {
	userId, err := request.UserID(ctx)

	if err != nil {
		return nil, err
	}

	return c.service.GetBalance(userId)
}

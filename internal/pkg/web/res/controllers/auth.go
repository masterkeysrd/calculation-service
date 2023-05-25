package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/masterkeysrd/calculation-service/internal/pkg/domain/auth"
	"github.com/masterkeysrd/calculation-service/internal/pkg/web/res/handlers"
	"go.uber.org/dig"
)

type AuthController struct {
	service auth.Service
}

type SignOutHeader = auth.SignOutRequest

type AuthControllerParams struct {
	dig.In
	Service auth.Service
}

func NewAuthController(options AuthControllerParams) *AuthController {
	return &AuthController{
		service: options.Service,
	}
}

func (c *AuthController) RegisterRoutes(group *gin.RouterGroup) {
	group.POST("/sign-up", handlers.HandleError(c.SignUp, handlers.NoContentResponseOptions))
	group.POST("/sign-in", handlers.HandleError(c.SignIn, handlers.DefaultResponseOptions))
	group.POST("/sign-out", handlers.HandleError(c.SignOut, handlers.NoContentResponseOptions))
}

func (c *AuthController) SignUp(ctx *gin.Context) (interface{}, error) {
	var request auth.SignUpRequest
	if err := ctx.ShouldBind(&request); err != nil {
		return nil, err
	}

	return nil, c.service.SignUp(request)
}

func (c *AuthController) SignIn(ctx *gin.Context) (interface{}, error) {
	var request auth.SignInRequest
	if err := ctx.ShouldBind(&request); err != nil {
		return nil, err
	}

	return c.service.SignIn(request)
}

func (c *AuthController) SignOut(ctx *gin.Context) (interface{}, error) {
	var headers SignOutHeader
	if err := ctx.BindHeader(&headers); err != nil {
		return nil, err
	}

	request := auth.SignOutRequest{
		AccessToken: headers.AccessToken,
	}

	return nil, c.service.SignOut(request)
}

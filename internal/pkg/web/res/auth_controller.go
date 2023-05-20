package res

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/masterkeysrd/calculation-service/internal/pkg/domain/auth"
)

type AuthController struct {
	service auth.Service
}

type SignOutHeader = auth.SignOutRequest

type NewAuthControllerOptions struct {
	Service auth.Service
}

func NewAuthController(options NewAuthControllerOptions) *AuthController {
	return &AuthController{
		service: options.Service,
	}
}

func (c *AuthController) RegisterRoutes(group *gin.RouterGroup) {
	group.POST("/sign-up", c.SignUp)
	group.POST("/sign-in", c.SignIn)
	group.POST("/sign-out", c.SignOut)
	group.POST("/refresh", c.Refresh)
}

func (c *AuthController) SignUp(ctx *gin.Context) {
	var request auth.SignUpRequest
	if err := ctx.ShouldBind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := c.service.SignUp(request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{})
}

func (c *AuthController) SignIn(ctx *gin.Context) {
	var request auth.SignInRequest
	if err := ctx.ShouldBind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	response, err := c.service.SignIn(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *AuthController) SignOut(ctx *gin.Context) {
	var headers SignOutHeader
	if err := ctx.BindHeader(&headers); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	request := auth.SignOutRequest{
		AccessToken: headers.AccessToken,
	}

	if err := c.service.SignOut(request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

func (c *AuthController) Refresh(ctx *gin.Context) {
	var request auth.RefreshRequest
	if err := ctx.ShouldBind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	response, err := c.service.Refresh(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

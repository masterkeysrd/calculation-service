package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/masterkeysrd/calculation-service/internal/pkg/domain/user"
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
	router.GET("me", c.FindMe)
	router.DELETE("me", c.DeleteMe)
	router.GET("me/balance", c.GetMyBalance)
}

func (c *UserController) FindMe(ctx *gin.Context) {
	userId := ctx.GetUint("userId")
	fmt.Println("userId: ", userId)

	if userId == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "Unauthorized",
		})
		return
	}

	user, err := c.service.Get(userId)

	if err != nil {
		ctx.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, user)
}

func (c *UserController) DeleteMe(ctx *gin.Context) {
	userId := ctx.GetUint("userId")

	if userId == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "Unauthorized",
		})
		return
	}

	err := c.service.Delete(user.DeleteUserRequest{
		UserID: userId,
	})

	if err != nil {
		ctx.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (c *UserController) GetMyBalance(ctx *gin.Context) {
	userId := ctx.GetUint("userId")

	if userId == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "Unauthorized",
		})
		return
	}

	balance, err := c.service.GetBalance(userId)
	if err != nil {
		ctx.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, balance)
}

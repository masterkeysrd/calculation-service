package auth

import "github.com/gin-gonic/gin"

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) RegisterRoutes(group *gin.RouterGroup) {
	group.POST("/sign-up", c.SignUp)
	group.POST("/sign-in", c.SignIn)
	group.POST("/sign-out", c.SignOut)
	group.POST("/refresh", c.Refresh)
}

func (c *Controller) SignUp(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}

func (c *Controller) SignIn(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}

func (c *Controller) SignOut(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}

func (c *Controller) Refresh(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}

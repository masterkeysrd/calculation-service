package res

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/masterkeysrd/calculation-service/internal/pkg/domain/user"
)

type UserController struct {
	service user.Service
}

type UserControllerOptions struct {
	Service user.Service
}

func NewUserController(options UserControllerOptions) *UserController {
	return &UserController{
		service: options.Service,
	}
}

func (c *UserController) RegisterRoutes(gin *gin.RouterGroup) {
	gin.GET("me", c.FindMe)
	gin.DELETE("me", c.DeleteMe)
}

func (c *UserController) FindMe(ctx *gin.Context) {
	user, err := c.service.FindByUserName("admin@test.com")

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
	err := c.service.Delete(user.DeleteUserRequest{
		UserName: "admin@test.com",
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

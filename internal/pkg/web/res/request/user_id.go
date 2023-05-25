package request

import (
	"github.com/gin-gonic/gin"
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/http/errors"
)

func UserID(ctx *gin.Context) (uint, error) {
	if userId, ok := ctx.Get("userId"); ok {
		return userId.(uint), nil
	}

	return 0, errors.Unauthorized("user id not found in context")
}

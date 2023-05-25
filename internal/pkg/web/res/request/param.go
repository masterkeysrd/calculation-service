package request

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/http/errors"
)

func ParamUint(ctx *gin.Context, key string) (uint, error) {
	param := ctx.Param(key)

	if param == "" {
		return 0, errors.BadRequest(fmt.Sprintf("param %s is required", key))
	}

	value, err := strconv.ParseUint(param, 10, 64)

	if err != nil {
		return 0, errors.BadRequest(fmt.Sprintf("param %s must be an integer", key))
	}

	return uint(value), nil
}

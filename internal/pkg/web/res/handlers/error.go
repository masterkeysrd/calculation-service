package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/http/errors"
)

var (
	DefaultResponseOptions = ResponseOptions{
		StatusCode: http.StatusOK,
	}
	NoContentResponseOptions = ResponseOptions{
		StatusCode: http.StatusNoContent,
	}
)

type ResponseOptions struct {
	StatusCode int
}

func HandleError(handler func(*gin.Context) (interface{}, error), options ResponseOptions) func(*gin.Context) {
	return func(ctx *gin.Context) {
		result, err := handler(ctx)

		if err != nil {
			code, response := ParseError(err)
			ctx.JSON(code, response)
			return
		}

		statusCode := options.StatusCode
		if statusCode == 0 {
			statusCode = http.StatusOK
		}

		ctx.JSON(statusCode, result)
	}
}

func ParseError(e error) (int, map[string]interface{}) {
	var response map[string]interface{}
	code := http.StatusInternalServerError

	switch err := e.(type) {
	case *strconv.NumError:
		code = http.StatusBadRequest
		response = ErrorToMap(http.StatusBadRequest, err.Error(), err)
	case errors.HTTPError:
		code = err.StatusCode()
		response = ErrorToMap(err.StatusCode(), err.Error(), err)
	default:
		response = ErrorToMap(http.StatusInternalServerError, err.Error(), err)
	}

	return code, response
}

func ErrorToMap(code int, message string, details interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":             code,
		"message":          message,
		"message_detailed": details,
	}
}

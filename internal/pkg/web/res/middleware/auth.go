package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/jwt"
)

type JWTAuthMiddleware = func() gin.HandlerFunc

func JWTAuthMiddlewareFactory(jwtService jwt.Service) JWTAuthMiddleware {
	return func() gin.HandlerFunc {
		return func(c *gin.Context) {

			if err := jwtService.ValidateToken(c); err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"code":  http.StatusUnauthorized,
					"error": fmt.Sprintf("Unauthorized: %s", err.Error()),
				})
				return
			}

			c.Next()
		}
	}
}

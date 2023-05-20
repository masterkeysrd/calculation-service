package jwt

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func ExtractToken(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")
	fmt.Println("ExtractToken: authHeader: ", authHeader)

	if len(authHeader) == 0 {
		return ""
	}

	token := strings.Split(authHeader, " ")
	if len(token) == 2 {
		return token[1]
	}

	return ""
}

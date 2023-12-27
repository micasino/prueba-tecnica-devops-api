package middleware

import (
	"log"
	helper "mg-inc-devops/pkg/helpers"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthorizationMiddleware(c *gin.Context) {
	s := c.Request.Header.Get("Authorization")

	token := strings.TrimPrefix(s, "Bearer ")

	if _, err := helper.ValidateToken(token); err != "" {
		log.Println("error:", err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}

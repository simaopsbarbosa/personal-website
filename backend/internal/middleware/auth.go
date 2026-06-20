package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"simaopsbarbosa/backend/internal/utils"
)

// protects routes by requiring a valid JWT in the Authorization header
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header must be Bearer token"})
			c.Abort()
			return
		}

		tokenString := authHeader[7:]
		_, err := utils.VerifyToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			c.Abort()
			return
		}

		c.Set("admin", true)
		c.Next()
	}
}

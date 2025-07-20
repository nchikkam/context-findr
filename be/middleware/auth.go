package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	utils "github.com/nchikkam/context-findr-be/utils/infrastructure"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Header missing XX Authorization"})
			c.Abort()
			return
		}

		const bearerPrefix = "Bearer "
		if !strings.HasPrefix(authHeader, bearerPrefix) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong Authorization header format"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, bearerPrefix)

		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token invalid"})
			c.Abort()
			return
		}

		c.Set("userEmail", claims.Email)

		c.Next()
	}
}

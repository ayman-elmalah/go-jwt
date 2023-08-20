package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IsGuest() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader != "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authenticated"})
			return
		}

		c.Next()
	}
}

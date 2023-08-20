package middlewares

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go-jwt/internal/modules/user/models"
	"go-jwt/internal/modules/user/responses"
	"go-jwt/pkg/config"
	"net/http"
	"strings"
)

func IsAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		cfg := config.Get()

		// Get the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			return
		}

		// Check if the header has the "Bearer" scheme
		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header"})
			return
		}

		tokenString := authHeaderParts[1]

		// Parse the token without claims validation
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.Jwt.Secret), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		user := claims["sub"]

		userModel := models.User{}
		userModelBytes, _ := json.Marshal(user)
		_ = json.Unmarshal(userModelBytes, &userModel)

		//Store the user in the context
		c.Set("auth", responses.ToUser(userModel))

		c.Next()
	}
}

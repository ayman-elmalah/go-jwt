package routes

import (
	"github.com/gin-gonic/gin"
	homeRoutes "go-jwt/internal/modules/home/routes"
	userRoutes "go-jwt/internal/modules/user/routes"
)

func RegisterRoutes(router *gin.Engine) {
	homeRoutes.Routes(router)
	userRoutes.Routes(router)
}

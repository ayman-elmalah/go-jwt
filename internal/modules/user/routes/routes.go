package routes

import (
	"github.com/gin-gonic/gin"
	"go-jwt/internal/middlewares"
	userCtrl "go-jwt/internal/modules/user/controllers"
)

func Routes(router *gin.Engine) {
	userController := userCtrl.New()

	guestGroup := router.Group("/api/v1/auth")
	guestGroup.Use(middlewares.IsGuest())
	{
		guestGroup.POST("/register", userController.HandleRegister)
		guestGroup.POST("/login", userController.HandleLogin)
	}

	authGroup := router.Group("/api/v1/auth")
	authGroup.Use(middlewares.IsAuth())
	{
		authGroup.GET("/user", userController.User)
	}
}

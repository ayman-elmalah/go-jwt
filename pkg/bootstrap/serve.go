package bootstrap

import (
	"go-jwt/pkg/config"
	"go-jwt/pkg/database"
	"go-jwt/pkg/routing"
)

func Serve() {
	config.Set()

	database.Connect()

	routing.Init()

	routing.RegisterRoutes()

	routing.Serve()
}

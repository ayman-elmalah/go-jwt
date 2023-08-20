package migrate

import (
	"go-jwt/internal/database/migration"
	"go-jwt/pkg/config"
	"go-jwt/pkg/database"
)

func Down() {
	config.Set()

	database.Connect()

	migration.MigrateDown()
}

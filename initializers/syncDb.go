package initializers

import "github.com/jm61/jwt/models"

func SyncDb() {
	DB.AutoMigrate(&models.User{})
}

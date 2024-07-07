package helpers

import (
	"chatter-backend-go/src/configs"
	"chatter-backend-go/src/models"
)

func Migration() {
	configs.DB.AutoMigrate(&models.User{})
}

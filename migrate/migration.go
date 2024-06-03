package migrate

import (
	"goAPI/database"
	"goAPI/models"
)

func SyncDB() {
	database.DB.AutoMigrate(&models.Users{}, &models.Photos{})
}

package initializers

import "GO-JWT/models"

func SyncDatabase() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		return
	}
}

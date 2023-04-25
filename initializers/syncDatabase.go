package initializers

import "github.com/Salavei/golang_gin/models"

func SyncDatabase() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		panic(err)
	}
}

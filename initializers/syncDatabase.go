package initializers

import (
	"PaymentAPI/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Transaction{})
}

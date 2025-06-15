package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/chinotanismail/internal-transfer-system/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=123 dbname=internal_transfer_system port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.AutoMigrate(&models.Account{}, &models.Transaction{})
	DB = db
}

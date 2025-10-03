package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"accountantapp/go-service/internal/models"
)

var DB *gorm.DB

func Connect() {
	// Настройка подключения к PostgreSQL
	dsn := "host=localhost user=postgres password=0000 dbname=accountant port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	// Авто-миграция моделей
	err = db.AutoMigrate(
		&models.User{},
		&models.Account{},
		&models.Transaction{},
		&models.Category{},
		&models.Template{},
		&models.AppSettings{},
	)
	if err != nil {
		panic("Failed to auto-migrate models!")
	}

	DB = db
	fmt.Println("Database connected and migrated")
}

package database

import (
	"os"

	// "gorm.io/driver/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"web_practicum/models"
)

var DB *gorm.DB

func Setup() error {

	connStr := os.Getenv("CONNECTION_STRING")
	var err error
	DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return err
	}
	return DB.AutoMigrate(&models.Article{})
}

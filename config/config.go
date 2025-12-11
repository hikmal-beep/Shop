package config

import (
	"Shop/models"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var JWTSecret = os.Getenv("JWT_SECRET")

func Connect() *gorm.DB {
	var (
		dbHost     = os.Getenv("DB_HOST")
		dbPort     = os.Getenv("DB_PORT")
		dbUsername = os.Getenv("DB_USERNAME")
		dbPassword = os.Getenv("DB_PASSWORD")
		dbName     = os.Getenv("DB_NAME")
		dbTimezone = os.Getenv("DB_TIMEZONE")
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", dbHost, dbUsername, dbPassword, dbName, dbPort, dbTimezone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{}, &models.Shop{}, &models.Product{})

	return db
}
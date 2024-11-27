package config

import (
	"fmt"
	"log"
	"news-go/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	ClickDB *gorm.DB
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func ConnectDB() {
	var err error
	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
	)

	DB, err = gorm.Open(postgres.Open(dbURI), &gorm.Config{})

	if err != nil {
			log.Fatal("Could not connect to the database", err)
	}

	if err := DB.AutoMigrate(&models.News{}); err != nil {
		log.Fatal("Could not create migrate models: RequestMetric", err)
	}

	if err := DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Could not create migrate models: RequestMetric", err)
	}

	if err := DB.AutoMigrate(&models.Client_metadata{}); err != nil {
		log.Fatal("Could not create migrate models: RequestMetric", err)
	}

	if err := DB.AutoMigrate(&models.RequestMetric{}); err != nil {
		log.Fatal("Could not create migrate models: RequestMetric", err)
	}
}

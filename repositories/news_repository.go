package repositories

import (
	"news-go/config"
	"news-go/models"
)

// GetAllNewsRepository returns all news from the database
func GetAllNewsRepository() ([]models.News, error) {
	var news []models.News
	result := config.DB.Find(&news)
	if result.Error != nil {
		return nil, result.Error
	}
	return news, nil
}

// CreateNewsRepository adds a new news item to the database
func CreateNewsRepository(news *models.News) error {
	return config.DB.Create(news).Error
}

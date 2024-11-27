package services

import (
	"news-go/models"
	"news-go/repositories"
)

// GetAllNews returns all news items using the repository
func GetAllNews() ([]models.News, error) {
	return repositories.GetAllNewsRepository()
}

// CreateNews adds a new news item via the repository
func CreateNews(news *models.News) error {
	return repositories.CreateNewsRepository(news)
}

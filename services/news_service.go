package services

import (
	"news-go/models"
	"news-go/repositories"
)

// GetAllNews возвращает все новости, используя репозиторий
func GetAllNews() ([]models.News, error) {
	return repositories.GetAllNewsRepository()
}

// CreateNews добавляет новость через репозиторий
func CreateNews(news *models.News) error {
	return repositories.CreateNewsRepository(news)
}

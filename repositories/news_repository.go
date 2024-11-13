package repositories

import (
	"news-go/config"
	"news-go/models"
)

// GetAllNewsRepository возвращает все новости из БД
func GetAllNewsRepository() ([]models.News, error) {
	var news []models.News
	result := config.DB.Find(&news)
	if result.Error != nil {
		return nil, result.Error
	}
	return news, nil
}

// CreateNewsRepository добавляет новую новость в БД
func CreateNewsRepository(news *models.News) error {
	return config.DB.Create(news).Error
}

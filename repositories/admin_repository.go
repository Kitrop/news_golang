package repositories

import (
	"news-go/config"
	"news-go/models"
)

// AdminDeleteNews удаляет новость из базы данных
func AdminDeleteNews(newsID uint) error {
	return config.DB.Model(&models.News{}).Where("id = ?", newsID).Delete(&models.News{}).Error
}

// BunUser меняет роль 
func BunUser(userID uint) error {
	return config.DB.Model(&models.News{}).Where("id = ?", userID).Update("role", "BANNED").Error
}
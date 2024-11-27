package repositories

import (
	"news-go/config"
	"news-go/models"
)

// AdminDeleteNews deletes news from the database
func AdminDeleteNews(newsID uint) error {
	return config.DB.Model(&models.News{}).Where("id = ?", newsID).Delete(&models.News{}).Error
}

// BunUser changes the role
func BunUser(userID uint) error {
	return config.DB.Model(&models.News{}).Where("id = ?", userID).Update("role", "BANNED").Error
}

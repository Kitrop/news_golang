package repositories

import (
	"errors"
	"news-go/config"
	"news-go/models"

	"gorm.io/gorm"
)

// CreateUserInDB создает пользователя в базе данных
func CreateUserInDB(user *models.User) error {
	return config.DB.Create(user).Error
}

// FindAllUsers возвращает всех пользователей из базы данных
func FindAllUsers() ([]models.User, error) {
	var users []models.User
	err := config.DB.Find(&users).Error
	return users, err
}

// FindUserByUsername ищет пользователя по имени пользователя
func FindUserByUsername(username string) (*models.User, error) {
	var user models.User
	result := config.DB.Where("username = ?", username).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}
	return &user, result.Error
}

// FindUserByEmail ищет пользователя по email
func FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := config.DB.Where("email = ?", email).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}
	return &user, result.Error
}

// UpdateUserPassword обновляет пароль пользователя в базе данных
func UpdateUserPassword(userID uint, newPassword string) error {
	return config.DB.Model(&models.User{}).Where("id = ?", userID).Update("password", newPassword).Error
}

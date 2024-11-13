package controller

import (
	"net/http"
	"news-go/models"
	"news-go/services"

	"github.com/gin-gonic/gin"
)

// CreateUserController создает нового пользователя
func CreateUserController(c *gin.Context) {
	var input models.User

	// Валидация входных данных
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Вызов сервиса для создания пользователя
	user, accessToken, err := services.CreateUser(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Установка JWT в cookies
	c.SetCookie("accessToken", accessToken, 36000, "/", "localhost", false, true)
	c.JSON(http.StatusCreated, gin.H{
		"message":     "User successfully created",
		"data":        user,
		"accessToken": accessToken,
	})
}

// GetAllUsersController возвращает список всех пользователей
func GetAllUsersController(c *gin.Context) {
	users, err := services.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// LoginController аутентификация пользователя и выдача JWT
func LoginController(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Логика аутентификации через сервис
	user, accessToken, err := services.LoginUser(input.Username, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Установка JWT в cookies
	c.SetCookie("accessToken", accessToken, 360000, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"data": user, "accessToken": accessToken})
}

// LogoutController удаляет accessToken из cookies
func LogoutController(c *gin.Context) {
	c.SetCookie("accessToken", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusNoContent, gin.H{"message": "Logged out successfully"})
}

// ChangePasswordController изменение пароля пользователя
func ChangePasswordController(c *gin.Context) {
	var input struct {
		OldPassword string `json:"oldPassword"`
		NewPassword string `json:"newPassword"`
	}

	// Валидация входных данных
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Получение accessToken из cookies
	accessToken, err := c.Cookie("accessToken")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization required"})
		return
	}

	// Вызов сервиса для смены пароля
	if err := services.ChangeUserPassword(accessToken, input.OldPassword, input.NewPassword); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password changed successfully"})
}
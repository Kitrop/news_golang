package controller

import (
	"net/http"
	"news-go/models"
	"news-go/services"

	"github.com/gin-gonic/gin"
)

// CreateUserController creates a new user
func CreateUserController(c *gin.Context) {
	var input models.User

	// Validate input data
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the service to create a user
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

// GetAllUsersController returns a list of all users
func GetAllUsersController(c *gin.Context) {
	users, err := services.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// LoginController authenticates the user and issues a JWT
func LoginController(c *gin.Context) {
	var input models.UserLogin
	
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Authentication logic through the service
	user, accessToken, err := services.LoginUser(input.Username, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// services.CreateNewSession(number.(user["userID"]), c.ClientIP(), c.Request.Header.Get("User-Agent"))

	// Установка JWT в cookies
	// Setting JWT in cookies
	c.SetCookie("accessToken", accessToken, 360000, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"data": user, "accessToken": accessToken})
}

// LogoutController removes the accessToken from cookies
func LogoutController(c *gin.Context) {
	c.SetCookie("accessToken", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusNoContent, gin.H{"message": "Logged out successfully"})
}

// ChangePasswordController changes the user's password
func ChangePasswordController(c *gin.Context) { 
	var input struct {
		OldPassword string `json:"oldPassword" binding:"required"`
		NewPassword string `json:"newPassword" binding:"required"`
	}

	// Валидация входных данных
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get accessToken from cookies
	accessToken, err := c.Cookie("accessToken")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization required"})
		return
	}

	// Call the service to change the password
	if err := services.ChangeUserPassword(accessToken, input.OldPassword, input.NewPassword); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Password changed successfully"})
}

package controller

import (
	"net/http"
	"news-go/config"
	"news-go/models"
	"news-go/utils"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := utils.ValidateUserInput(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password, try another input password"})
		return
	}

	newUser := models.User{Username: input.Username, Password: hashPassword, Email: input.Email}
	if err := utils.CreateUserInDB(&newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user in database"})
		return
	}

	accessToken, err := utils.GenerateJWT(input.Username, input.Email, string(newUser.Role))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate JWT"})
		return
	}

	c.SetCookie("accessToken", accessToken, 36000, "/", "localhost", false, true)

	userData := map[string]string{
		"username": newUser.Username,
		"email":    newUser.Email,
		"role":     string(newUser.Role),
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User successfully created",
		"data": userData,
		"accessToken": accessToken,
	})
}


func GetAllUsers(c *gin.Context) {
	var users []models.User

	allUsers := config.DB.Find(&users)

	if allUsers.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": allUsers.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func Login(c *gin.Context) {
	type User struct {
    Username string `json:"username"`
    Password string `json:"password"`
	}
	
	var input User
	
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	resultFindUser, err := utils.FindUserInDB(input.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if !utils.CheckPasswordHash(input.Password, resultFindUser.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
	}

	accessToken, err := utils.GenerateJWT(resultFindUser.Username, resultFindUser.Email, string(resultFindUser.Role))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.SetCookie("accessToken", accessToken, 360000, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"data": resultFindUser, "accessToken": accessToken})
}

func Logout(c *gin.Context) {
	c.SetCookie("accessToken", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusNoContent, gin.H{})
}
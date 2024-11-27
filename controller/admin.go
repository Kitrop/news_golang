package controller

import (
	"fmt"
	"net/http"
	"news-go/models"
	"news-go/services"

	"github.com/gin-gonic/gin"
)

// Setting the user's role to "BANNED", preventing them from interacting with the api
func BanUser(c *gin.Context) {
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.BanUser(input.ID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message := fmt.Sprintf("user with id - %d successfully bun", input.ID)
	c.JSON(http.StatusOK, gin.H{"message": message})
}

// Deleting news for the admin
func DeleteNews(c *gin.Context) {
	var input models.News

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.DeleteNews(input.ID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message := fmt.Sprintf("news with id - %d successfully delete", input.ID)
	c.JSON(http.StatusOK, gin.H{"message": message})
}

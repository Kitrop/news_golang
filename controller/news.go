package controller

import (
	"net/http"
	"news-go/models"
	"news-go/services"

	"github.com/gin-gonic/gin"
)

// GetAllNewsController возвращает все новости
func GetAllNewsController(c *gin.Context) {
	news, err := services.GetAllNews()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": news})
}

// CreateNewsController создает новость
func CreateNewsController(c *gin.Context) {
	var input models.News

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	news := &models.News{Text: input.Text}

	if err := services.CreateNews(news); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "News created", "data": news})
}
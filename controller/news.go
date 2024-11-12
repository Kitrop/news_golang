package controller

import (
	"net/http"
	"news-go/config"
	"news-go/models"
	"github.com/gin-gonic/gin"
)
func GetAllNews(c *gin.Context) {
	var news []models.News

	allNews := config.DB.Find(&news)

	if allNews.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": allNews.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": news})
}

func CreateNews(c *gin.Context) {
	var input models.News

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	news := models.News{Text: input.Text}
	result := config.DB.Create(&news)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{"message": "News created"})
}
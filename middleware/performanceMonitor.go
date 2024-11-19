package middleware

import (
	"news-go/config"
	"news-go/models"
	"time"

	"github.com/gin-gonic/gin"
)

func PerformanceMonitor() gin.HandlerFunc {
	db := config.DB
	return func(c *gin.Context) {
		startTime := time.Now()

		// Выполняем запрос
		c.Next()

		// Вычисляем время выполнения
		duration := time.Since(startTime).Seconds() * 1000 // миллисекунды

		// Сохраняем метрики в БД
		metric := models.RequestMetric{
			Path:         c.FullPath(),
			Method:       c.Request.Method,
			StatusCode:   c.Writer.Status(),
			ResponseTime: duration,
			Timestamp:    time.Now(),
		}

		if err := db.Create(&metric).Error; err != nil {
			c.Error(err)
		}
	}
}
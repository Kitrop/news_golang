package middleware

import (
	"news-go/config"
	"news-go/models"
	"time"

	"github.com/gin-gonic/gin"
)

// Middleware for collecting endpoints performance information
func PerformanceMonitor() gin.HandlerFunc {
	db := config.DB
	return func(c *gin.Context) {
		// Record the start time
		startTime := time.Now()

		// Execute the query
		c.Next()

		// Calculate the execution time
		duration := time.Since(startTime).Seconds() * 1000 // миллисекунды

	// Save the metrics to the database
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
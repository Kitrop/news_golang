package routes

import (
	"news-go/controller"
	"news-go/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterNewsRoutes registers routes for news management
func RegisterNewsRoutes(router *gin.RouterGroup) {
	router.POST("/new", middleware.Autification, controller.CreateNewsController)
	router.GET("/all", middleware.Autification, controller.GetAllNewsController)
}

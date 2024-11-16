package routes

import (
	"news-go/controller"
	"news-go/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterNewsRoutes(router *gin.RouterGroup) {
	router.POST("/new", middleware.Autification, controller.CreateNewsController)
	router.GET("/all", middleware.Autification, controller.GetAllNewsController)
}

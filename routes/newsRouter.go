package routes

import (
	"news-go/controller"

	"github.com/gin-gonic/gin"
)

func RegisterNewsRoutes(router *gin.RouterGroup) {
	router.POST("/new", controller.CreateNewsController)
	router.GET("/all", controller.GetAllNewsController)
}

package routes

import (
	"news-go/controller"
	"news-go/middleware"

	"github.com/gin-gonic/gin"
)

// Роутер для админки 
func RegisterAdminRoutes(router *gin.RouterGroup) {
	router.POST("/deleteNews", middleware.СheckIsAdmin, controller.DeleteNews)
	router.GET("/bunUser", middleware.СheckIsAdmin, controller.BanUser)
}

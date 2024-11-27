package routes

import (
	"github.com/gin-gonic/gin"
	"news-go/controller"
	"news-go/middleware"
)

// RegisterUserRoutes registers routes for user management
func RegisterUserRoutes(router *gin.RouterGroup) {
	router.POST("/create", middleware.UnAuthorized, controller.CreateUserController)
	router.GET("/all", middleware.Autification, controller.GetAllUsersController)
	router.POST("/login", middleware.UnAuthorized, controller.LoginController)
	router.GET("/logout", middleware.Autification, controller.LogoutController)
	router.POST("/change-password", middleware.Autification, controller.ChangePasswordController)
}

package routes

import (
	"news-go/controller"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.RouterGroup) {
	router.POST("/create", controller.CreateUserController)
	router.GET("/all", controller.GetAllUsersController)
	router.POST("/login", controller.LoginController)
	router.GET("/logout", controller.LogoutController)
	router.POST("/change-password", controller.ChangePasswordController)
}

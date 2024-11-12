package routes

import (
	"news-go/controller"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine) {
	newsRoutes := router.Group("/news")
	{
		newsRoutes.POST("/new", controller.CreateNews)
		newsRoutes.GET("/all", controller.GetAllNews)
	}

	usersRoutes := router.Group("/users")
	{
		usersRoutes.POST("/create", controller.CreateUser)
		usersRoutes.GET("/all", controller.GetAllUsers)
		usersRoutes.POST("/login", controller.Login)
		usersRoutes.GET("/logout", controller.Logout)
	}
}
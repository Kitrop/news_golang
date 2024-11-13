package main

import (
	"news-go/config"
	"news-go/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	r := gin.Default()

	api := r.Group("/api/v1")
	routes.RegisterUserRoutes(api.Group("/users"))
	routes.RegisterNewsRoutes(api.Group("/news"))

	r.Run(":8080")
}

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
	routes.RegisterUserRoutes(r)

	r.Run(":8080")
}
package main

import (
	"fmt"
	"log"
	"news-go/config"
	"news-go/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	r := gin.Default()

	routes.RegisterUserRoutes(r.Group("/users"))
	routes.RegisterNewsRoutes(r.Group("/news"))

	err := r.Run(":8080")

	if err != nil {
		errorString := fmt.Sprintf("Server not started by error: %s", err.Error())
		log.Fatal(errorString)
	}
}

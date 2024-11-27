package main

import (
	"fmt"
	"log"
	"news-go/config"
	"news-go/middleware"
	"news-go/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// DB and env connect
	config.LoadEnv()
	config.ConnectDB()

	r := gin.Default()

	// Middleware
	r.Use(middleware.CorsMiddleware())
	r.Use(middleware.GetAllClientData)
	r.Use(middleware.PerformanceMonitor())
	
	// Routes
	routes.RegisterUserRoutes(r.Group("/users"))
	routes.RegisterNewsRoutes(r.Group("/news"))
	routes.RegisterAdminRoutes(r.Group("/admin"))

	err := r.Run(":8080")

	if err != nil {
		errorString := fmt.Sprintf("Server not started by error: %s", err.Error())
		log.Fatal(errorString)
	}
}

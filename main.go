package main

import (
	"github.com/yordanos-habtamu/GinCrud/config"
	"github.com/yordanos-habtamu/GinCrud/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize configuration (e.g., environment variables)
	config.Init()

	// Create Gin router
	r := gin.Default()

	// Set up routes
	routes.SetupBookRoutes(r)

	// Run the server
	r.Run(":8080")
}

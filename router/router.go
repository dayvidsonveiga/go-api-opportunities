package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Initialize Router
	router := gin.Default()

	// Initialize Routes
	InitializeRoutes(router)

	// Run Server
	router.Run(":8080")
}

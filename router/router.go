package router

import "github.com/gin-gonic/gin"

func InitializeRouter() {
	// Inicializa o Router utilizando as configurações padrão do Gin
	router := gin.Default()
	// Define uma Rota
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	router.Run(":8080")
}

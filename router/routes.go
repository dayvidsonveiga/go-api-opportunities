package router

import (
	"github.com/dayvidsonveiga/go-api-opportunities/handler"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.GetOpening)
		v1.POST("/opening", handler.CreateOpening)
		v1.DELETE("/opening", handler.DeleteOpening)
		v1.PUT("/opening", handler.PutOpening)
		v1.GET("/openings", handler.GetAllOpenings)
	}
}

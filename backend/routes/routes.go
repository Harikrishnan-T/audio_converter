package routes

import (
	"audio-converter/backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/api/save", controllers.SaveConversion)
	router.GET("/api/files", controllers.GetFiles)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "API is up!"})
	})
}

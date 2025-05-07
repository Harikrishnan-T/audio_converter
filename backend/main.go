package main

import (
	"audio-converter/backend/config"
	"audio-converter/backend/models"
	"audio-converter/backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	config.DB.AutoMigrate(&models.AudioFile{})

	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run(":8080")
}

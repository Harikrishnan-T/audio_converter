package main

import (
	"audio-converter/backend/config"
	"audio-converter/backend/models"
	"audio-converter/backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	config.DB.AutoMigrate(&models.AudioFile{})

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))
	routes.SetupRoutes(router)
	router.Run(":8080")
}

package controllers

import (
	"audio-converter/backend/config"
	"audio-converter/backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SaveConversion(c *gin.Context) {
	var data struct {
		FileName string `json:"fileName"`
		Format   string `json:"format"`
	}

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	file := models.AudioFile{FileName: data.FileName, Format: data.Format, CreatedAt: time.Now()}
	config.DB.Create(&file)
	c.JSON(http.StatusOK, gin.H{"message": "Saved successfully"})
}

func GetFiles(c *gin.Context) {
	var files []models.AudioFile
	config.DB.Find(&files)
	c.JSON(http.StatusOK, files)
}

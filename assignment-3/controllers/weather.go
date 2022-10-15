package controllers

import (
	"assignment3/helpers"
	"assignment3/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetWeather(c *gin.Context) {
	weather := models.Weather{
		Water: helpers.GenerateRandomNumberRange(),
		Wind:  helpers.GenerateRandomNumberRange(),
	}

	status := helpers.GetStatus(weather)

	dataJson := models.Data{
		Status: weather,
	}

	// Create Folder
	folder := "storage/"
	filename := "weather.json"
	fullPath := folder + filename
	helpers.CreateFolder(folder)

	// Write File
	helpers.WriteFile(dataJson, fullPath)

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"status": status,
		"wind":   weather.Wind,
		"water":  weather.Water,
	})

	return
}

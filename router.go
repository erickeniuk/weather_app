package main

import (
	"net/http"

	"github.com/erickeniuk/geography"
	"github.com/erickeniuk/weather"
	"github.com/gin-gonic/gin"
)

func main() {
	// run app
	r := gin.Default()
	r.LoadHTMLGlob("./templates/*.html")
	r.StaticFile("/weather_icon.png", "./assets/weather_icon.png")
	r.StaticFile("/weather.ico", "./assets/weather.ico")
	r.Static("/bootstrap", "./templates/bootstrap")
	r.Static("/css", "./templates/css")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"content": "This is an index page..."})
	})

	r.GET("/weather", func(c *gin.Context) {

		city := c.Query("city")

		latLong, err := geography.GetLatLong(city)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		weatherResponse, err := weather.GetWeather(*latLong)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		weatherDisplay, err := weather.ExtractHourlyWeatherData(city, *weatherResponse)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.HTML(http.StatusOK, "weather.html", weatherDisplay)
	})

	r.Run(":8080")
}

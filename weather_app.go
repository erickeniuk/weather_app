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

	/*
		// Example of a path param
		r.GET("/weather/:city", func(c *gin.Context) {
			city := c.Param("city")
			// ..
		})
	*/
	r.GET("/weather", func(c *gin.Context) {
		city := c.Query("city")
		lat_long, err := geography.Get_Lat_Long(city)
		// check if Get_Lat_Long returns an error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		weather, err := weather.Get_Weather(*lat_long)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"weather": weather})
	})

	r.Run()
}

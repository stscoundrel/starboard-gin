package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stscoundrel/starboard/stars"
)

// Local mapping of stars to avoid Github serialization names.
type StarStruct struct {
	Repository     string `json:"repository"`
	Count          int    `json:"count"`
	Link           string `json:"link"`
	StarGazersLink string `json:"stargazer_link"`
}

func starsEndpoint(c *gin.Context) {
	stars, starError := stars.GetStars(c.Param("username"))

	if starError != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	responseStars := []StarStruct{}

	for _, star := range stars {
		responseStars = append(responseStars, StarStruct(star))
	}
	c.JSON(http.StatusOK, responseStars)
}

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/stars/:username", starsEndpoint)
	}
	router.Run()
}

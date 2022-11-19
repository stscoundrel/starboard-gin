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

type Summary struct {
	Count   int
	Starred []StarStruct
}

func starsEndpointV1(c *gin.Context) {
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

func starsEndpointV2(c *gin.Context) {
	stars, starError := stars.GetStars(c.Param("username"))

	if starError != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	response := Summary{
		0,
		[]StarStruct{},
	}

	for _, star := range stars {
		starredRepo := StarStruct(star)
		response.Count = response.Count + starredRepo.Count
		response.Starred = append(response.Starred, starredRepo)
	}
	c.JSON(http.StatusOK, response)
}

func healthEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/stars/:username", starsEndpointV1)
		v1.GET("health", healthEndpoint)
	}
	v2 := router.Group("/v2")
	{
		v2.GET("/stars/:username", starsEndpointV2)
		v2.GET("health", healthEndpoint)
	}
	router.Run()
}

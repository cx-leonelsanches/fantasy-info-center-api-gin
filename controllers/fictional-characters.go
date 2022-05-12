package controllers

import (
	"fantasy-info-center-api-gin/models"
	"fantasy-info-center-api-gin/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

var fictionalCharacters = []models.FictionalCharacter{
	{ID: 1, Name: "Fred Flintstone"},
	{ID: 2, Name: "Barney Rubble"},
	{ID: 3, Name: "Wilma Flintstone"},
}

func GetFictionalCharacters(c *gin.Context) {
	records := repositories.GetFictionalCharactersFromDatabase()
	c.IndentedJSON(http.StatusOK, records)
}

func GetFictionalCharacterByID(c *gin.Context) {
	fictionalCharacter := repositories.GetFictionalCharacterByID(c.Param("id"))
	c.IndentedJSON(http.StatusOK, fictionalCharacter)
}

func PostFictionalCharacters(c *gin.Context) {
	var newFictionalCharacter models.FictionalCharacter

	if err := c.BindJSON(&newFictionalCharacter); err != nil {
		return
	}

	repositories.InsertFictionalCharacterInDatabaseDatabase(newFictionalCharacter.Name)
	fictionalCharacters = append(fictionalCharacters, newFictionalCharacter)
	c.IndentedJSON(http.StatusCreated, newFictionalCharacter)
}

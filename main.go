package main

import (
	"fantasy-info-center-api-gin/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/fictionalCharacters", controllers.GetFictionalCharacters)
	router.GET("/fictionalCharacters/:id", controllers.GetFictionalCharacterByID)
	router.POST("/fictionalCharacters", controllers.PostFictionalCharacters)

	router.Run("localhost:8081")
}

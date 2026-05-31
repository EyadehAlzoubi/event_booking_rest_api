package main

import (
	"net/http"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/events", getEvents)
	server.POST("/events", createEvents)

	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events := models.GetAll()
	context.JSON(http.StatusOK, events)
}

func createEvents(context *gin.Context) {
	var newEvent models.Event
	err := context.ShouldBindJSON(&newEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"ERROR": err.Error()})
		return
	}

	newEvent.ID = 1
	newEvent.UserID = 1

	newEvent.Save()

	context.JSON(http.StatusCreated, gin.H{"message": "event created successfully", "event": newEvent})

}

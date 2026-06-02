package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAll()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"ERROR1": err.Error()})
		return

	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"ERROR_getEvent": err.Error()})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"ERROR_getEvent": err.Error()})
		return
	}

	context.JSON(http.StatusOK, event)

}

func createEvents(context *gin.Context) {
	var newEvent models.Event
	err := context.ShouldBindJSON(&newEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"ERROR": err.Error()})
		return
	}
	userId := context.GetInt64("userId")
	newEvent.UserID = userId

	err = newEvent.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"ERROR2": err.Error()})
		return

	}

	context.JSON(http.StatusCreated, gin.H{"message": "event created successfully", "event": newEvent})

}

func updateEvent(context *gin.Context) {

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"updateEvent": err.Error()})
		return
	}

	_, err = models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"updateEvent": err.Error()})
		return
	}

	var updateEvent models.Event

	err = context.ShouldBindJSON(&updateEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"ERROR": err.Error()})
		return
	}

	updateEvent.ID = eventId
	err = updateEvent.Update()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"ERROR": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "event Updated successfully"})

}

func deleteEvent(context *gin.Context) {

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"deleteEvent": err.Error()})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"updateEvent": err.Error()})
		return
	}

	err = event.Delete()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"ERROR": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "event Deleted successfully"})

}

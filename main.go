package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"ig4llc.com/models"
)

func main() {
	server := gin.Default()

	//register endpoints as a handler for http request
	server.GET("/events", getEvents) //GET, POST, PUT, PATCH, DELETE
	server.POST("/events", createEvent)

	server.Run(":8081") // localhost:8081
}

func getEvents(context *gin.Context) {

	events := models.GetAllEvents()

	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"}) //err.Error()
		return
	}

	event.ID = 1
	event.UserID = 1
	event.Save()
	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})

	//
	// context.JSON(http.StatusCreated, event)
	// context.ShouldBindJSON(&event)
}

// func updateEvent(context *gin.Context) {

// }

// func deleteEvent(context *gin.Context) {

// }

// func getEvent(context *gin.Context) {

// }

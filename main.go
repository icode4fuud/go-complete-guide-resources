package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"ig4llc.com/db"
	"ig4llc.com/models"
)

func main() {
	if err := db.InitDB(); err != nil {
		//panic(err)
		log.Fatal("Database initialization failed: ", err)
	}
	defer db.CloseDB()

	server := gin.Default()

	//register endpoints as a handler for http request
	server.GET("/events", getEvents) //GET, POST, PUT, PATCH, DELETE
	server.POST("/events", createEvent)

	server.Run(":8081") // localhost:8081
}

func getEvents(context *gin.Context) {

	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

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
	err = event.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not create event. Try again later."}) //err.Error()
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}

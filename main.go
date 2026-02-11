package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"ig4llc.com/db"
	"ig4llc.com/models"
	"ig4llc.com/services"
)

var eventService *services.EventService
var eventRepo *db.EventRepo

func main() {
	if err := db.InitDB(); err != nil {
		//panic(err)
		log.Fatal("Database initialization failed: ", err)
	}
	defer db.CloseDB()

	//initialize the the repo & service layer <= service layer references the repository layer
	eventRepo = db.NewEventRepo()
	eventService = services.NewEventService(eventRepo)

	server := gin.Default()

	//register endpoints as a handler for http request
	server.GET("/events", getEvents) //GET, POST, PUT, PATCH, DELETE
	server.POST("/events", createEvent)

	server.Run(":8081") // localhost:8081
}

// handlers
func getEvents(context *gin.Context) {

	events, err := eventService.GetAllEvents()
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

	//Evolution of Go project structure/architecture
	//err = event.Save() 1st from Udemy
	if err := eventService.CreateEvent(&event); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to INSERT INTO database!"})
		return
	}

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not create event. Try again later."}) //err.Error()
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}

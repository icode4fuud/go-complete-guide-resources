package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	//handler for http request
	server.GET("/events", getEvents) //GET, POST, PUT, PATCH, DELETE

	server.Run(":8081") // localhost:8081
}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
}

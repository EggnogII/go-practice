package main

import (
	"net/http"

	"example.com/rest-project/db"
	"example.com/rest-project/models"
	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()
	server := gin.Default()          // Configure HTTP Server, with logger and recovery attached
	server.GET("/events", getEvents) //GET, POST, PUT, PATCH, DELETE
	server.POST("/events", createEvent)
	server.Run(":8080") // localhost:8080
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.BindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	event.ID = 1
	event.UserID = 1

	event.Save()
	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}

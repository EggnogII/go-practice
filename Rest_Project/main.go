package main

import (
	"net/http"
	"strconv"

	"example.com/rest-project/db"
	"example.com/rest-project/models"
	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()
	server := gin.Default()             // Configure HTTP Server, with logger and recovery attached
	server.GET("/events", getEvents)    //GET, POST, PUT, PATCH, DELETE
	server.GET("/events/:id", getEvent) // : denotes dynamic path
	server.POST("/events", createEvent)
	server.Run(":8080") // localhost:8080
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events."})
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event with supplied ID."})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch event with supplied ID."})
		return
	}

	context.JSON(http.StatusOK, event)
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

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event."})
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}

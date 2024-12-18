package routes

import (
	"example.com/rest-project/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)    //GET, POST, PUT, PATCH, DELETE
	server.GET("/events/:id", getEvent) // : denotes dynamic path

	authenticationGroups := server.Group("/")
	authenticationGroups.Use(middleware.Authenticate)
	authenticationGroups.POST("/events", createEvent)
	authenticationGroups.PUT("/events/:id", updateEvent)
	authenticationGroups.DELETE("/events/:id", deleteEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)
}

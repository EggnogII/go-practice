package main

import (
	"example.com/rest-project/db"
	"example.com/rest-project/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()
	server := gin.Default() // Configure HTTP Server, with logger and recovery attached
	routes.RegisterRoutes(server)
	server.Run(":8080") // localhost:8080
}

package main

import (
	"example.com/event/db"
	"example.com/event/routes"
	"github.com/gin-gonic/gin"
)

func main () {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}


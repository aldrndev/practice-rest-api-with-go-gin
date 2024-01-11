package routes

import (
	"example.com/event/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// Public routes
	server.GET("/events", getEvent)
	server.GET("/events/:id", getEventId)
	server.POST("/signup", signup)
	server.POST("/login", login)

	// Protected routes
	protected := server.Group("/").Use(middlewares.Authenticate)
	protected.POST("/events", createEvent)
	protected.PUT("/events/:id", updateEvent)
	protected.DELETE("/events/:id", deleteEvent)
	protected.POST("/events/:id/register", registerForEvent)
	protected.DELETE("/events/:id/cancel", cancelRegistration)
}

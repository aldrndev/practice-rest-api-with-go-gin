package routes

import (
	"log"
	"net/http"
	"strconv"

	"example.com/event/models"
	"github.com/gin-gonic/gin"
)

func getEvent(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEventId(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed"})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed"})
		return
	}

	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	var event models.Event

	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed"})
		return
	}

	userId := context.GetInt64("userId")

	event.UserID = userId

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Success", "event": event})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		log.Printf("Error parsing event ID: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventById(eventId)

	if err != nil {
		log.Printf("Error finding event: %v", err)
		context.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Failed to update, not authorized"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		log.Printf("Error binding JSON: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		log.Printf("Error executing update: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully", "data": updatedEvent})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Printf("Error parsing event ID: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventById(eventId)
	if err != nil {
		log.Printf("Error finding event: %v", err)
		context.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Failed to delete, not authorized"})
		return
	}

	err = event.Delete()

	if err != nil {
		log.Printf("Error executing update: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "success deleted event"})
}

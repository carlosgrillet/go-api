package routes

import (
	"net/http"
	"time"

	"github.com/carlosgrillet/go-api/models"
	"github.com/carlosgrillet/go-api/utils"
	"github.com/gin-gonic/gin"
)

func addEvent(context *gin.Context) {
  userId := context.GetString("userId")
	var event models.Event
  err := context.ShouldBindJSON(&event)
	if err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"error": "Bad request format"})
		return
	}

  event.ID = utils.GenerateID()
  event.UserID = userId
  event.Timestamp = time.Now()
	go event.Save()
	context.JSON(http.StatusOK, gin.H{"message": "Event created"})
}

func getEvent(context *gin.Context) {
	event, err := models.GetEventById(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, event)
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
  if err != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
	context.JSON(http.StatusOK, events)
}

func updateEvent(context *gin.Context) {
	eventRequested, err := models.GetEventById(context.Param("id"))
  userId := context.GetString("userId")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

  if eventRequested.UserID != userId {
    context.JSON(http.StatusUnauthorized, gin.H{"error": "not event owner"}) 
    return
  }
  go eventRequested.Delete()

  event := models.Event{}
	err = context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Bad request format"})
		return
	}

	go event.Save()
	context.JSON(http.StatusOK, gin.H{"message": "Event updated"})
}

func deleteEvent(context *gin.Context) {
	eventRequested, err := models.GetEventById(context.Param("id"))
  userId := context.GetString("userId")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
  if eventRequested.UserID != userId {
    context.JSON(http.StatusUnauthorized, gin.H{"error": "not event owner"}) 
    return
  }
  go eventRequested.Delete()
	context.JSON(http.StatusOK, gin.H{"message": "Event deleted"})
}

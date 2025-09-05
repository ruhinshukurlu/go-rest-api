package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-rest-api/database"
	"github.com/go-rest-api/models"
)

func main() {
	server := gin.Default()
	database.InitDB()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Some error..."})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event

	err := context.ShouldBindJSON(&event)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	event.ID = 1
	event.UserID = 1

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't create an event..."})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Created Successfully!", "event": event})
}

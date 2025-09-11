package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-rest-api/models"
)

// registerForEvent godoc
// @Summary Register for an event
// @Description Register the authenticated user for the specified event
// @Tags registrations
// @Produce json
// @Security BearerAuth
// @Param id path int true "Event ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /events/{id}/register [post]
func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Some error..."})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Event not found"})
		return
	}

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't register event for the user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registered."})
}

// cancelRegistration godoc
// @Summary Cancel event registration
// @Description Cancel the authenticated user's registration for the specified event
// @Tags registrations
// @Produce json
// @Security BearerAuth
// @Param id path int true "Event ID"
// @Success 204 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /events/{id}/register [delete]
func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Some error..."})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Event not found"})
		return
	}

	err = event.Cancel(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't cancel event for the user"})
		return
	}

	context.JSON(http.StatusNoContent, gin.H{"message": "Canceled."})

}

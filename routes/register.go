package routes

import (
	"net/http"
	"project/GoBooking/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Add registration to the database
func registerForEvent(context *gin.Context) {
	//Get the userId from context after being set from the middleware while validating the token
	userId := context.GetInt64("userId")

	//Get the event Id from the URL route
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse URL."})
		return
	}

	//Check if the event with ID = "id" exist
	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return 
	}

	//Add the registration to db
	err = event.SaveRegistration(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Can't register to the event"})
		return 
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registred!"})
}

// Cancel a registartion by removing it from the db
func cancelRegistration(context *gin.Context) {
	//Get the userId from context after being set from the middleware while validating the token
	userId := context.GetInt64("userId")

	//Get the event Id from the URL route
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse URL."})
		return
	}

	var event models.Event
	event.ID = eventId

	//Delete registration from the database
	err = event.CancelRegistration(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't cancel your registration."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registration cancelled."})
}
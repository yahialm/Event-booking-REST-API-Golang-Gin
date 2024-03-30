package routes

import (
	"net/http"
	"project/GoBooking/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

/** gin.Context is a pointer to a struct provided by the Gin framework. It encapsulates the HTTP request and response details for a particular HTTP request.
By passing it as a parameter to the getEvents function, you're allowing the function to access and manipulate the HTTP request and response for the "/events" endpoint. */

// GET all events
func getEvents(context *gin.Context){
	events, err := models.GetAllEvents()

	if err!= nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events from the server." })
		return
	}

	context.JSON(http.StatusOK, events)
}


//GET event by Id
func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Parsing Error. Can't parse the event ID"})
		return
	};

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't fetch event."})
		return
	}

	context.JSON(http.StatusOK, event)
}

//POST a new event
func postEvent(context *gin.Context){

	// Get the event from the request Body
	var e models.Event

	//Convert the body of the request to a useful event STRUCT
	err := context.ShouldBindJSON(&e)

	//Handle errors
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Parsing failed."})
		return
	}

	//Get the userId from the context
	userId := context.GetInt64("userId")

	// We need the user Id for the event, this Id is retrieved using the verify token method
	e.UserID = userId

	//Save to db
	err = e.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event added successfully"})

}

//PUT(UPDATE) an existant event
func updateEvent(context *gin.Context){
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Parsing Error. Can't parse the event ID"})
		return
	};


	//In order to check if the event we want to update exist
	event , err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Event with this ID doesn't exist."})
		return
	}

	// Make sure that only the event creator can update
	userId := context.GetInt64("userId")

	if userId != event.UserID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update this event."})
		return
	}
	
	// A new var to hold the updated event from request body
	var updatedEvent models.Event

	//Convert the body to a useful event STRUCT
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse request data."})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Updated successfully."})

}

//DELETE an event
func deleteEvent(context *gin.Context){

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Parsing Error. Can't parse the event ID"})
		return
	};

	//In order to check if the event we want to update exist
	event , err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Event with this ID doesn't exist."})
		return
	}

	//Make sure that only the event's creator is able to delete
	userId:= context.GetInt64("userId")

	if userId != event.UserID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete this event."})
		return
	}

	//Delete event
	err = event.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Event couldn't be deleted."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully."})
}



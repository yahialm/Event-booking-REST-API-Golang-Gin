package routes

import (
	"project/GoBooking/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	//----------No protected routes---------------

	//Get all events
	server.GET("/events", getEvents)

	//Get a specific event by ID
	server.GET("/events/:id", getEvent)

	//Signp users
	server.POST("/signup", postUser)

	//Login users
	server.POST("/login", loginUser)

	//------------Protected routes-------------

	//Create a group of routers
	server.Group("/events")

	//Set the middleware
	server.Use(middlewares.Authenticate)

	/*Every time a request is sent to the protected routes, a middleware is executed first to check token validity.
	If something went wrong, an error response is generated and no route handler is executed.
	If the middleware was executed successfully (valid token), then handlers will be executed thank to context.Next() */

	// "protected" means that the user should be authenticated

	//Add new event to the database
	server.POST("/events", postEvent)

	//Update an event
	server.PUT("/events/:id", updateEvent)

	//Delete event
	server.DELETE("/events/:id", deleteEvent)

	//Register to an event
	server.POST("/events/:id/register", registerForEvent)

	//Cancel registration
	server.DELETE("/events/:id/register", cancelRegistration)




}
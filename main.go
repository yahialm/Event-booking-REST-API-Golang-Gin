package main

import (
	"project/GoBooking/database"
	"project/GoBooking/routes"

	"github.com/gin-gonic/gin"
)


func main(){
	
	//Initialize the Database
	db.InitDB()

	//Create a basic http server
	server := gin.Default()

	//Use routes handlers
	routes.RegisterRoutes(server)

	server.Run()
}


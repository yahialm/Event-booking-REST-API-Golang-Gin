package routes

import (
	"net/http"
	"project/GoBooking/models"
	"project/GoBooking/utils"

	"github.com/gin-gonic/gin"
)

func postUser(context *gin.Context) {
	var u models.User

	//Convert the body of the request to a useful event STRUCT
	err := context.ShouldBindJSON(&u)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't parse data."})
		return
	}

	err = u.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User registred successfully."})

}

func loginUser(context *gin.Context) {
	
	var u models.User

	//Convert the body of the request to a useful event STRUCT
	err := context.ShouldBindJSON(&u)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't parse data."})
		return
	}

	//Check password validity
	err = u.CheckPassword()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Incorrect email or password."})
		return
	}

	id, err := u.GetIdByEmail()

	// fmt.Println(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Something goes wrong!"})
		return	
	} 

	token, err := utils.GenerateToken(u.Email, id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successfully.", "token": token})

}
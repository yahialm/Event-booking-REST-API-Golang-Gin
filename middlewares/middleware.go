package middlewares

import (
	"net/http"
	"project/GoBooking/utils"

	"github.com/gin-gonic/gin"
)

// A middleware is executed before router's handlers functions are executed
// If any error is returned, the middleware should write the http response and no handler functions are executed (They can't modify the response since it's written once and never modified)

func Authenticate(context *gin.Context) {

	//Get the header for JWT authentication
	token := context.Request.Header.Get("Authorization")

	//Check if the token is empty
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Authorization error"})
		return
	}

	//Check the validity of the token if not empty
	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token or token expired."})
		return
	}

	context.Set("userId", userId)

	context.Next()

}
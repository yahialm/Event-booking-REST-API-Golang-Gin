package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "mySimpleSecretKey"

func GenerateToken(email string, userID int64) (string, error){

	// Create the payload part of the JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"userId": userID,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})

	// Create and return the signed token based on the payload and the header specified above in string format
	// "secret key" must be of byte type
	return token.SignedString([]byte(secretKey))
}


// Checking tokens
func VerifyToken(token string) (int64, error) {

	// "func(token *jwt.Token)" returns the secret key if the signing method is the one expected
	// This part of code is used to parse the token and verify the signature of the token since we got the secret key from "func(token *jwt.Token)"
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

	if !ok {
		return nil, errors.New("unexpected signing method")
	}

		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, errors.New("couldn't parse token")
	}

	//Validate the token by checking expiration time, issuer, ... which are fields from the payload
	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return 0, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("invalid token claims")
	}

	// email := claims["email"].(string)
	// We need the user ID to use it while processing event creation. We have a userId field in each event that should have an ID value of the responsible user of that event creation
	// The map claims strore IDs as floats so we should convert them to int64
	userId := int64(claims["userId"].(float64))

	return userId, err 

}
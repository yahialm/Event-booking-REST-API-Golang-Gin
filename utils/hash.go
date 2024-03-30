package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error){

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 15)
	return string(bytes), err

}

func ComparePasswords(hashedPassword, password string ) bool{
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
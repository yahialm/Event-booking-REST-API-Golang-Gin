package models

import (
	"errors"
	db "project/GoBooking/database"
	utils "project/GoBooking/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user User) Save() error {
	query := "INSERT INTO users(email, password) VALUES(?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedUserPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(user.Email, hashedUserPassword)

	if err != nil {
		return err
	}

	_ , err = result.LastInsertId()

	//user.ID = userId

	return err
}

func (u User) CheckPassword() error{
	query := "SELECT password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	//To hold the retrieved hashed password
	var retrievedhashedPassword string 

	//Get the hashed password
	err := row.Scan(&retrievedhashedPassword)

	if err != nil {
		return err
	}

	//Compare the hashed version retrieved with the one obtained from the POST request 
	passwordIsValid := utils.ComparePasswords(retrievedhashedPassword, u.Password)

	if !passwordIsValid {
		return errors.New("invalid password, please try again")
	}

	return nil
	
}

func (u User)GetIdByEmail() (int64, error) {
	query := "SELECT id FROM users WHERE email = ?"

	row := db.DB.QueryRow(query, u.Email)

	var retrievedID int64

	err := row.Scan(&retrievedID)

	if err != nil {
		return 0, err
	}

	return retrievedID, nil 
}
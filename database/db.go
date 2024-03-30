package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Creates a variable to hold the connection
var DB *sql.DB

func InitDB() {
	//Open the database for connection and db configuration
	//DB is like a DB handler or a way to reach the database config
	
	var err error
	
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		//Crash the app
		panic("Could not connect to the database.")
	}
	// Max and Idle connections
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(6)

	createTables()

}

func createTables(){

	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`
	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic("Couldn't create users database.")
	}


	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`
	_, err = DB.Exec(createEventsTable)

	if err != nil {
		panic("Database can't be created:" + err.Error())
	}


	createRegistrationsTable := `
	CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY(event_id) REFERENCES events(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createRegistrationsTable)

	if err != nil {
		panic("Database can't be created: " + err.Error())
	}



}
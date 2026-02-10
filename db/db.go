package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

// global variable; uppercase makes it public access modifier
var DB *sql.DB

// initalize the database
func InitDB() {
	DB, err := sql.Open("sqlite", "api.db")

	if err != nil {
		panic("Could not connect to database")
	}

	//configure connection pooling to control how many open cnxn are allowed
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5) // how many cnxn we want to keep open if noone is using a cnxn

	//call createTables() inside InitDB() so your schema exist before queries run
	//Microsoft Copilot is wrong here, this line causes an error
	//createTables()
}

func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		datetime TEXT NOT NULL,
		location TEXT NOT NULL,
		description TEXT NOT NULL,
		user_id INTEGER NOT NULL
	)`

	_, err := DB.Exec(createEventsTable)

	if err != nil {
		panic("Could not create events table")
	}
}

func CloseDB() {
	DB.Close()
}

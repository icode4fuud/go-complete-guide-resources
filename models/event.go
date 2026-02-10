package models

//have all the logic that deals with storing event data

import (
	"time"

	"ig4llc.com/db"
)

// Event struct
type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	Location    string    `binding:"required"`
	Description string    `binding:"required"`
	UserID      int
}

var events []Event = []Event{}

func (e Event) Save() error {
	query :=
		`INSERT INTO events (name, datetime, location, description, user_id) 
	VALUES (?, ?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.DateTime, e.Location, e.Description, e.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.ID = int64(id)
	return err

	//later will add to a database
	//events = append(events, e)
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		//pass a pointer(&)
		rows.Scan(&event.ID, &event.Name, &event.DateTime, &event.Location, &event.Description, &event.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

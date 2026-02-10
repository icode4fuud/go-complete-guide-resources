package models

//have all the logic that deals with storing event data

import "time"

//Event struct
type Event struct {
	ID          int
	Name        string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	Location    string    `binding:"required"`
	Description string    `binding:"required"`
	UserID      int
}

var events []Event = []Event{}

func (e Event) Save() {
	//later will add to a database
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events

}

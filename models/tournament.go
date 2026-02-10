package models

//have all the logic that deals with storing event data

import "time"

//Event struct
type Tournament struct {
	ID          int
	Name        string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	Location    string    `binding:"required"`
	Description string    `binding:"required"`
	UserID      int
}

var tournaments []Tournament = []Tournament{}

func (e Tournament) Save() {
	//later will add to a database
	tournaments = append(tournaments, e)
}

func GetAllTournaments() []Tournament {
	return tournaments

}

package models

//have all the logic that deals with storing event data (Udemy)
// Decouple HTTP from persistence. Microsoft Copilot

import (
	"time"
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

type EventRepository interface {
	Create(e *Event) error
	GetAll() ([]Event, error)
}

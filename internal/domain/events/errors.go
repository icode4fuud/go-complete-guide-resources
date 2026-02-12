package events

import "errors"

var (
	ErrEventNotFound = errors.New("Event not found")
	ErrInvalidEvent  = errors.New("Invalid event data")
	// ErrInvalidEventID = errors.New("Invalid event ID")
	// ErrDMLFailed      = errors.New("Database persistence failed!")
)

package events

import "errors"

var (
	ErrEventNotFound = errors.New("Event not found")
	ErrInvalidEvent  = errors.New("Invalid event data")
	ErrUnauthorized  = errors.New("unauthorized access") // New
	ErrForbidden     = errors.New("forbidden")           // New
)

package services

import(
	ig4llc.com/db
)

type EventService interface { 
	CreateEvent(e *Event) error 
	GetAllEvents() ([]Event, error) 
}
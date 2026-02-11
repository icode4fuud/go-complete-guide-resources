package services

import (
	"ig4llc.com/models"
)

type EventService interface {
	CreateEvent(e *models.Event) error
	GetAllEvents() ([]models.Event, error)
}

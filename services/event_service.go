//Why this is great
//You can add validation, authorization, transformations, etc.
//You can swap repos (e.g., mock repo for tests).
//Handlers stay tiny.

package services

import (
	"ig4llc.com/db"
	"ig4llc.com/models"
)

type EventService struct {
	repo *db.EventRepo
}

func NewEventService(repo *db.EventRepo) *EventService {
	return &EventService{repo: repo}
}

func (s *EventService) CreateEvent(e *models.Event) error {
	// business rules go here later
	return s.repo.Create(e)
}

func (s *EventService) GetAllEvents() ([]models.Event, error) {
	return s.repo.GetAll()
}

func (s *EventService) GetEventByID(id int64) (*models.Event, error) {
	return s.repo.GetByID(id)
}

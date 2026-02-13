// Why this is great
// You can add validation, authorization, transformations, etc.
// You can swap repos (e.g., mock repo for tests).
// Handlers stay tiny.
package events

import (
	"strings"

	//"ig4llc.com/internal/infrastructure/cache"
	"ig4llc.com/internal/infrastructure/logging"
)

type Repository interface {
	Create(e *Event) error
	GetAll() ([]Event, error)
	GetByID(id int64) (*Event, error)
	Update(e *Event) error
	Delete(id int64) error
	Register(eventID int64, userID int) error
	Unregister(eventID int64, userID int) error
}

type Service interface {
	CreateEvent(e *Event) error
	GetAllEvents() ([]Event, error)
	GetEventByID(id int64) (*Event, error)
	UpdateEvent(e *Event) error
	DeleteEvent(id int64) error
	RegisterUser(eventID int64, userID int) error
	UnregisterUser(eventID int64, userID int) error
}

type service struct {
	repo Repository
	//cache *cache.EventsCache
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateEvent(e *Event) error {
	if strings.TrimSpace(e.Name) == "" { // (tightened validation)Handlers stay thin; they just bind DTO → domain and map errors
		return ErrInvalidEvent
	}
	if strings.TrimSpace(e.Location) == "" { // added validation equivalent to FluentValidations
		return ErrInvalidEvent
	}
	if e.DateTime.IsZero() {
		logging.Logger.Printf("method: service.CreateEvent e.DateTime is zero: DateTime=%s", e.DateTime.String())
		return ErrInvalidEvent
	}

	//test logger
	logging.Logger.Printf("method: service.CreateEvent creating event: user=%d name=%s", e.UserID, e.Name)

	return s.repo.Create(e)
}

func (s *service) GetAllEvents() ([]Event, error) {
	return s.repo.GetAll()
}

func (s *service) GetEventByID(id int64) (*Event, error) {
	return s.repo.GetByID(id)
}

func (s *service) UpdateEvent(e *Event) error {
	return s.repo.Update(e)
}

func (s *service) DeleteEvent(id int64) error {
	return s.repo.Delete(id)
}

func (s *service) RegisterUser(eventID int64, userID int) error {
	return s.repo.Register(eventID, userID)
}

func (s *service) UnregisterUser(eventID int64, userID int) error {
	return s.repo.Unregister(eventID, userID)
}

// func (s *service) ListEvents(page, pageSize int, sortBy, order string, userID *int) ([]Event, error) {
// 	if userID == nil && page == 1 && sortBy == "datetime" && strings.ToLower(order) == "desc" {
// 		if data, ok := s.cache.Get(); ok {
// 			return data, nil
// 		}
// 	}

// 	events, err := s.repo.List(page, pageSize, sortBy, order, userID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if userID == nil && page == 1 && sortBy == "datetime" && strings.ToLower(order) == "desc" {
// 		s.cache.Set(events)
// 	}

// 	return events, nil
// }

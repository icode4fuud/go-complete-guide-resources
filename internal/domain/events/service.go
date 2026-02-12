// Why this is great
// You can add validation, authorization, transformations, etc.
// You can swap repos (e.g., mock repo for tests).
// Handlers stay tiny.
package events

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
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateEvent(e *Event) error {
	if e.Name == "" {
		return ErrInvalidEvent
	}
	if e.Location == "" { // added validation equivalent to FluentValidations
		return ErrInvalidEvent
	}
	// if e.DateTime.IsZero() {
	// 	return ErrInvalidEvent
	// }

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

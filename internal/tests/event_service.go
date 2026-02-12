package tests

import (
	"testing"
	"time"

	"ig4llc.com/internal/domain/events"
)

type mockRepo struct {
	created []*events.Event
	list    []events.Event
	err     error
}

func (m *mockRepo) Create(e *events.Event) error {
	if m.err != nil {
		return m.err
	}
	e.ID = 99
	m.created = append(m.created, e)
	return nil
}

func (m *mockRepo) GetAll() ([]events.Event, error) {
	return m.list, m.err
}

func (m *mockRepo) GetByID(id int64) (*events.Event, error) {
	if m.err != nil {
		return nil, m.err
	}
	return &events.Event{ID: id, Name: "Mock"}, nil
}

func (m *mockRepo) Update(e *events.Event) error               { return m.err }
func (m *mockRepo) Delete(id int64) error                      { return m.err }
func (m *mockRepo) Register(eventID int64, userID int) error   { return m.err }
func (m *mockRepo) Unregister(eventID int64, userID int) error { return m.err }

func TestCreateEvent(t *testing.T) {
	mr := &mockRepo{}
	svc := events.NewService(mr)

	e := &events.Event{
		Name:        "Test",
		DateTime:    time.Now(),
		Location:    "Here",
		Description: "Desc",
		UserID:      1,
	}

	if err := svc.CreateEvent(e); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if e.ID != 99 {
		t.Fatalf("expected ID 99, got %d", e.ID)
	}
}

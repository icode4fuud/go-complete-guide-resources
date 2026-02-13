package tests

import (
	"errors"
	"testing"

	. "ig4llc.com/internal/domain/events"
)

type fakeRepo struct {
	created []*Event
}

func (f *fakeRepo) Create(e *Event) error {
	f.created = append(f.created, e)
	return nil
}

func (f *fakeRepo) GetAll() ([]Event, error) {
	return nil, nil
}

func (f *fakeRepo) GetByID(id int64) (*Event, error) {
	return nil, nil
}

func (f *fakeRepo) Update(e *Event) error {
	return nil
}

func (f *fakeRepo) Delete(id int64) error {
	return nil
}

func (f *fakeRepo) Register(eventID int64, userID int) error {
	return nil
}

func (f *fakeRepo) Unregister(eventID int64, userID int) error {
	return nil
}

func TestCreateEvent_Validation(t *testing.T) {
	repo := &fakeRepo{}
	svc := NewService(repo)

	err := svc.CreateEvent(&Event{Name: ""})
	if !errors.Is(err, ErrInvalidEvent) {
		t.Fatalf("expected ErrInvalidEvent, got %v", err)
	}
}

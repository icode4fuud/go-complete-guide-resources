// db/event_repo_test.go
package db_test

import (
	"testing"
	"time"

	"ig4llc.com/db"
	"ig4llc.com/models"
)

func TestEventRepo_CreateAndGetAll(t *testing.T) {
	_ = db.NewTestDB(t)
	repo := db.NewEventRepo()

	e := &models.Event{
		Name:        "Test",
		DateTime:    time.Now(),
		Location:    "Here",
		Description: "Desc",
		UserID:      1,
	}

	if err := repo.Create(e); err != nil {
		t.Fatalf("Create: %v", err)
	}
	if e.ID == 0 {
		t.Fatalf("expected ID to be set")
	}

	events, err := repo.GetAll()
	if err != nil {
		t.Fatalf("GetAll: %v", err)
	}
	if len(events) != 1 {
		t.Fatalf("expected 1 event, got %d", len(events))
	}
}

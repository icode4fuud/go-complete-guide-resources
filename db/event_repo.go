// db/event_repo.go
package db

import (
	"database/sql"
	"time"

	"ig4llc.com/models"
)

type EventRepo struct {
	db *sql.DB
}

func NewEventRepo() *EventRepo {
	return &EventRepo{db: DB}
}

func (r *EventRepo) Create(e *models.Event) error {
	stmt, err := PrepareCached(`
        INSERT INTO events (name, datetime, location, description, user_id)
        VALUES (?, ?, ?, ?, ?)
    `)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(e.Name, e.DateTime.Format(time.RFC3339), e.Location, e.Description, e.UserID)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	e.ID = id
	return nil
}

func (r *EventRepo) GetAll() ([]models.Event, error) {
	rows, err := r.db.Query(`SELECT id, name, datetime, location, description, user_id FROM events`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []models.Event
	for rows.Next() {
		var e models.Event
		var dt string
		if err := rows.Scan(&e.ID, &e.Name, &dt, &e.Location, &e.Description, &e.UserID); err != nil {
			return nil, err
		}
		t, _ := time.Parse(time.RFC3339, dt)
		e.DateTime = t
		events = append(events, e)
	}
	return events, rows.Err()
}

func (r *EventRepo) GetByID(id int64) (*models.Event, error) {
	row := r.db.QueryRow("SELECT * FROM events WHERE id = ?", id)

	var e models.Event
	var dt string

	if err := row.Scan(&e.ID, &e.Name, &dt, &e.Location, &e.Description, &e.UserID); err != nil {
		return nil, err
	}
	t, _ := time.Parse(time.RFC3339, dt)
	e.DateTime = t
	return &e, nil
}

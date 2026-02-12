// db/event_repo.go
package db

import (
	"database/sql"
	"time"

	"ig4llc.com/internal/domain/events"
)

type EventRepo struct {
	db *sql.DB
}

func NewEventRepo(database *sql.DB) *EventRepo {
	return &EventRepo{db: DB}
}

func (r *EventRepo) Create(e *events.Event) error {
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

func (r *EventRepo) GetAll() ([]events.Event, error) {
	rows, err := r.db.Query(`SELECT id, name, datetime, location, description, user_id FROM events`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []events.Event
	for rows.Next() {
		var e events.Event
		var dt string
		if err := rows.Scan(&e.ID, &e.Name, &dt, &e.Location, &e.Description, &e.UserID); err != nil {
			return nil, err
		}
		t, _ := time.Parse(time.RFC3339, dt)
		e.DateTime = t
		out = append(out, e)
	}
	return out, rows.Err()
}

func (r *EventRepo) GetByID(id int64) (*events.Event, error) {
	row := r.db.QueryRow("SELECT * FROM events WHERE id = ?", id)

	var e events.Event
	var dt string

	if err := row.Scan(&e.ID, &e.Name, &dt, &e.Location, &e.Description, &e.UserID); err != nil {
		return nil, err
	}
	t, _ := time.Parse(time.RFC3339, dt)
	e.DateTime = t
	return &e, nil
}

func (r *EventRepo) Register(eventID int64, userID int) error {
	_, err := r.db.Exec(`
        INSERT INTO registrations (event_id, user_id)
        VALUES (?, ?)
    `, eventID, userID)
	return err
}

func (r *EventRepo) Unregister(eventID int64, userID int) error {
	_, err := r.db.Exec(`
        DELETE FROM registrations
        WHERE event_id = ? AND user_id = ?
    `, eventID, userID)
	return err
}

func (r *EventRepo) Update(e *events.Event) error {
	_, err := r.db.Exec(`
        UPDATE events
        SET name = ?, datetime = ?, location = ?, description = ?
        WHERE id = ?
    `, e.Name, e.DateTime.Format(time.RFC3339), e.Location, e.Description, e.ID)
	return err
}

func (r *EventRepo) Delete(id int64) error {
	_, err := r.db.Exec(`DELETE FROM events WHERE id = ?`, id)
	return err
}

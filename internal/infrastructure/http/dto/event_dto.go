package dto

import (
	"time"

	"ig4llc.com/internal/domain/events"
)

type CreateEventRequest struct {
	Name        string    `json:"name" binding:"required"` //Automatic Validation: By using binding:"required", Gin will reject the request before it even hits your service if dateTime is missing.
	DateTime    time.Time `json:"dateTime" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	Description string    `json:"description"`
	UserID      int64     `json:"userId"`
}

type EventResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	DateTime    string `json:"dateTime"`
	Location    string `json:"location"`
	Description string `json:"description"`
	UserID      int64  `json:"userId"`
}

func ToEventResponse(e *events.Event) EventResponse {
	return EventResponse{
		ID:          e.ID,
		Name:        e.Name,
		DateTime:    e.DateTime.Format(time.RFC3339),
		Location:    e.Location,
		Description: e.Description,
		UserID:      e.UserID,
	}
}

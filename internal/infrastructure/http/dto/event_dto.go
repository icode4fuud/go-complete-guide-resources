package dto

import "ig4llc.com/internal/domain/events"

type CreateEventRequest struct {
	Name        string `json:"name"`
	DateTime    string `json:"dateTime"`
	Location    string `json:"location"`
	Description string `json:"description"`
	UserID      int    `json:"userId"`
}

type EventResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	DateTime    string `json:"dateTime"`
	Location    string `json:"location"`
	Description string `json:"description"`
	UserID      int    `json:"userId"`
}

func ToEventResponse(e *events.Event) EventResponse {
	return EventResponse{
		ID:          e.ID,
		Name:        e.Name,
		DateTime:    e.DateTime.Format("2006-01-02T15:04:05Z"),
		Location:    e.Location,
		Description: e.Description,
		//UserID:      e.UserID,
	}
}

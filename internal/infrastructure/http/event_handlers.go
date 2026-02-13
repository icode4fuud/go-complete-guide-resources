package http

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"ig4llc.com/internal/domain/events"
	"ig4llc.com/internal/infrastructure/http/dto"
	"ig4llc.com/internal/infrastructure/http/middleware"
	"ig4llc.com/internal/infrastructure/logging"
	//"ig4llc.com/internal/infrastructure/logging"
)

const baseuri = "/events"
const baseuri2 = "/events/:id"
const baseuri3 = "/events/:id/register"

type EventHandler struct {
	svc events.Service
}

func NewEventHandler(svc events.Service) *EventHandler {
	return &EventHandler{svc: svc}
}

func (h *EventHandler) RegisterRoutes(r *gin.Engine) {
	//register authentication middleware
	api := r.Group(baseuri)
	api.Use(middleware.AuthRequired())
	api.Use(middleware.RateLimit(50, time.Minute)) // register rate limiting middleware

	r.GET(baseuri, h.getEvents)
	r.GET(baseuri2, h.getEvent)
	r.POST(baseuri, h.createEvent)
	r.PUT(baseuri2, h.updateEvent)
	r.DELETE(baseuri2, h.deleteEvent)
	r.POST(baseuri3, h.register)
	r.DELETE(baseuri3, h.unregister)
}

// helper function for error mapping
// Use this to swap the repetitive switch logic and manual c.JSON calls
func (h *EventHandler) mapErrorToResponse(c *gin.Context, err error, defaultMsg string) {
	// Log the actual error for internal debugging
	logging.Logger.Printf("Error occurred: %v", err)

	switch err {
	case events.ErrInvalidEvent:
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	case events.ErrEventNotFound:
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	// Add other domain errors here as you create them
	case events.ErrUnauthorized:
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	case events.ErrForbidden:
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"message": defaultMsg})
	}
}

// handler implementation
func (h *EventHandler) getEvents(c *gin.Context) {
	eventz, err := h.svc.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve events"})
		return
	}

	c.JSON(http.StatusOK, eventz)
}

func (h *EventHandler) getEvent(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		h.mapErrorToResponse(c, events.ErrInvalidEvent, "Invalid event ID") //[cite: 1]
		return
	}

	event, err := h.svc.GetEventByID(id)
	if err != nil {
		h.mapErrorToResponse(c, err, "Event not found") //[cite: 1]
		return
	}

	c.JSON(http.StatusOK, event)
}

// You should also update your createEvent handler so that the UserID in the domain model comes from the authenticated session,
// not just the JSON payload. This prevents users from "impersonating" other user IDs in the request body.
func (h *EventHandler) createEvent(c *gin.Context) {
	var req dto.CreateEventRequest // init DTO
	//var event events.Event

	// 1. Bind JSON directly to the DTO
	// If "dateTime" is missing or invalid in the JSON, this will return an error automatically.
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body", "error": err.Error()})
		return
	}

	// 1. Safely retrieve userID from context
	val, exists := c.Get("userId")
	var finalUserID int64

	if !exists {
		// Fallback for development until Auth is implemented
		logging.Logger.Println("Auth not yet implemented: using default userID 1")
		finalUserID = 1
	} else {
		// 2. Safely assert type
		id, ok := val.(int64)
		if !ok {
			h.mapErrorToResponse(c, events.ErrUnauthorized, "Invalid user session")
			return
		}
		finalUserID = id
	}

	//add logic for DTO+clean response
	// 2. Map DTO to Domain Model
	// Because req.DateTime is now a time.Time, no manual time.Parse is needed.
	//logging.Logger.Printf("ev := &events.Event is executing: dateTime=%s", req.DateTime.String())
	ev := &events.Event{
		Name:        req.Name,
		DateTime:    req.DateTime,
		Location:    req.Location,
		Description: req.Description,
		UserID:      int64(finalUserID),
	}

	//3, pass to the service
	if err := h.svc.CreateEvent(ev); err != nil { //[cite: 1]
		h.mapErrorToResponse(c, err, "Failed to insert into database")
		return
	}

	//c.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
	c.JSON(http.StatusCreated, dto.ToEventResponse(ev)) //return DTO ToEventResponse

}

func (h *EventHandler) updateEvent(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		h.mapErrorToResponse(c, events.ErrInvalidEvent, "Invalid ID format") //[cite: 1]
		return
	}

	var event events.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event payload"})
		return
	}
	event.ID = id

	if err := h.svc.UpdateEvent(&event); err != nil {
		h.mapErrorToResponse(c, err, "Failed to update event in database") //[cite: 1]
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event updated", "event": event})
}

func (h *EventHandler) deleteEvent(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		h.mapErrorToResponse(c, events.ErrInvalidEvent, "Invalid ID format") //[cite: 1]
		return
	}

	if err := h.svc.DeleteEvent(id); err != nil {
		h.mapErrorToResponse(c, err, "Failed to delete event") //[cite: 1]
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event deleted"})
}

func (h *EventHandler) register(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		h.mapErrorToResponse(c, events.ErrInvalidEvent, "Invalid ID")
		return
	}

	// In the future, you'll get this ID from your Auth middleware/JWT context
	//currentUserID := 1
	//Now retrieving from the middleware.AuthRequired()
	// Retrieve the userID set by the AuthRequired middleware
	// We use c.Get() and type assert it to an int
	userID, exists := c.Get("userId")
	if !exists {
		h.mapErrorToResponse(c, events.ErrUnauthorized, "User context not found")
		return
	}

	if err := h.svc.RegisterUser(id, userID.(int)); err != nil {
		h.mapErrorToResponse(c, err, "Could not register user for event")
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered for event"})
}

func (h *EventHandler) unregister(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	if err := h.svc.UnregisterUser(id, 1); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User unregistered from event"})
}

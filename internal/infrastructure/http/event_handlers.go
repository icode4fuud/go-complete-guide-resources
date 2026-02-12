package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"ig4llc.com/internal/domain/events"
	"ig4llc.com/internal/infrastructure/http/middleware"
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

	r.GET(baseuri, h.getEvents)
	r.GET(baseuri2, h.getEvent)
	r.POST(baseuri, h.createEvent)
	r.PUT(baseuri2, h.updateEvent)
	r.DELETE(baseuri2, h.deleteEvent)
	r.POST(baseuri3, h.register)
	r.DELETE(baseuri3, h.unregister)
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
		//use Error Mapping to errors.go
		switch err {
		case events.ErrEventNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		case events.ErrInvalidEvent:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}
		return
	}

	event, err := h.svc.GetEventByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}

	c.JSON(http.StatusOK, event)
}

func (h *EventHandler) createEvent(c *gin.Context) {
	var event events.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}

	if err := h.svc.CreateEvent(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to INSERT INTO database!"})
		return
	}

	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"message": "Could not create event. Try again later."}) //err.Error()
	// 	return
	// }

	c.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})

}

func (h *EventHandler) updateEvent(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	var event events.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event payload"})
		return
	}

	event.ID = id

	if err := h.svc.UpdateEvent(&event); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event updated", "event": event})
}

func (h *EventHandler) deleteEvent(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	if err := h.svc.DeleteEvent(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event deleted"})
}

func (h *EventHandler) register(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	if err := h.svc.RegisterUser(id, 1); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
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

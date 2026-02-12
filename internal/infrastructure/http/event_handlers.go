package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"ig4llc.com/internal/domain/events"
)

const baseuri = "/events"

type EventHandler struct {
	svc events.Service
}

func NewEventHandler(svc events.Service) *EventHandler {
	return &EventHandler{svc: svc}
}

func (h *EventHandler) RegisterRoutes(r *gin.Engine) {
	r.GET(baseuri, h.getEvents)
	r.GET("/events/:id", h.getEvent)
	r.POST("/events", h.createEvent)
	r.PUT("/events/:id", h.updateEvent)
	r.DELETE("/events/:id", h.deleteEvent)
	r.POST("/events/:id/register", h.register)
	r.DELETE("/events/:id/register", h.unregister)
}

// handler implementation
func (h *EventHandler) getEvents(c *gin.Context) {
	events, err := h.svc.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve events"})
		return
	}

	c.JSON(http.StatusOK, events)
}

func (h *EventHandler) getEvent(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
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

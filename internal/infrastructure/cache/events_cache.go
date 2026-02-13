package cache

import (
	"sync"
	"time"

	"ig4llc.com/internal/domain/events"
)

type EventsCache struct {
	mu     sync.RWMutex
	data   []events.Event
	expiry time.Time
	ttl    time.Duration
}

func NewEventsCache(ttl time.Duration) *EventsCache {
	return &EventsCache{ttl: ttl}
}

func (c *EventsCache) Get() ([]events.Event, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if time.Now().After(c.expiry) {
		return nil, false
	}
	return c.data, true
}

func (c *EventsCache) Set(events []events.Event) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data = events
	c.expiry = time.Now().Add(c.ttl)
}

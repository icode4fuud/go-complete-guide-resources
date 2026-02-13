package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type tokenBucket struct {
	tokens      int
	maxTokens   int
	refillEvery time.Duration
	lastRefill  time.Time
	mu          sync.Mutex
}

func newTokenBucket(max int, refillEvery time.Duration) *tokenBucket {
	return &tokenBucket{
		tokens:      max,
		maxTokens:   max,
		refillEvery: refillEvery,
		lastRefill:  time.Now(),
	}
}

func (b *tokenBucket) allow() bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	now := time.Now()
	if now.Sub(b.lastRefill) >= b.refillEvery {
		b.tokens = b.maxTokens
		b.lastRefill = now
	}

	if b.tokens <= 0 {
		return false
	}
	b.tokens--
	return true
}

func RateLimit(max int, refillEvery time.Duration) gin.HandlerFunc {
	bucket := newTokenBucket(max, refillEvery)

	return func(c *gin.Context) {
		if !bucket.allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "rate limit exceeded"})
			c.Abort()
			return
		}
		c.Next()
	}
}

package crawler

import (
	"sync"
	"time"
)

// RateLimiter controls how many requests can be made in a time window.
type RateLimiter struct {
	mu          sync.Mutex
	lastRequest time.Time
	delay       time.Duration
}

// NewRateLimiter creates a new RateLimiter with specified delay.
func NewRateLimiter(delay time.Duration) *RateLimiter {
	return &RateLimiter{
		delay: delay,
	}
}

// Allow checks whether a new request can be made now.
func (r *RateLimiter) Allow() bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()
	if now.Sub(r.lastRequest) >= r.delay {
		r.lastRequest = now
		return true
	}
	return false
}

package crawler

import "sync"

// Deduplicator prevents crawling the same URL multiple times.
type Deduplicator struct {
	mu      sync.RWMutex
	visited map[string]bool
}

// NewDeduplicator creates a new Deduplicator.
func NewDeduplicator() *Deduplicator {
	return &Deduplicator{
		visited: make(map[string]bool),
	}
}

// IsDuplicate checks and marks the URL if not visited before.
func (d *Deduplicator) IsDuplicate(url string) bool {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.visited[url] {
		return true
	}
	d.visited[url] = true
	return false
}

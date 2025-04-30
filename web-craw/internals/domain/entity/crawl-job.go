package entity

import "time"

// CrawlJob represents a crawling task/session.
type CrawlJob struct {
	JobID       string    `bson:"job_id" json:"job_id"`                       // Unique ID for the job
	Source      string    `bson:"source" json:"source"`                       // Source name (e.g., Guardian)
	StartURL    string    `bson:"start_url" json:"start_url"`                 // Starting URL for crawling
	Status      string    `bson:"status" json:"status"`                       // pending, running, completed, failed
	StartedAt   time.Time `bson:"started_at" json:"started_at"`               // When job started
	CompletedAt time.Time `bson:"completed_at,omitempty" json:"completed_at"` // When job completed
	Errors      []string  `bson:"errors,omitempty" json:"errors"`             // Any errors encountered
}

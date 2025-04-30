package entity

import "time"

// Article represents a news article fetched by the crawler.
type Article struct {
	ID          string    `bson:"_id,omitempty" json:"id"`          // Unique ID (UUID or ObjectID)
	Title       string    `bson:"title" json:"title"`               // Title of the article
	Author      string    `bson:"author,omitempty" json:"author"`   // Author name, optional
	PublishedAt time.Time `bson:"published_at" json:"published_at"` // Original publication timestamp
	Content     string    `bson:"content" json:"content"`           // Full content of the article
	URL         string    `bson:"url" json:"url"`                   // URL where article was found
	Source      string    `bson:"source" json:"source"`             // News source name (e.g., Guardian)
	Tags        []string  `bson:"tags,omitempty" json:"tags"`       // Tags like "Politics", "Technology"
	ScrapedAt   time.Time `bson:"scraped_at" json:"scraped_at"`     // When the article was scraped
}

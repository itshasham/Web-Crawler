package entity

// Source represents a news source supported by the crawler.
type Source struct {
	SourceID    string `bson:"source_id" json:"source_id"`       // Unique ID
	Name        string `bson:"name" json:"name"`                 // e.g., Guardian, BBC
	BaseURL     string `bson:"base_url" json:"base_url"`         // Base URL for crawling
	AdapterName string `bson:"adapter_name" json:"adapter_name"` // Adapter module name (GuardianAdapter etc.)
}

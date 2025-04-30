package adapters

import (
	"context"
	"web-craw/internals/domain/entity"
)

// SiteAdapter defines an interface for crawling and parsing a site.
type SiteAdapter interface {
	// Parse parses a page and returns an Article, list of new URLs, or an error.
	Parse(ctx context.Context, url string) (*entity.Article, []string, error)
}

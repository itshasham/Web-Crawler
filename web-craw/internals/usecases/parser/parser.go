package parser

import (
	"context"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// FetchDocument fetches and parses a web page into a goquery document.
func FetchDocument(ctx context.Context, url string) (*goquery.Document, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return goquery.NewDocumentFromReader(resp.Body)
}

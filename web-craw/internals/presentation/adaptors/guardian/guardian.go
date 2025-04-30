package guardian

import (
	"context"
	"net/http"
	"web-craw/internals/domain/entity"
	adapters "web-craw/internals/presentation/adaptors"

	"github.com/PuerkitoBio/goquery"
)

type GuardianAdapter struct{}

func NewGuardianAdapter() adapters.SiteAdapter {
	return &GuardianAdapter{}
}

func (g *GuardianAdapter) Parse(ctx context.Context, url string) (*entity.Article, []string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	title := doc.Find("h1").First().Text()
	content := doc.Find("div.content__article-body").Text()

	article := &entity.Article{
		Title:   title,
		Content: content,
		URL:     url,
		Source:  "Guardian",
	}

	var newUrls []string
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists && href != "" {
			newUrls = append(newUrls, href)
		}
	})

	return article, newUrls, nil
}

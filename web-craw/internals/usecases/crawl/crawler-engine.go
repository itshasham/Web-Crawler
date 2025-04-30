package crawler

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"
	"web-craw/internals/domain/entity"
	"web-craw/internals/domain/repository"
	adapters "web-craw/internals/presentation/adaptors"
)

type Engine struct {
	articleRepo  repository.ArticleRepository
	crawlJobRepo repository.CrawlJobRepository
	sourceRepo   repository.SourceRepository
	siteAdapters map[string]adapters.SiteAdapter
	rateLimiter  *RateLimiter
	deduplicator *Deduplicator
}

// NewEngine creates a new Crawler Engine.
func NewEngine(
	articleRepo repository.ArticleRepository,
	crawlJobRepo repository.CrawlJobRepository,
	sourceRepo repository.SourceRepository,
	siteAdapters map[string]adapters.SiteAdapter,
	rateLimiter *RateLimiter,
	deduplicator *Deduplicator,
) *Engine {
	return &Engine{
		articleRepo:  articleRepo,
		crawlJobRepo: crawlJobRepo,
		sourceRepo:   sourceRepo,
		siteAdapters: siteAdapters,
		rateLimiter:  rateLimiter,
		deduplicator: deduplicator,
	}
}

// StartCrawl starts crawling a given source.
func (e *Engine) StartCrawl(ctx context.Context, sourceName string) (string, error) {
	source, err := e.sourceRepo.FindByName(ctx, sourceName)
	if err != nil {
		return "", err
	}

	if source == nil {
		return "", errors.New("source not found: " + sourceName) // ✅ Safe check added here
	}

	adapter, ok := e.siteAdapters[source.AdapterName]
	if !ok {
		return "", errors.New("no adapter found for source: " + sourceName)
	}

	job := &entity.CrawlJob{
		JobID:     generateUUID(),
		Source:    source.Name,
		StartURL:  source.BaseURL,
		Status:    "running",
		StartedAt: time.Now(),
	}
	err = e.crawlJobRepo.Create(ctx, job)
	if err != nil {
		return "", err
	}

	go e.crawl(ctx, job, adapter)

	return job.JobID, nil
}

func (e *Engine) crawl(ctx context.Context, job *entity.CrawlJob, adapter adapters.SiteAdapter) {
	defer func() {
		e.crawlJobRepo.UpdateStatus(ctx, job.JobID, "completed")
	}()

	urls := []string{job.StartURL}
	var wg sync.WaitGroup

	for _, url := range urls {
		if e.deduplicator.IsDuplicate(url) {
			continue
		}

		if !e.rateLimiter.Allow() {
			log.Println("Rate limited, sleeping...")
			time.Sleep(time.Second)
		}

		wg.Add(1)
		go func(url string) {
			defer wg.Done()

			article, newUrls, err := adapter.Parse(ctx, url)
			if err != nil {
				log.Printf("Error parsing %s: %v", url, err)
				e.crawlJobRepo.AppendError(ctx, job.JobID, err.Error())
				return
			}

			if article != nil {
				err = e.articleRepo.Save(ctx, article)
				if err != nil {
					log.Printf("Error saving article: %v", err)
				}
			}

			urls = append(urls, newUrls...)
		}(url)
	}

	wg.Wait()
}

func generateUUID() string {
	// Replace with your preferred UUID generator
	return time.Now().Format("20060102150405")
}

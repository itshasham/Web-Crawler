package repository

import (
	"context"
	"web-craw/internals/domain/entity"
)

// CrawlJobRepository defines the methods for managing crawl jobs.
type CrawlJobRepository interface {
	Create(ctx context.Context, job *entity.CrawlJob) error
	UpdateStatus(ctx context.Context, jobID string, status string) error
	AppendError(ctx context.Context, jobID string, errorMessage string) error
	FindByID(ctx context.Context, jobID string) (*entity.CrawlJob, error)
}

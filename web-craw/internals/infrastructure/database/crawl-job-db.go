package database

import (
	"context"
	"web-craw/internals/domain/entity"
	"web-craw/internals/domain/repository"
)

type InMemoryCrawlJobRepo struct {
	jobs map[string]*entity.CrawlJob
}

func NewInMemoryCrawlJobRepo() repository.CrawlJobRepository {
	return &InMemoryCrawlJobRepo{
		jobs: make(map[string]*entity.CrawlJob),
	}
}

func (r *InMemoryCrawlJobRepo) Create(ctx context.Context, job *entity.CrawlJob) error {
	r.jobs[job.JobID] = job
	return nil
}

func (r *InMemoryCrawlJobRepo) UpdateStatus(ctx context.Context, jobID string, status string) error {
	if job, ok := r.jobs[jobID]; ok {
		job.Status = status
	}
	return nil
}

func (r *InMemoryCrawlJobRepo) AppendError(ctx context.Context, jobID string, errorMessage string) error {
	if job, ok := r.jobs[jobID]; ok {
		job.Errors = append(job.Errors, errorMessage)
	}
	return nil
}

func (r *InMemoryCrawlJobRepo) FindByID(ctx context.Context, jobID string) (*entity.CrawlJob, error) {
	job, ok := r.jobs[jobID]
	if !ok {
		return nil, nil
	}
	return job, nil
}

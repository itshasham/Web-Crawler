package repository

import (
	"context"
	"web-craw/internals/domain/entity"
)

// SourceRepository defines the methods for managing news sources.
type SourceRepository interface {
	Save(ctx context.Context, source *entity.Source) error
	FindByName(ctx context.Context, name string) (*entity.Source, error)
	FindAll(ctx context.Context) ([]*entity.Source, error)
	DeleteByID(ctx context.Context, sourceID string) error
}

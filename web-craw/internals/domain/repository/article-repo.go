package repository

import (
	"context"
	"web-craw/internals/domain/entity"
)

// ArticleRepository defines the methods for article persistence.
type ArticleRepository interface {
	Save(ctx context.Context, article *entity.Article) error
	FindByID(ctx context.Context, id string) (*entity.Article, error)
	FindBySource(ctx context.Context, source string) ([]*entity.Article, error)
	Search(ctx context.Context, filters map[string]interface{}) ([]*entity.Article, error)
	DeleteByID(ctx context.Context, id string) error
}

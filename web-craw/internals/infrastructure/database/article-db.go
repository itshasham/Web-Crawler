package database

import (
	"context"
	"web-craw/internals/domain/entity"
	"web-craw/internals/domain/repository"
)

type InMemoryArticleRepo struct {
	articles map[string]*entity.Article
}

func NewInMemoryArticleRepo() repository.ArticleRepository {
	return &InMemoryArticleRepo{
		articles: make(map[string]*entity.Article),
	}
}

func (r *InMemoryArticleRepo) Save(ctx context.Context, article *entity.Article) error {
	r.articles[article.ID] = article
	return nil
}

func (r *InMemoryArticleRepo) FindByID(ctx context.Context, id string) (*entity.Article, error) {
	article, ok := r.articles[id]
	if !ok {
		return nil, nil
	}
	return article, nil
}

func (r *InMemoryArticleRepo) FindBySource(ctx context.Context, source string) ([]*entity.Article, error) {
	var result []*entity.Article
	for _, article := range r.articles {
		if article.Source == source {
			result = append(result, article)
		}
	}
	return result, nil
}

func (r *InMemoryArticleRepo) Search(ctx context.Context, filters map[string]interface{}) ([]*entity.Article, error) {
	var result []*entity.Article
	for _, article := range r.articles {
		match := true
		if source, ok := filters["source"]; ok && article.Source != source {
			match = false
		}
		if match {
			result = append(result, article)
		}
	}
	return result, nil
}

func (r *InMemoryArticleRepo) DeleteByID(ctx context.Context, id string) error {
	delete(r.articles, id)
	return nil
}

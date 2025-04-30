package database

import (
	"context"
	"web-craw/internals/domain/entity"
	"web-craw/internals/domain/repository"
)

type InMemorySourceRepo struct {
	sources map[string]*entity.Source
}

func NewInMemorySourceRepo() repository.SourceRepository {
	return &InMemorySourceRepo{
		sources: make(map[string]*entity.Source),
	}
}

func (r *InMemorySourceRepo) Save(ctx context.Context, source *entity.Source) error {
	r.sources[source.SourceID] = source
	return nil
}

func (r *InMemorySourceRepo) FindByName(ctx context.Context, name string) (*entity.Source, error) {
	for _, src := range r.sources {
		if src.Name == name {
			return src, nil
		}
	}
	return nil, nil
}

func (r *InMemorySourceRepo) FindAll(ctx context.Context) ([]*entity.Source, error) {
	var result []*entity.Source
	for _, src := range r.sources {
		result = append(result, src)
	}
	return result, nil
}

func (r *InMemorySourceRepo) DeleteByID(ctx context.Context, sourceID string) error {
	delete(r.sources, sourceID)
	return nil
}

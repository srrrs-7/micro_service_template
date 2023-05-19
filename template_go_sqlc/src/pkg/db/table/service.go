package domain

import (
	"context"
	"template/pkg/db/table/entity"
	"template/pkg/db/table/repository"
	"template/pkg/db/table/usecase"
)

type Service struct {
	store repository.Store
}

func NewService(
	store repository.Store,
) usecase.Store {
	return &Service{
		store: store,
	}
}

func (s Service) Find(ctx context.Context) ([]*entity.Entity, error) {
	return nil, nil
}

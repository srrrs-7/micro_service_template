package table

import (
	"template/driver/db/table/model"
	"template/pkg/domain"
	"template/util/env"

	"gorm.io/gorm"
)

type Repository struct {
	env *env.Env
	db  *gorm.DB
}

func NewRepository(
	env *env.Env,
	db *gorm.DB,
) domain.Store {
	return &Repository{
		env: env,
		db:  db,
	}
}

func (r *Repository) GetTable() ([]*domain.Entity, error) {
	var models []*model.Table

	entity, err := entityTranslator(models)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (r *Repository) CreateTable() ([]*domain.Entity, error) {
	return nil, nil
}

func (r *Repository) UpdateTable() ([]*domain.Entity, error) {
	return nil, nil
}

func (r *Repository) DeleteTable() ([]*domain.Entity, error) {
	return nil, nil
}

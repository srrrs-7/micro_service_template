package table

import (
	"context"
	"database/sql"
	"template/util/env"
)

type Repository struct {
	env *env.Env
	db  *sql.DB
}

func NewRepository(
	env *env.Env,
	db *sql.DB,
) *Repository {
	return &Repository{
		env: env,
		db:  db,
	}
}

func (r *Repository) Find(ctx context.Context) (authors []*Author, err error) {
	return nil, nil
}
func (r *Repository) Create(ctx context.Context, body map[string]any) error {
	return nil
}
func (r *Repository) Update(ctx context.Context) error {
	return nil
}
func (r *Repository) Delete(ctx context.Context) error {
	return nil
}
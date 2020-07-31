package feature

import (
	"context"
	"database/sql"
)

type Repository interface {
	Method(ctx context.Context) (string, error)
	NewMethod(ctx context.Context) (DTOFeature, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r repository) Method(_ context.Context) (string, error) {
	return "Hello", nil
}

func (r repository) NewMethod(_ context.Context) (DTOFeature, error) {

	return DTOFeature{
		Title:       "asdasd",
		Description: "sdsadasd",
	}, nil
}

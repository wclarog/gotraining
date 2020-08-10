package library

import (
	"context"
	"database/sql"
)

type Repository interface {
	Method(ctx context.Context) (string, error)
	NewMethod(ctx context.Context) (DTOLibrary, error)
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

func (r repository) NewMethod(_ context.Context) (DTOLibrary, error) {

	return DTOLibrary{
		Title:       "asdasd",
		Description: "sdsadasd",
	}, nil
}

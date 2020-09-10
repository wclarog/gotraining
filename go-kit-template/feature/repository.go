package feature

import (
	"context"
	"go-kit-template/database"
	"go-kit-template/ent"
)

type Repository interface {
	Method(ctx context.Context) (string, error)
	NewMethod(ctx context.Context) (DTOFeature, error)
	database.RepositoryTx
}

type repository struct {
	database.RepositoryTxImpl
}

func NewRepository(client *ent.Client) Repository {
	return &repository{
		database.RepositoryTxImpl{Client: client},
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

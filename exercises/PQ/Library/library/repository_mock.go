package library

import (
	"context"
	"time"
)

var (
	MockMaterials = []DTOMaterial{{
		UniqueCode:     "123",
		Name:           "HOla",
		DateOfEmission: time.Time{},
		NumberOfPages:  12,
		MaterialType:   1,
	}}
)

type repositoryMock struct {
}

func NewRepositoryMock() Repository {
	return repositoryMock{}
}

func (r repositoryMock) GetMaterials(ctx context.Context) ([]DTOMaterial, error) {
	if ctx.Value("GetMaterialsError") == true {
		return nil, ErrDefault
	}
	return MockMaterials, nil
}

func (r repositoryMock) GetMaterialByCode(ctx context.Context, uniqueCode string) (DTOMaterial, error) {
	panic("implement me")
}

func (r repositoryMock) DeleteMaterial(ctx context.Context, uniqueCode string) error {
	panic("implement me")
}

func (r repositoryMock) GetBooks(ctx context.Context) ([]DTOBook, error) {
	panic("implement me")
}

func (r repositoryMock) GetBookByCode(ctx context.Context, uniqueCode string) (DTOBook, error) {
	panic("implement me")
}

func (r repositoryMock) AddBook(ctx context.Context, book DTOBook) (DTOBook, error) {
	panic("implement me")
}

func (r repositoryMock) UpdateBook(ctx context.Context, uniqueCode string, book DTOBook) (DTOBook, error) {
	panic("implement me")
}

func (r repositoryMock) GetNewspapers(ctx context.Context) ([]DTONewspaper, error) {
	panic("implement me")
}

func (r repositoryMock) GetNewspaperByCode(ctx context.Context, uniqueCode string) (DTONewspaper, error) {
	panic("implement me")
}

func (r repositoryMock) AddNewspaper(ctx context.Context, newspaper DTONewspaper) (DTONewspaper, error) {
	panic("implement me")
}

func (r repositoryMock) UpdateNewspaper(ctx context.Context, uniqueCode string, newspaper DTONewspaper) (DTONewspaper, error) {
	panic("implement me")
}

func (r repositoryMock) GetMagazines(ctx context.Context) ([]DTOMagazine, error) {
	panic("implement me")
}

func (r repositoryMock) GetMagazineByCode(ctx context.Context, uniqueCode string) (DTOMagazine, error) {
	panic("implement me")
}

func (r repositoryMock) AddMagazine(ctx context.Context, magazine DTOMagazine) (DTOMagazine, error) {
	panic("implement me")
}

func (r repositoryMock) UpdateMagazine(ctx context.Context, uniqueCode string, magazine DTOMagazine) (DTOMagazine, error) {
	panic("implement me")
}

package library

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type Repository interface {
	GetMaterial(ctx context.Context, uniqueCode string) (interface{}, error)
	AddMaterial(ctx context.Context, material interface{}) (interface{}, error)
	UpdateMaterial(ctx context.Context, material interface{}) (interface{}, error)
	DeleteMaterial(ctx context.Context, uniqueCode string) error
}

type repository struct {
	db       *sql.DB
	testData map[string]interface{}
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db:       db,
		testData: initTestData(),
	}
}

func (r repository) GetMaterial(_ context.Context, uniqueCode string) (interface{}, error) {
	material, exists := r.testData[uniqueCode]

	if !exists {
		return Material{}, errors.New(fmt.Sprintf("material %s not defined in the database", uniqueCode))
	}

	return material, nil
}

func (r repository) AddMaterial(_ context.Context, material interface{}) (interface{}, error) {
	materialElement := material.(Material)

	result, exists := r.testData[materialElement.UniqueCode]

	if exists {
		return Material{}, errors.New(fmt.Sprintf("material %s already defined in the database", materialElement.UniqueCode))
	}

	return result, nil
}

func (r repository) UpdateMaterial(_ context.Context, material interface{}) (interface{}, error) {
	materialElement := material.(Material)

	_, exists := r.testData[materialElement.UniqueCode]

	if !exists {
		return Material{}, errors.New(fmt.Sprintf("material %s not defined in the database", materialElement.UniqueCode))
	}

	r.testData[materialElement.UniqueCode] = material

	return material, nil
}

func (r repository) DeleteMaterial(_ context.Context, uniqueCode string) error {
	_, exists := r.testData[uniqueCode]

	if !exists {
		return errors.New(fmt.Sprintf("material %s not defined in the database", uniqueCode))
	}

	delete(r.testData, uniqueCode)

	return nil
}

func initTestData() map[string]interface{} {
	var idx int
	var typeLabel string
	var code string
	materials := make(map[string]interface{}, 500)

	for idx = 0; idx < 100; idx++ {
		typeLabel = "Book"
		code = initCodeTestData(idx, typeLabel)
		materials[code] = initBookTestData(idx, typeLabel, code)
	}

	for idx = 0; idx < 100; idx++ {
		typeLabel = "Newspaper"
		code = initCodeTestData(idx, typeLabel)
		materials[code] = initNewspaperTestData(idx, typeLabel, code)
	}

	for idx = 0; idx < 100; idx++ {
		typeLabel = "Magazine"
		code = initCodeTestData(idx, typeLabel)
		materials[code] = initMagazineTestData(idx, typeLabel, code)
	}

	return materials
}

func initCodeTestData(idx int, typeLabel string) string {
	code := fmt.Sprintf("%s_%d", typeLabel, idx)

	return code
}

func initMaterialTestData(idx int, typeLabel string, code string) DTOMaterial {
	material := DTOMaterial{
		UniqueCode:     code,
		Name:           fmt.Sprintf("%s %d", typeLabel, idx),
		DateOfEmission: time.Now(),
		NumberOfPages:  500 + idx,
	}

	return material
}

func initBookTestData(idxBook int, typeLabel string, code string) DTOBook {
	book := DTOBook{
		DTOMaterial: initMaterialTestData(idxBook, typeLabel, code),
		AuthorName:  fmt.Sprintf("AuthorName %d", idxBook),
		Genre:       fmt.Sprintf("Genre %d", idxBook),
	}

	return book
}

func initNewspaperTestData(idxNewspaper int, typeLabel string, code string) DTONewspaper {
	magazine := DTONewspaper{
		DTOMaterial: initMaterialTestData(idxNewspaper, typeLabel, code),
		Url:         fmt.Sprintf("Url %d", idxNewspaper),
	}

	return magazine
}

func initMagazineTestData(idxMagazine int, typeLabel string, code string) DTOMagazine {
	magazine := DTOMagazine{
		DTOMaterial: initMaterialTestData(idxMagazine, typeLabel, code),
		Sections:    initSectionsTestData(idxMagazine),
		Url:         fmt.Sprintf("Url %d", idxMagazine),
	}

	return magazine
}

func initSectionsTestData(idxMagazine int) []string {
	sections := make([]string, 5)

	for idx := 0; idx < 5; idx++ {
		sections[idx] = fmt.Sprintf("Sections %d %d", idxMagazine, idx)
	}
	return sections
}

package materials

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/facebook/ent/examples/start/ent"
	"log"
	"time"
)

type Repository interface {
	GetMaterials(ctx context.Context) ([]DTOMaterial, error)
	GetMaterialByCode(ctx context.Context, uniqueCode string) (DTOMaterial, error)
	DeleteMaterial(ctx context.Context, uniqueCode string) error

	GetBooks(ctx context.Context) ([]DTOBook, error)
	GetBookByCode(ctx context.Context, uniqueCode string) (DTOBook, error)
	AddBook(ctx context.Context, book DTOBook) (DTOBook, error)
	UpdateBook(ctx context.Context, uniqueCode string, book DTOBook) (DTOBook, error)

	GetNewspapers(ctx context.Context) ([]DTONewspaper, error)
	GetNewspaperByCode(ctx context.Context, uniqueCode string) (DTONewspaper, error)
	AddNewspaper(ctx context.Context, newspaper DTONewspaper) (DTONewspaper, error)
	UpdateNewspaper(ctx context.Context, uniqueCode string, newspaper DTONewspaper) (DTONewspaper, error)

	GetMagazines(ctx context.Context) ([]DTOMagazine, error)
	GetMagazineByCode(ctx context.Context, uniqueCode string) (DTOMagazine, error)
	AddMagazine(ctx context.Context, magazine DTOMagazine) (DTOMagazine, error)
	UpdateMagazine(ctx context.Context, uniqueCode string, magazine DTOMagazine) (DTOMagazine, error)
}

type repository struct {
	db       *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db:       db,
	}
}

// Materials
func (r repository) GetMaterials(_ context.Context, client *ent.Client) ([]DTOMaterial, error) {
	materials, err := client.Material.
		Query().
		Where(user.NameEQ("a8m")).
		// `Only` fails if no user found,
		// or more than 1 user returned.
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %v", err)
	}
	log.Println("user returned: ", u)
	return materials, nil
}

func (r repository) GetMaterialByCode(_ context.Context, client *ent.Client, uniqueCode string) (DTOMaterial, error) {
	material, err := client..
		Query().
		Where(user.NameEQ("a8m")).
		// `Only` fails if no user found,
		// or more than 1 user returned.
		Only(ctx)
	if err != nil {
		return DTOMaterial{}, fmt.Errorf("failed querying user: %v", err)
	}
	log.Println("user returned: ", u)
	return material, nil

}

func (r repository) DeleteMaterial(_ context.Context, uniqueCode string) error {
	_, exists := r.testData[uniqueCode]

	if !exists {
		return errors.New(fmt.Sprintf("material %s not found in the database", uniqueCode))
	}

	delete(r.testData, uniqueCode)

	return nil
}

// Books

func (r repository) GetBooks(_ context.Context) ([]DTOBook, error) {
	books := make([]DTOBook, len(r.testData))

	idx := 0
	for _, material := range r.testData {
		resMaterial, err := r.getMaterial(material)
		if err != nil {
			return nil, err
		}

		if resMaterial.MaterialType == BookType {
			books[idx] = material.(DTOBook)
			idx++
		}
	}

	return books[:idx], nil
}

func (r repository) GetBookByCode(_ context.Context, uniqueCode string) (DTOBook, error) {
	material, exists := r.testData[uniqueCode]

	if !exists {
		return DTOBook{}, errors.New(fmt.Sprintf("book %s not found in the database", uniqueCode))
	}

	resMaterial, err := r.getMaterial(material)
	if err != nil {
		return DTOBook{}, err
	}

	if resMaterial.MaterialType != BookType {
		return DTOBook{}, errors.New(fmt.Sprintf("material %s is not a book in the database", uniqueCode))
	}

	return material.(DTOBook), nil
}

func (r repository) AddBook(_ context.Context, book DTOBook) (DTOBook, error) {
	uniqueCode := book.UniqueCode

	_, exists := r.testData[uniqueCode]

	if exists {
		return DTOBook{}, errors.New(fmt.Sprintf("book %s already defined in the database", uniqueCode))
	}

	r.testData[uniqueCode] = book

	return book, nil
}

func (r repository) UpdateBook(_ context.Context, uniqueCode string, book DTOBook) (DTOBook, error) {
	objUniqueCode := book.UniqueCode

	if objUniqueCode != uniqueCode {
		return DTOBook{}, errors.New(fmt.Sprintf("parameter unique code %s and book unique code %s do not match", uniqueCode, objUniqueCode))
	}

	material, exists := r.testData[uniqueCode]

	if !exists {
		return DTOBook{}, errors.New(fmt.Sprintf("book %s not found in the database", uniqueCode))
	}

	m := material.(MaterialTyped)
	if m.GetMaterialType() != BookType {
		return DTOBook{}, errors.New(fmt.Sprintf("material %s is not a book in the database", uniqueCode))
	}

	r.testData[uniqueCode] = book

	return book, nil
}

// Newspapers

func (r repository) GetNewspapers(_ context.Context) ([]DTONewspaper, error) {
	newspapers := make([]DTONewspaper, len(r.testData))

	idx := 0
	for _, material := range r.testData {
		resMaterial, err := r.getMaterial(material)
		if err != nil {
			return nil, err
		}

		if resMaterial.MaterialType == NewspaperType {
			newspapers[idx] = material.(DTONewspaper)
			idx++
		}
	}

	return newspapers[:idx], nil
}

func (r repository) GetNewspaperByCode(_ context.Context, uniqueCode string) (DTONewspaper, error) {
	material, exists := r.testData[uniqueCode]

	if !exists {
		return DTONewspaper{}, errors.New(fmt.Sprintf("newspaper %s not found in the database", uniqueCode))
	}

	resMaterial, err := r.getMaterial(material)
	if err != nil {
		return DTONewspaper{}, err
	}

	if resMaterial.MaterialType != NewspaperType {
		return DTONewspaper{}, errors.New(fmt.Sprintf("material %s is not a newspaper in the database", uniqueCode))
	}

	return material.(DTONewspaper), nil
}

func (r repository) AddNewspaper(_ context.Context, newspaper DTONewspaper) (DTONewspaper, error) {
	uniqueCode := newspaper.UniqueCode

	_, exists := r.testData[uniqueCode]

	if exists {
		return DTONewspaper{}, errors.New(fmt.Sprintf("newspaper %s already defined in the database", uniqueCode))
	}

	r.testData[uniqueCode] = newspaper

	return newspaper, nil
}

func (r repository) UpdateNewspaper(_ context.Context, uniqueCode string, newspaper DTONewspaper) (DTONewspaper, error) {
	objUniqueCode := newspaper.UniqueCode

	if objUniqueCode != uniqueCode {
		return DTONewspaper{}, errors.New(fmt.Sprintf("parameter unique code %s and newspaper unique code %s do not match", uniqueCode, objUniqueCode))
	}

	material, exists := r.testData[uniqueCode]

	if !exists {
		return DTONewspaper{}, errors.New(fmt.Sprintf("newspaper %s not found in the database", uniqueCode))
	}

	m := material.(MaterialTyped)
	if m.GetMaterialType() != NewspaperType {
		return DTONewspaper{}, errors.New(fmt.Sprintf("material %s is not a newspaper in the database", uniqueCode))
	}

	r.testData[uniqueCode] = newspaper

	return newspaper, nil
}

// Magazines

func (r repository) GetMagazines(_ context.Context) ([]DTOMagazine, error) {
	magazines := make([]DTOMagazine, len(r.testData))

	idx := 0
	for _, material := range r.testData {
		resMaterial, err := r.getMaterial(material)
		if err != nil {
			return nil, err
		}

		if resMaterial.MaterialType == MagazineType {
			magazines[idx] = material.(DTOMagazine)
			idx++
		}
	}

	return magazines[:idx], nil
}

func (r repository) GetMagazineByCode(_ context.Context, uniqueCode string) (DTOMagazine, error) {
	material, exists := r.testData[uniqueCode]

	if !exists {
		return DTOMagazine{}, errors.New(fmt.Sprintf("magazine %s not found in the database", uniqueCode))
	}

	resMaterial, err := r.getMaterial(material)
	if err != nil {
		return DTOMagazine{}, err
	}

	if resMaterial.MaterialType != MagazineType {
		return DTOMagazine{}, errors.New(fmt.Sprintf("material %s is not a magazine in the database", uniqueCode))
	}

	return material.(DTOMagazine), nil
}

func (r repository) AddMagazine(_ context.Context, magazine DTOMagazine) (DTOMagazine, error) {
	uniqueCode := magazine.UniqueCode

	_, exists := r.testData[uniqueCode]

	if exists {
		return DTOMagazine{}, errors.New(fmt.Sprintf("material %s already defined in the database", uniqueCode))
	}

	r.testData[uniqueCode] = magazine

	return magazine, nil
}

func (r repository) UpdateMagazine(_ context.Context, uniqueCode string, magazine DTOMagazine) (DTOMagazine, error) {
	objUniqueCode := magazine.UniqueCode

	if objUniqueCode != uniqueCode {
		return DTOMagazine{}, errors.New(fmt.Sprintf("parameter unique code %s and magazine unique code %s do not match", uniqueCode, objUniqueCode))
	}

	material, exists := r.testData[uniqueCode]

	if !exists {
		return DTOMagazine{}, errors.New(fmt.Sprintf("magazine %s not found in the database", uniqueCode))
	}

	m := material.(MaterialTyped)
	if m.GetMaterialType() != MagazineType {
		return DTOMagazine{}, errors.New(fmt.Sprintf("material %s is not a magazine in the database", uniqueCode))
	}

	r.testData[uniqueCode] = magazine

	return magazine, nil
}

// Other functions

func (r repository) getMaterial(material MaterialTyped) (DTOMaterial, error) {
	switch material.GetMaterialType() {
	case BookType:
		return material.(DTOBook).DTOMaterial, nil

	case MagazineType:
		return material.(DTOMagazine).DTOMaterial, nil

	case NewspaperType:
		return material.(DTONewspaper).DTOMaterial, nil

	default:
		return DTOMaterial{}, errors.New(fmt.Sprintf("unsupported material type %d in GetMaterials", material.GetMaterialType()))
	}
}

// Create fake test data

func initTestData() map[string]MaterialTyped {
	var idx int
	var typeLabel string
	var code string
	materials := make(map[string]MaterialTyped, 500)

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

func initMaterialTestData(idx int, typeLabel string, code string, materialType MaterialType) DTOMaterial {
	material := DTOMaterial{
		UniqueCode:     code,
		Name:           fmt.Sprintf("%s %d", typeLabel, idx),
		DateOfEmission: time.Now(),
		NumberOfPages:  500 + idx,
		MaterialType:   materialType,
	}

	return material
}

func initBookTestData(idxBook int, typeLabel string, code string) DTOBook {
	book := DTOBook{
		DTOMaterial: initMaterialTestData(idxBook, typeLabel, code, BookType),
		AuthorName:  fmt.Sprintf("AuthorName %d", idxBook),
		Genre:       fmt.Sprintf("Genre %d", idxBook),
	}

	return book
}

func initNewspaperTestData(idxNewspaper int, typeLabel string, code string) DTONewspaper {
	magazine := DTONewspaper{
		DTOMaterial: initMaterialTestData(idxNewspaper, typeLabel, code, NewspaperType),
		Url:         fmt.Sprintf("Url %d", idxNewspaper),
	}

	return magazine
}

func initMagazineTestData(idxMagazine int, typeLabel string, code string) DTOMagazine {
	magazine := DTOMagazine{
		DTOMaterial: initMaterialTestData(idxMagazine, typeLabel, code, MagazineType),
		Sections:    initSectionsTestData(idxMagazine),
		Url:         fmt.Sprintf("Url %d", idxMagazine),
	}

	return magazine
}

func initSectionsTestData(idxMagazine int) []DTOSection {
	sections := make([]DTOSection, 5)

	for idx := 0; idx < 5; idx++ {
		sections[idx].Code = fmt.Sprintf("Section code %d %d", idxMagazine, idx)
		sections[idx].Content = fmt.Sprintf("Section content %d %d", idxMagazine, idx)
	}
	return sections
}

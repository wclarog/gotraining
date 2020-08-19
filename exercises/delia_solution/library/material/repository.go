package material

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	uuid2 "github.com/google/uuid"
	"strconv"
	"time"
)

type Repository interface {
	GetMaterials(ctx context.Context) ([]MaterialDTO, error)
	GetMaterialByCode(ctx context.Context, uniqueCode string) (MaterialDTO, error)
	DeleteMaterial(ctx context.Context, uniqueCode string) error

	GetBooks(ctx context.Context) ([]BookDTO, error)
	GetBookByCode(ctx context.Context, uniqueCode string) (BookDTO, error)
	AddBook(ctx context.Context, book BookDTO) (BookDTO, error)
	UpdateBook(ctx context.Context, uniqueCode string, book BookDTO) (BookDTO, error)

	GetNewspapers(ctx context.Context) ([]NewsPaperDTO, error)
	GetNewspaperByCode(ctx context.Context, uniqueCode string) (NewsPaperDTO, error)
	AddNewspaper(ctx context.Context, newspaper NewsPaperDTO) (NewsPaperDTO, error)
	UpdateNewspaper(ctx context.Context, uniqueCode string, newspaper NewsPaperDTO) (NewsPaperDTO, error)

	GetMagazines(ctx context.Context) ([]MagazineDTO, error)
	GetMagazineByCode(ctx context.Context, uniqueCode string) (MagazineDTO, error)
	AddMagazine(ctx context.Context, magazine MagazineDTO) (MagazineDTO, error)
	UpdateMagazine(ctx context.Context, uniqueCode string, magazine MagazineDTO) (MagazineDTO, error)
}

type repository struct {
	db       *sql.DB
	mockData []GenericMaterial
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db:       db,
		mockData: createMockData(),
	}
}

func (r repository) GetMaterials(ctx context.Context) ([]MaterialDTO, error) {

	materials := make([]MaterialDTO, len(r.mockData))

	for index, material := range r.mockData {
		mat, err := r.getMaterial(material)
		if err != nil {
			return nil, err
		}
		materials[index] = mat
	}

	return materials, nil
}

func (r repository) GetMaterialByCode(ctx context.Context, uniqueCode string) (MaterialDTO, error) {

	item, found := r.findItem(uniqueCode)

	if found == -1 {
		return MaterialDTO{}, errors.New(fmt.Sprintf("Material %s not found.", uniqueCode))
	}

	return item.(MaterialDTO), nil
}

func (r repository) DeleteMaterial(ctx context.Context, uniqueCode string) error {

	_, found := r.findItem(uniqueCode)

	if found == -1 {
		return errors.New(fmt.Sprintf("Material %s not found.", uniqueCode))
	}

	copy(r.mockData[found:], r.mockData[found+1:]) // Shift a[i+1:] left one index.
	r.mockData[len(r.mockData)-1] = MaterialDTO{}
	r.mockData = r.mockData[:len(r.mockData)-1]

	return nil
}

func (r repository) GetBooks(ctx context.Context) ([]BookDTO, error) {

	books := make([]BookDTO, 0, len(r.mockData))

	for _, material := range r.mockData {

		resMaterial, err := r.getMaterial(material)

		if err != nil {
			return nil, err
		}

		if resMaterial.Type == BookType {
			books = append(books, material.(BookDTO))
		}
	}

	return books, nil
}

func (r repository) GetBookByCode(ctx context.Context, uniqueCode string) (BookDTO, error) {

	item, found := r.findItem(uniqueCode)

	if found == -1 {
		return BookDTO{}, errors.New(fmt.Sprintf("Book %s not found.", uniqueCode))
	}

	if item.GetTypeMaterial() != BookType {
		return BookDTO{}, errors.New(fmt.Sprintf("Material %s is not a book.", uniqueCode))
	}

	return r.mockData[found].(BookDTO), nil
}

func (r repository) AddBook(ctx context.Context, book BookDTO) (BookDTO, error) {

	uniqueCode := book.UniqueCode

	_, found := r.findItem(uniqueCode)

	if found != -1 {
		return BookDTO{}, errors.New(fmt.Sprintf("Book %s already exists", uniqueCode))
	}

	r.mockData = append(r.mockData, book)

	return book, nil
}

func (r repository) UpdateBook(ctx context.Context, uniqueCode string, book BookDTO) (BookDTO, error) {

	item, index := r.findItem(uniqueCode)

	if index == -1 {
		return BookDTO{}, errors.New(fmt.Sprintf("Book not found: %s", uniqueCode))
	}

	if item.GetTypeMaterial() != BookType {
		return BookDTO{}, errors.New(fmt.Sprintf("Material %s is not a book.", uniqueCode))
	}

	r.mockData[index] = book

	return book, nil
}

func (r repository) GetNewspapers(ctx context.Context) ([]NewsPaperDTO, error) {

	newspapers := make([]NewsPaperDTO, 0, len(r.mockData))

	for _, material := range r.mockData {

		resMaterial, err := r.getMaterial(material)

		if err != nil {
			return nil, err
		}

		if resMaterial.Type == NewsPaperType {
			newspapers = append(newspapers, material.(NewsPaperDTO))
		}
	}

	return newspapers, nil
}

func (r repository) GetNewspaperByCode(ctx context.Context, uniqueCode string) (NewsPaperDTO, error) {

	item, found := r.findItem(uniqueCode)

	if found == -1 {
		return NewsPaperDTO{}, errors.New(fmt.Sprintf("Material %s not found.", uniqueCode))
	}

	if item.GetTypeMaterial() != NewsPaperType {
		return NewsPaperDTO{}, errors.New(fmt.Sprintf("Material %s is not a book.", uniqueCode))
	}

	return r.mockData[found].(NewsPaperDTO), nil
}

func (r repository) AddNewspaper(ctx context.Context, newspaper NewsPaperDTO) (NewsPaperDTO, error) {

	uniqueCode := newspaper.UniqueCode

	_, found := r.findItem(uniqueCode)

	if found != -1 {
		return NewsPaperDTO{}, errors.New(fmt.Sprintf("Book %s already exists", uniqueCode))
	}

	r.mockData = append(r.mockData, newspaper)

	return newspaper, nil
}

func (r repository) UpdateNewspaper(ctx context.Context, uniqueCode string, newspaper NewsPaperDTO) (NewsPaperDTO, error) {

	item, index := r.findItem(uniqueCode)

	if index == -1 {
		return NewsPaperDTO{}, errors.New(fmt.Sprintf("Book not found: %s", uniqueCode))
	}

	if item.GetTypeMaterial() != NewsPaperType {
		return NewsPaperDTO{}, errors.New(fmt.Sprintf("Material %s is not a book.", uniqueCode))
	}

	r.mockData[index] = newspaper

	return newspaper, nil
}

func (r repository) GetMagazines(ctx context.Context) ([]MagazineDTO, error) {

	magazines := make([]MagazineDTO, 0, len(r.mockData))

	for _, material := range r.mockData {

		resMaterial, err := r.getMaterial(material)

		if err != nil {
			return nil, err
		}

		if resMaterial.Type == NewsPaperType {
			magazines = append(magazines, material.(MagazineDTO))
		}
	}

	return magazines, nil
}

func (r repository) GetMagazineByCode(ctx context.Context, uniqueCode string) (MagazineDTO, error) {

	item, found := r.findItem(uniqueCode)

	if found == -1 {
		return MagazineDTO{}, errors.New(fmt.Sprintf("Material %s not found.", uniqueCode))
	}

	if item.GetTypeMaterial() != MagazineType {
		return MagazineDTO{}, errors.New(fmt.Sprintf("Material %s is not a book.", uniqueCode))
	}

	return r.mockData[found].(MagazineDTO), nil
}

func (r repository) AddMagazine(ctx context.Context, magazine MagazineDTO) (MagazineDTO, error) {

	uniqueCode := magazine.UniqueCode

	_, found := r.findItem(uniqueCode)

	if found != -1 {
		return MagazineDTO{}, errors.New(fmt.Sprintf("Book %s already exists", uniqueCode))
	}

	r.mockData = append(r.mockData, magazine)

	return magazine, nil
}

func (r repository) UpdateMagazine(ctx context.Context, uniqueCode string, magazine MagazineDTO) (MagazineDTO, error) {
	item, index := r.findItem(uniqueCode)

	if index == -1 {
		return MagazineDTO{}, errors.New(fmt.Sprintf("Book not found: %s", uniqueCode))
	}

	if item.GetTypeMaterial() != MagazineType {
		return MagazineDTO{}, errors.New(fmt.Sprintf("Material %s is not a book.", uniqueCode))
	}

	r.mockData[index] = magazine

	return magazine, nil
}

func createMockData() []GenericMaterial {
	mockMaterial := make([]GenericMaterial, 100)

	for i := 0; i < len(mockMaterial); i++ {
		guid := uuid2.New().String()
		// books
		if i < 40 {
			bookName := fmt.Sprintf("%s%d", "book_", i)
			authorName := fmt.Sprintf("%s%d", "author_", i)
			mockMaterial[i] = createBook(guid, bookName, authorName, i*100, "genre")
		}
		// news papers
		if i >= 40 && i < 70 {
			newsPaperName := fmt.Sprintf("%s%d", "newspaper_", i)
			url := fmt.Sprintf("%s%d", "url", i)
			mockMaterial[i] = createNewsPaper(guid, newsPaperName, url, i*100)
		}
		// magazines
		if i >= 70 {
			magazineName := fmt.Sprintf("%s%d", "magazine_", i)
			urlMag := fmt.Sprintf("%s%d", "url", i)
			mockMaterial[i] = createMagazine(guid, magazineName, urlMag, i%10)
		}
	}

	return mockMaterial[:]
}

func createMockMaterial(name string, code string, materialType TypeMaterial, pages int) MaterialDTO {
	material := MaterialDTO{
		UniqueCode:    code,
		Name:          name,
		EmissionDate:  time.Now(),
		NumberOfPages: pages,
		Type:          materialType,
	}

	return material
}

func createBook(code string, bookName string, author string, pages int, genre string) BookDTO {
	book := BookDTO{
		MaterialDTO: createMockMaterial(bookName, code, BookType, pages),
		Author:      author,
		Genre:       genre,
	}

	return book
}

func createNewsPaper(code string, name string, url string, pages int) NewsPaperDTO {
	newsPaper := NewsPaperDTO{
		MaterialDTO: createMockMaterial(name, code, NewsPaperType, pages),
		Url:         url,
	}

	return newsPaper
}

func createMagazine(code string, name string, url string, pages int) MagazineDTO {
	magazine := MagazineDTO{
		MaterialDTO: createMockMaterial(name, code, MagazineType, pages),
		Url:         url,
		Sections:    createSections(),
	}

	return magazine
}

func createSections() []Section {
	sections := make([]Section, 3)

	for i := 0; i < 3; i++ {
		sections[i].Code = strconv.Itoa(i)
		sections[i].Name = "section " + strconv.Itoa(i)
	}
	return sections
}

func (r repository) getMaterial(material GenericMaterial) (MaterialDTO, error) {

	switch material.GetTypeMaterial() {

	case BookType:
		return material.(BookDTO).MaterialDTO, nil

	case MagazineType:
		return material.(MagazineDTO).MaterialDTO, nil

	case NewsPaperType:
		return material.(NewsPaperDTO).MaterialDTO, nil

	default:
		return MaterialDTO{}, errors.New(fmt.Sprintf("Invalid material type %d ", material.GetTypeMaterial()))
	}
}

func (r repository) findItem(code string) (GenericMaterial, int) {
	for index := range r.mockData {
		mat, _ := r.getMaterial(r.mockData[index])
		if mat.UniqueCode == code {
			return mat, index
		}
	}
	return MaterialDTO{}, -1
}

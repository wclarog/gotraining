package library

import (
	"context"
	"errors"
	"github.com/go-kit/kit/log"
)

type Service interface {
	GetMaterials(ctx context.Context) ([]Material, error)
	GetMaterialByCode(ctx context.Context, uniqueCode string) (Material, error)
	DeleteMaterial(ctx context.Context, uniqueCode string) error

	GetBooks(ctx context.Context) ([]Book, error)
	GetBookByCode(ctx context.Context, uniqueCode string) (Book, error)
	AddBook(ctx context.Context, book Book) (Book, error)
	UpdateBook(ctx context.Context, uniqueCode string, book Book) (Book, error)

	GetNewspapers(ctx context.Context) ([]Newspaper, error)
	GetNewspaperByCode(ctx context.Context, uniqueCode string) (Newspaper, error)
	AddNewspaper(ctx context.Context, newspaper Newspaper) (Newspaper, error)
	UpdateNewspaper(ctx context.Context, uniqueCode string, newspaper Newspaper) (Newspaper, error)

	GetMagazines(ctx context.Context) ([]Magazine, error)
	GetMagazineByCode(ctx context.Context, uniqueCode string) (Magazine, error)
	AddMagazine(ctx context.Context, magazine Magazine) (Magazine, error)
	UpdateMagazine(ctx context.Context, uniqueCode string, magazine Magazine) (Magazine, error)

	dtoToMaterials(ms []DTOMaterial) ([]Material, error)
	dtoToMaterial(m DTOMaterial) (Material, error)
	dtoToBooks(bs []DTOBook) ([]Book, error)
	dtoToBook(b DTOBook) (Book, error)
	dtoToMagazines(ms []DTOMagazine) ([]Magazine, error)
	dtoToMagazine(m DTOMagazine) (Magazine, error)
	dtoToNewspapers(ns []DTONewspaper) ([]Newspaper, error)
	dtoToNewspaper(n DTONewspaper) (Newspaper, error)
}

type service struct {
	repository Repository
	logger     log.Logger
}

func NewService(r Repository, logger log.Logger) Service {
	return &service{
		repository: r,
		logger:     logger,
	}
}

// Materials

func (s service) GetMaterials(ctx context.Context) ([]Material, error) {
	dtoResults, err := s.repository.GetMaterials(ctx)
	if err != nil {
		return nil, err
	}

	return s.dtoToMaterials(dtoResults)
}

func (s service) GetMaterialByCode(ctx context.Context, uniqueCode string) (Material, error) {
	dtoResult, err := s.repository.GetMaterialByCode(ctx, uniqueCode)
	if err != nil {
		return Material{}, err
	}

	return s.dtoToMaterial(dtoResult)
}

func (s service) DeleteMaterial(ctx context.Context, uniqueCode string) error {
	err := s.repository.DeleteMaterial(ctx, uniqueCode)

	return err
}

// Books

func (s service) GetBooks(ctx context.Context) ([]Book, error) {
	dtoBooks, err := s.repository.GetBooks(ctx)
	if err != nil {
		return nil, err
	}

	return s.dtoToBooks(dtoBooks)
}

func (s service) GetBookByCode(ctx context.Context, uniqueCode string) (Book, error) {
	dtoBook, err := s.repository.GetBookByCode(ctx, uniqueCode)
	if err != nil {
		return Book{}, err
	}

	return s.dtoToBook(dtoBook)
}

func (s service) AddBook(ctx context.Context, book Book) (Book, error) {
	var dtoResult DTOBook

	dtoBook, err := s.bookToDto(book)
	if err != nil {
		return Book{}, err
	}

	dtoResult, err = s.repository.AddBook(ctx, dtoBook)
	if err != nil {
		return Book{}, err
	}

	return s.dtoToBook(dtoResult)
}

func (s service) UpdateBook(ctx context.Context, uniqueCode string, book Book) (Book, error) {
	var dtoResult DTOBook

	dtoBook, err := s.bookToDto(book)
	if err != nil {
		return Book{}, err
	}

	dtoResult, err = s.repository.UpdateBook(ctx, uniqueCode, dtoBook)
	if err != nil {
		return Book{}, err
	}

	return s.dtoToBook(dtoResult)
}

// Magazines

func (s service) GetMagazines(ctx context.Context) ([]Magazine, error) {
	dtoMagazines, err := s.repository.GetMagazines(ctx)
	if err != nil {
		return nil, err
	}

	return s.dtoToMagazines(dtoMagazines)
}

func (s service) GetMagazineByCode(ctx context.Context, uniqueCode string) (Magazine, error) {
	dtoMagazine, err := s.repository.GetMagazineByCode(ctx, uniqueCode)
	if err != nil {
		return Magazine{}, err
	}

	return s.dtoToMagazine(dtoMagazine)
}

func (s service) AddMagazine(ctx context.Context, magazine Magazine) (Magazine, error) {
	var dtoResult DTOMagazine

	dtoMagazine, err := s.magazineToDto(magazine)
	if err != nil {
		return Magazine{}, err
	}

	dtoResult, err = s.repository.AddMagazine(ctx, dtoMagazine)
	if err != nil {
		return Magazine{}, err
	}

	return s.dtoToMagazine(dtoResult)
}

func (s service) UpdateMagazine(ctx context.Context, uniqueCode string, magazine Magazine) (Magazine, error) {
	var dtoResult DTOMagazine

	dtoMagazine, err := s.magazineToDto(magazine)
	if err != nil {
		return Magazine{}, err
	}

	dtoResult, err = s.repository.UpdateMagazine(ctx, uniqueCode, dtoMagazine)
	if err != nil {
		return Magazine{}, err
	}

	return s.dtoToMagazine(dtoResult)
}

// Newspapers

func (s service) GetNewspapers(ctx context.Context) ([]Newspaper, error) {
	dtoNewspapers, err := s.repository.GetNewspapers(ctx)
	if err != nil {
		return nil, err
	}

	return s.dtoToNewspapers(dtoNewspapers)
}

func (s service) GetNewspaperByCode(ctx context.Context, uniqueCode string) (Newspaper, error) {
	dtoNewspaper, err := s.repository.GetNewspaperByCode(ctx, uniqueCode)
	if err != nil {
		return Newspaper{}, err
	}

	return s.dtoToNewspaper(dtoNewspaper)
}

func (s service) AddNewspaper(ctx context.Context, newspaper Newspaper) (Newspaper, error) {
	var dtoResult DTONewspaper

	dtoNewspaper, err := s.newspaperToDto(newspaper)
	if err != nil {
		return Newspaper{}, err
	}

	dtoResult, err = s.repository.AddNewspaper(ctx, dtoNewspaper)
	if err != nil {
		return Newspaper{}, err
	}

	return s.dtoToNewspaper(dtoResult)
}

func (s service) UpdateNewspaper(ctx context.Context, uniqueCode string, newspaper Newspaper) (Newspaper, error) {
	var dtoResult DTONewspaper

	dtoNewspaper, err := s.newspaperToDto(newspaper)
	if err != nil {
		return Newspaper{}, err
	}

	dtoResult, err = s.repository.UpdateNewspaper(ctx, uniqueCode, dtoNewspaper)
	if err != nil {
		return Newspaper{}, err
	}

	return s.dtoToNewspaper(dtoResult)
}

// Materials

func (s service) materialToDto(m Material) (DTOMaterial, error) {
	return DTOMaterial{
		UniqueCode:     m.UniqueCode,
		Name:           m.Name,
		DateOfEmission: m.DateOfEmission,
		NumberOfPages:  m.NumberOfPages,
		MaterialType:   m.MaterialType,
	}, nil
}

func (s service) materialsToDto(ms []Material) ([]DTOMaterial, error) {
	var err error
	dtoMaterials := make([]DTOMaterial, len(ms))

	for idx, m := range ms {
		dtoMaterials[idx], err = s.materialToDto(m)

		if err != nil {
			return nil, errors.New("invalid material object in materialsToDto")
		}
	}

	return dtoMaterials, nil
}

func (s service) dtoToMaterial(m DTOMaterial) (Material, error) {
	return Material{
		UniqueCode:     m.UniqueCode,
		Name:           m.Name,
		DateOfEmission: m.DateOfEmission,
		NumberOfPages:  m.NumberOfPages,
		MaterialType:   m.MaterialType,
	}, nil
}

func (s service) dtoToMaterials(ms []DTOMaterial) ([]Material, error) {
	var err error
	materials := make([]Material, len(ms))

	for idx, m := range ms {
		materials[idx], err = s.dtoToMaterial(m)

		if err != nil {
			return nil, errors.New("invalid material object in dtoToMaterials")
		}

	}

	return materials, nil
}

// Books

func (s service) bookToDto(b Book) (DTOBook, error) {
	return DTOBook{
		DTOMaterial: DTOMaterial{
			UniqueCode:     b.UniqueCode,
			Name:           b.Name,
			DateOfEmission: b.DateOfEmission,
			NumberOfPages:  b.NumberOfPages,
			MaterialType:   b.MaterialType,
		},
		AuthorName: b.AuthorName,
		Genre:      b.Genre,
	}, nil
}

func (s service) booksToDto(bs []Book) ([]DTOBook, error) {
	var err error
	dtoBooks := make([]DTOBook, len(bs))

	for idx, b := range bs {
		dtoBooks[idx], err = s.bookToDto(b)

		if err != nil {
			return nil, errors.New("invalid book object in booksToDto")
		}
	}

	return dtoBooks, nil
}

func (s service) dtoToBook(b DTOBook) (Book, error) {
	return Book{
		Material: Material{
			UniqueCode:     b.UniqueCode,
			Name:           b.Name,
			DateOfEmission: b.DateOfEmission,
			NumberOfPages:  b.NumberOfPages,
			MaterialType:   b.MaterialType,
		},
		AuthorName: b.AuthorName,
		Genre:      b.Genre,
	}, nil
}

func (s service) dtoToBooks(bs []DTOBook) ([]Book, error) {
	var err error
	books := make([]Book, len(bs))

	for idx, b := range bs {
		books[idx], err = s.dtoToBook(b)

		if err != nil {
			return nil, errors.New("invalid book object in dtoToBooks")
		}
	}

	return books, nil
}

// Magazines

func (s service) magazineToDto(m Magazine) (DTOMagazine, error) {
	secs, err := s.sectionsToDto(m.Sections)
	if err != nil {
		return DTOMagazine{}, err
	}

	return DTOMagazine{
		DTOMaterial: DTOMaterial{
			UniqueCode:     m.UniqueCode,
			Name:           m.Name,
			DateOfEmission: m.DateOfEmission,
			NumberOfPages:  m.NumberOfPages,
			MaterialType:   m.MaterialType,
		},
		Sections: secs,
		Url:      m.Url,
	}, nil
}

func (s service) magazinesToDto(ms []Magazine) ([]DTOMagazine, error) {
	var err error
	dtoMagazines := make([]DTOMagazine, len(ms))

	for idx, m := range ms {
		dtoMagazines[idx], err = s.magazineToDto(m)

		if err != nil {
			return nil, errors.New("invalid magazine object in magazinesToDto")
		}
	}

	return dtoMagazines, nil
}

func (s service) dtoToMagazine(m DTOMagazine) (Magazine, error) {
	secs, err := s.dtoToSections(m.Sections)
	if err != nil {
		return Magazine{}, err
	}

	return Magazine{
		Material: Material{
			UniqueCode:     m.UniqueCode,
			Name:           m.Name,
			DateOfEmission: m.DateOfEmission,
			NumberOfPages:  m.NumberOfPages,
			MaterialType:   m.MaterialType,
		},
		Sections: secs,
		Url:      m.Url,
	}, nil
}

func (s service) dtoToMagazines(ms []DTOMagazine) ([]Magazine, error) {
	var err error
	magazines := make([]Magazine, len(ms))

	for idx, m := range ms {
		magazines[idx], err = s.dtoToMagazine(m)

		if err != nil {
			return nil, errors.New("invalid magazine object in dtoToMagazines")
		}
	}

	return magazines, nil
}

// Sections

func (s service) sectionToDto(sec Section) (DTOSection, error) {
	return DTOSection{
		Code:    sec.Code,
		Content: sec.Content,
	}, nil
}

func (s service) sectionsToDto(secs []Section) ([]DTOSection, error) {
	var err error
	dtoSections := make([]DTOSection, len(secs))

	for idx, sec := range secs {
		dtoSections[idx], err = s.sectionToDto(sec)

		if err != nil {
			return nil, errors.New("invalid section object in sectionsToDto")
		}
	}

	return dtoSections, nil
}

func (s service) dtoToSection(sec DTOSection) (Section, error) {
	return Section{
		Code:    sec.Code,
		Content: sec.Content,
	}, nil
}

func (s service) dtoToSections(secs []DTOSection) ([]Section, error) {
	var err error
	sections := make([]Section, len(secs))

	for idx, sec := range secs {
		sections[idx], err = s.dtoToSection(sec)

		if err != nil {
			return nil, errors.New("invalid section object in dtoToSections")
		}
	}

	return sections, nil
}

// Newspapers

func (s service) newspaperToDto(n Newspaper) (DTONewspaper, error) {
	return DTONewspaper{
		DTOMaterial: DTOMaterial{
			UniqueCode:     n.UniqueCode,
			Name:           n.Name,
			DateOfEmission: n.DateOfEmission,
			NumberOfPages:  n.NumberOfPages,
			MaterialType:   n.MaterialType,
		},
		Url: n.Url,
	}, nil
}

func (s service) newspapersToDto(ns []Newspaper) ([]DTONewspaper, error) {
	var err error
	dtoMagazines := make([]DTONewspaper, len(ns))

	for idx, n := range ns {
		dtoMagazines[idx], err = s.newspaperToDto(n)

		if err != nil {
			return nil, errors.New("invalid newspaper object in newspapersToDto")
		}
	}

	return dtoMagazines, nil
}

func (s service) dtoToNewspaper(n DTONewspaper) (Newspaper, error) {
	return Newspaper{
		Material: Material{
			UniqueCode:     n.UniqueCode,
			Name:           n.Name,
			DateOfEmission: n.DateOfEmission,
			NumberOfPages:  n.NumberOfPages,
			MaterialType:   n.MaterialType,
		},
		Url: n.Url,
	}, nil
}

func (s service) dtoToNewspapers(ns []DTONewspaper) ([]Newspaper, error) {
	var err error
	newspapers := make([]Newspaper, len(ns))

	for idx, n := range ns {
		newspapers[idx], err = s.dtoToNewspaper(n)

		if err != nil {
			return nil, errors.New("invalid newspaper object in dtoToNewspapers")
		}
	}

	return newspapers, nil
}

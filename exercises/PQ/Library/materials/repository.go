package materials

import (
	"context"
	"errors"
	"excercise-library/ent"
	material2 "excercise-library/ent/material"
	"fmt"
	"log"
	//"time"
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
	client *ent.Client
}

func NewRepository(client *ent.Client) Repository {
	return &repository{
		client: client,
	}
}

// Materials
func (r repository) GetMaterials(ctx context.Context) ([]DTOMaterial, error) {
	repoMaterials, err := r.client.
		Material.
		Query().
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying materials: %v", err)
	}

	materials, err := r.repoMaterialsToDto(repoMaterials)
	if err != nil {
		return nil, err
	}

	log.Println("materials returned")

	return materials, nil
}

func (r repository) GetMaterialByCode(ctx context.Context, uniqueCode string) (DTOMaterial, error) {
	repoMaterial, err := r.client.
		Material.
		Query().
		Where(material2.UniqueCodeEQ(uniqueCode)).
		Only(ctx)
	if err != nil {
		return DTOMaterial{}, fmt.Errorf("failed querying material (%v): %v", uniqueCode, err)
	}

	material, err := r.repoMaterialToDto(repoMaterial)
	if err != nil {
		return DTOMaterial{}, err
	}

	log.Println("materials returned")

	return material, nil
}

func (r repository) DeleteMaterial(ctx context.Context, uniqueCode string) error {
	_, err := r.client.
		Material.
		Delete().
		Where(material2.UniqueCodeEQ(uniqueCode)).
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

// Books

func (r repository) GetBooks(ctx context.Context) ([]DTOBook, error) {
	repoMaterials, err := r.client.
		Material.
		Query().
		Where(material2.MaterialTypeEQ(int(BookType))).
		WithBook().
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying books: %v", err)
	}

	repoBooks := make([]*ent.Book, len(repoMaterials))

	for idx, m := range repoMaterials {
		repoBooks[idx] = m.Edges.Book

	}

	log.Println("books returned")

	var books []DTOBook
	books, err = r.repoBooksToDto(repoMaterials, repoBooks)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (r repository) GetBookByCode(ctx context.Context, uniqueCode string) (DTOBook, error) {
	repoMaterial, err := r.client.
		Material.
		Query().
		Where(material2.UniqueCodeEQ(uniqueCode)).
		Where(material2.MaterialTypeEQ(int(BookType))).
		Only(ctx)
	if err != nil {
		return DTOBook{}, fmt.Errorf("failed querying books: %v", err)
	}

	var repoBook *ent.Book

	repoBook, err = repoMaterial.QueryBook().Only(ctx)
	if err != nil {
		return DTOBook{}, errors.New("missing book for material in GetBookByCode")
	}

	log.Println("book returned")

	var book DTOBook
	book, err = r.repoBookToDto(repoMaterial, repoBook)
	if err != nil {
		return DTOBook{}, err
	}

	return book, nil
}

func (r repository) AddBook(ctx context.Context, book DTOBook) (DTOBook, error) {
	repoMaterial, repoBook, err := r.dtoToRepoBook(book)
	if err != nil {
		return DTOBook{}, err
	}

	_, err = r.client.
		Material.
		Create().
		SetUniqueCode(repoMaterial.UniqueCode).
		SetName(repoMaterial.Name).
		SetDateOfEmission(repoMaterial.DateOfEmission).
		SetNumberOfPages(repoMaterial.NumberOfPages).
		SetMaterialType(repoMaterial.MaterialType).
		SetBook(repoBook).
		Save(ctx)
	if err != nil {
		return DTOBook{}, fmt.Errorf("failed adding book: %v", err)
	}

	log.Println("book saved")

	return book, nil
}

func (r repository) UpdateBook(ctx context.Context, uniqueCode string, book DTOBook) (DTOBook, error) {
	repoMaterial, repoBook, err := r.dtoToRepoBook(book)
	if err != nil {
		return DTOBook{}, err
	}

	_, err = r.client.
		Material.
		Update().
		Where(material2.UniqueCodeEQ(uniqueCode)).
		SetName(repoMaterial.Name).
		SetDateOfEmission(repoMaterial.DateOfEmission).
		SetNumberOfPages(repoMaterial.NumberOfPages).
		SetMaterialType(repoMaterial.MaterialType).
		SetBook(repoBook).
		Save(ctx)
	if err != nil {
		return DTOBook{}, fmt.Errorf("failed updating book: %v", err)
	}

	log.Println("book updated")

	return book, nil
}

// Newspapers

func (r repository) GetNewspapers(ctx context.Context) ([]DTONewspaper, error) {
	repoMaterials, err := r.client.
		Material.
		Query().
		Where(material2.MaterialTypeEQ(int(NewspaperType))).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying newspapers: %v", err)
	}

	repoNewspapers := make([]*ent.Newspaper, len(repoMaterials))

	for idx, m := range repoMaterials {
		repoNewspapers[idx], err = m.QueryNewspaper().Only(ctx)
		if err != nil {
			return nil, errors.New("missing newspaper for material in GetNewspapers")
		}
	}

	log.Println("newspapers returned")

	var newspapers []DTONewspaper
	newspapers, err = r.repoNewspapersToDto(repoMaterials, repoNewspapers)
	if err != nil {
		return nil, err
	}

	return newspapers, nil
}

func (r repository) GetNewspaperByCode(ctx context.Context, uniqueCode string) (DTONewspaper, error) {
	repoMaterial, err := r.client.
		Material.
		Query().
		Where(material2.UniqueCodeEQ(uniqueCode)).
		Where(material2.MaterialTypeEQ(int(NewspaperType))).
		Only(ctx)
	if err != nil {
		return DTONewspaper{}, fmt.Errorf("failed querying newspapers: %v", err)
	}

	var repoNewspaper *ent.Newspaper

	repoNewspaper, err = repoMaterial.QueryNewspaper().Only(ctx)
	if err != nil {
		return DTONewspaper{}, errors.New("missing newspaper for material in GetNewspaperByCode")
	}

	log.Println("newspaper returned")

	var newspaper DTONewspaper
	newspaper, err = r.repoNewspaperToDto(repoMaterial, repoNewspaper)
	if err != nil {
		return DTONewspaper{}, err
	}

	return newspaper, nil
}

func (r repository) AddNewspaper(ctx context.Context, newspaper DTONewspaper) (DTONewspaper, error) {
	repoMaterial, repoNewspaper, err := r.dtoToRepoNewspaper(newspaper)
	if err != nil {
		return DTONewspaper{}, err
	}

	_, err = r.client.
		Material.
		Create().
		SetUniqueCode(repoMaterial.UniqueCode).
		SetName(repoMaterial.Name).
		SetDateOfEmission(repoMaterial.DateOfEmission).
		SetNumberOfPages(repoMaterial.NumberOfPages).
		SetMaterialType(repoMaterial.MaterialType).
		SetNewspaper(repoNewspaper).
		Save(ctx)
	if err != nil {
		return DTONewspaper{}, fmt.Errorf("failed adding book: %v", err)
	}

	log.Println("newspaper saved")

	return newspaper, nil
}

func (r repository) UpdateNewspaper(ctx context.Context, uniqueCode string, newspaper DTONewspaper) (DTONewspaper, error) {
	repoMaterial, repoNewspaper, err := r.dtoToRepoNewspaper(newspaper)
	if err != nil {
		return DTONewspaper{}, err
	}

	_, err = r.client.
		Material.
		Update().
		Where(material2.UniqueCodeEQ(uniqueCode)).
		SetName(repoMaterial.Name).
		SetDateOfEmission(repoMaterial.DateOfEmission).
		SetNumberOfPages(repoMaterial.NumberOfPages).
		SetMaterialType(repoMaterial.MaterialType).
		SetNewspaper(repoNewspaper).
		Save(ctx)
	if err != nil {
		return DTONewspaper{}, fmt.Errorf("failed updating newspaper: %v", err)
	}

	log.Println("newspaper updated")

	return newspaper, nil
}

// Magazines

func (r repository) GetMagazines(ctx context.Context) ([]DTOMagazine, error) {
	repoMaterials, err := r.client.
		Material.
		Query().
		Where(material2.MaterialTypeEQ(int(MagazineType))).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying magazines: %v", err)
	}

	repoMagazines := make([]*ent.Magazine, len(repoMaterials))

	for idx, m := range repoMaterials {
		repoMagazines[idx], err = m.QueryMagazine().Only(ctx)
		if err != nil {
			return nil, errors.New("missing book for material in GetBooks")
		}
	}

	log.Println("magazines returned")

	var magazines []DTOMagazine
	magazines, err = r.repoMagazinesToDto(repoMaterials, repoMagazines)
	if err != nil {
		return nil, err
	}

	return magazines, nil
}

func (r repository) GetMagazineByCode(ctx context.Context, uniqueCode string) (DTOMagazine, error) {
	repoMaterial, err := r.client.
		Material.
		Query().
		Where(material2.UniqueCodeEQ(uniqueCode)).
		Where(material2.MaterialTypeEQ(int(MagazineType))).
		Only(ctx)
	if err != nil {
		return DTOMagazine{}, fmt.Errorf("failed querying magazines: %v", err)
	}

	var repoMagazine *ent.Magazine

	repoMagazine, err = repoMaterial.QueryMagazine().Only(ctx)
	if err != nil {
		return DTOMagazine{}, errors.New("missing magazine for material in GetMagazineByCode")
	}

	log.Println("magazine returned")

	var magazine DTOMagazine
	magazine, err = r.repoMagazineToDto(repoMaterial, repoMagazine)
	if err != nil {
		return DTOMagazine{}, err
	}

	return magazine, nil
}

func (r repository) AddMagazine(ctx context.Context, magazine DTOMagazine) (DTOMagazine, error) {
	repoMaterial, repoMagazine, err := r.dtoToRepoMagazine(magazine)
	if err != nil {
		return DTOMagazine{}, err
	}

	_, err = r.client.
		Material.
		Create().
		SetUniqueCode(repoMaterial.UniqueCode).
		SetName(repoMaterial.Name).
		SetDateOfEmission(repoMaterial.DateOfEmission).
		SetNumberOfPages(repoMaterial.NumberOfPages).
		SetMaterialType(repoMaterial.MaterialType).
		SetMagazine(repoMagazine).
		Save(ctx)
	if err != nil {
		return DTOMagazine{}, fmt.Errorf("failed adding magazine: %v", err)
	}

	log.Println("magazine saved")

	return magazine, nil
}

func (r repository) UpdateMagazine(ctx context.Context, uniqueCode string, magazine DTOMagazine) (DTOMagazine, error) {
	repoMaterial, repoMagazine, err := r.dtoToRepoMagazine(magazine)
	if err != nil {
		return DTOMagazine{}, err
	}

	_, err = r.client.
		Material.
		Update().
		Where(material2.UniqueCodeEQ(uniqueCode)).
		SetName(repoMaterial.Name).
		SetDateOfEmission(repoMaterial.DateOfEmission).
		SetNumberOfPages(repoMaterial.NumberOfPages).
		SetMaterialType(repoMaterial.MaterialType).
		SetMagazine(repoMagazine).
		Save(ctx)
	if err != nil {
		return DTOMagazine{}, fmt.Errorf("failed updating magazine: %v", err)
	}

	log.Println("magazine updated")

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

func (r repository) materialTypeToInt(mt MaterialType) (int, error) {
	switch mt {
	case BookType:
		return 0, nil

	case MagazineType:
		return 1, nil

	case NewspaperType:
		return 2, nil

	default:
		return -1, errors.New(fmt.Sprintf("unsupported material type %d in materialTypeToInt", mt))
	}
}

func (r repository) intToMaterialType(mt int) (MaterialType, error) {
	switch mt {
	case 0:
		return BookType, nil

	case 1:
		return MagazineType, nil

	case 2:
		return NewspaperType, nil

	default:
		return BookType, errors.New(fmt.Sprintf("unsupported material type %d in intToMaterialType", mt))
	}
}

// Materials

func (r repository) repoMaterialToDto(m *ent.Material) (DTOMaterial, error) {
	return DTOMaterial{
		UniqueCode:     m.UniqueCode,
		Name:           m.Name,
		DateOfEmission: m.DateOfEmission,
		NumberOfPages:  m.NumberOfPages,
		MaterialType:   MaterialType(m.MaterialType),
	}, nil
}

func (r repository) repoMaterialsToDto(ms []*ent.Material) ([]DTOMaterial, error) {
	var err error
	dtoMaterials := make([]DTOMaterial, len(ms))

	for idx, m := range ms {
		dtoMaterials[idx], err = r.repoMaterialToDto(m)

		if err != nil {
			return nil, errors.New("invalid material object in repoMaterialsToDto")
		}
	}

	return dtoMaterials, nil
}

func (r repository) dtoToRepoMaterial(m DTOMaterial) (*ent.Material, error) {
	return &ent.Material{
		UniqueCode:     m.UniqueCode,
		Name:           m.Name,
		DateOfEmission: m.DateOfEmission,
		NumberOfPages:  m.NumberOfPages,
		MaterialType:   int(m.MaterialType),
	}, nil
}

func (r repository) dtoToRepoMaterials(ms []DTOMaterial) ([]*ent.Material, error) {
	var err error
	materials := make([]*ent.Material, len(ms))

	for idx, m := range ms {
		materials[idx], err = r.dtoToRepoMaterial(m)

		if err != nil {
			return nil, errors.New("invalid material object in dtoToRepoMaterials")
		}

	}

	return materials, nil
}

// Books

func (r repository) repoBookToDto(m *ent.Material, b *ent.Book) (DTOBook, error) {
	mt, err := r.intToMaterialType(m.MaterialType)
	if err != nil {
		return DTOBook{}, errors.New("invalid material type in repoBookToDto")
	}

	return DTOBook{
		DTOMaterial: DTOMaterial{
			UniqueCode:     m.UniqueCode,
			Name:           m.Name,
			DateOfEmission: m.DateOfEmission,
			NumberOfPages:  m.NumberOfPages,
			MaterialType:   mt,
		},
		AuthorName: b.AuthorName,
		Genre:      b.Genre,
	}, nil
}

func (r repository) repoBooksToDto(ms []*ent.Material, bs []*ent.Book) ([]DTOBook, error) {
	var err error
	dtoBooks := make([]DTOBook, len(ms))

	for idx, m := range ms {
		dtoBooks[idx], err = r.repoBookToDto(m, bs[idx])

		if err != nil {
			return nil, errors.New("invalid book object in repoBooksToDto")
		}
	}

	return dtoBooks, nil
}

func (r repository) dtoToRepoBook(b DTOBook) (*ent.Material, *ent.Book, error) {
	return &ent.Material{
			UniqueCode:     b.UniqueCode,
			Name:           b.Name,
			DateOfEmission: b.DateOfEmission,
			NumberOfPages:  b.NumberOfPages,
			MaterialType:   int(b.MaterialType),
		},
		&ent.Book{
			AuthorName: b.AuthorName,
			Genre:      b.Genre,
		},
		nil
}

func (r repository) dtoToRepoBooks(bs []DTOBook) ([]*ent.Material, []*ent.Book, error) {
	var err error
	materials := make([]*ent.Material, len(bs))
	books := make([]*ent.Book, len(bs))

	for idx, b := range bs {
		materials[idx], books[idx], err = r.dtoToRepoBook(b)
		if err != nil {
			return nil, nil, errors.New("invalid material object in dtoToRepoBooks")
		}

	}

	return materials, books, nil
}

// Newspaper

func (r repository) repoNewspaperToDto(m *ent.Material, n *ent.Newspaper) (DTONewspaper, error) {
	mt, err := r.intToMaterialType(m.MaterialType)
	if err != nil {
		return DTONewspaper{}, errors.New("invalid material type in repoNewspaperToDto")
	}

	return DTONewspaper{
		DTOMaterial: DTOMaterial{
			UniqueCode:     m.UniqueCode,
			Name:           m.Name,
			DateOfEmission: m.DateOfEmission,
			NumberOfPages:  m.NumberOfPages,
			MaterialType:   mt,
		},
		Url: n.URL,
	}, nil
}

func (r repository) repoNewspapersToDto(ns []*ent.Material, b []*ent.Newspaper) ([]DTONewspaper, error) {
	var err error
	dtoNewspapers := make([]DTONewspaper, len(ns))

	for idx, n := range ns {
		dtoNewspapers[idx], err = r.repoNewspaperToDto(n, b[idx])

		if err != nil {
			return nil, errors.New("invalid newspaper object in repoNewspapersToDto")
		}
	}

	return dtoNewspapers, nil
}

func (r repository) dtoToRepoNewspaper(n DTONewspaper) (*ent.Material, *ent.Newspaper, error) {
	return &ent.Material{
			UniqueCode:     n.UniqueCode,
			Name:           n.Name,
			DateOfEmission: n.DateOfEmission,
			NumberOfPages:  n.NumberOfPages,
			MaterialType:   int(n.MaterialType),
		},
		&ent.Newspaper{
			URL: n.Url,
		},
		nil
}

func (r repository) dtoToRepoNewspapers(ns []DTONewspaper) ([]*ent.Material, []*ent.Newspaper, error) {
	var err error
	materials := make([]*ent.Material, len(ns))
	newspapers := make([]*ent.Newspaper, len(ns))

	for idx, n := range ns {
		materials[idx], newspapers[idx], err = r.dtoToRepoNewspaper(n)
		if err != nil {
			return nil, nil, errors.New("invalid material object in dtoToRepoNewspapers")
		}

	}

	return materials, newspapers, nil
}

// Magazine

func (r repository) repoMagazineToDto(m *ent.Material, mg *ent.Magazine) (DTOMagazine, error) {
	mt, err := r.intToMaterialType(m.MaterialType)
	if err != nil {
		return DTOMagazine{}, errors.New("invalid material type in repoMagazineToDto")
	}

	return DTOMagazine{
		DTOMaterial: DTOMaterial{
			UniqueCode:     m.UniqueCode,
			Name:           m.Name,
			DateOfEmission: m.DateOfEmission,
			NumberOfPages:  m.NumberOfPages,
			MaterialType:   mt,
		},
		Url: mg.URL,
	}, nil
}

func (r repository) repoMagazinesToDto(ms []*ent.Material, mgs []*ent.Magazine) ([]DTOMagazine, error) {
	var err error
	dtoMagazines := make([]DTOMagazine, len(ms))

	for idx, m := range ms {
		dtoMagazines[idx], err = r.repoMagazineToDto(m, mgs[idx])

		if err != nil {
			return nil, errors.New("invalid magazine object in repoMagazinesToDto")
		}
	}

	return dtoMagazines, nil
}

func (r repository) dtoToRepoMagazine(mg DTOMagazine) (*ent.Material, *ent.Magazine, error) {
	return &ent.Material{
			UniqueCode:     mg.UniqueCode,
			Name:           mg.Name,
			DateOfEmission: mg.DateOfEmission,
			NumberOfPages:  mg.NumberOfPages,
			MaterialType:   int(mg.MaterialType),
		},
		&ent.Magazine{
			URL: mg.Url,
		},
		nil
}

func (r repository) dtoToRepoMagazines(mgs []DTOMagazine) ([]*ent.Material, []*ent.Magazine, error) {
	var err error
	materials := make([]*ent.Material, len(mgs))
	magazines := make([]*ent.Magazine, len(mgs))

	for idx, mg := range mgs {
		materials[idx], magazines[idx], err = r.dtoToRepoMagazine(mg)
		if err != nil {
			return nil, nil, errors.New("invalid material object in dtoToRepoMagazines")
		}

	}

	return materials, magazines, nil
}

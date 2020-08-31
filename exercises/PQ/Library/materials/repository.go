package materials

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"go-kit-template/ent"
	material2 "go-kit-template/ent/material"
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
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

// Materials
func (r repository) GetMaterials(ctx context.Context) ([]DTOMaterial, error) {
	client, errOpen := ent.Open("mysql", "root:RooT27134668@tcp(localhost:3306)/dev_library")
	if errOpen != nil {
		return nil, errOpen
	}
	defer client.Close()

	repoMaterials, err := client.
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
	client, errOpen := ent.Open("mysql", "root:RooT27134668@tcp(localhost:3306)/dev_library")
	if errOpen != nil {
		return DTOMaterial{}, errOpen
	}
	defer client.Close()

	repoMaterial, err := client.
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
	client, errOpen := ent.Open("mysql", "root:RooT27134668@tcp(localhost:3306)/dev_library")
	if errOpen != nil {
		return errOpen
	}
	defer client.Close()

	_, err := client.
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
	client, errOpen := ent.Open("mysql", "root:RooT27134668@tcp(localhost:3306)/dev_library")
	if errOpen != nil {
		return nil, errOpen
	}
	defer client.Close()

	repoMaterials, err := client.
		Material.
		Query().
		Where(material2.MaterialTypeEQ(int(BookType))).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying books: %v", err)
	}

	repoBooks := make([]*ent.Book, len(repoMaterials))

	for idx, m := range repoMaterials {
		repoBooks[idx], err = m.QueryBook().Only(ctx)
		if err != nil {
			return nil, errors.New("missing book for material in GetBooks")
		}
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
	client, errOpen := ent.Open("mysql", "root:RooT27134668@tcp(localhost:3306)/dev_library")
	if errOpen != nil {
		return DTOBook{}, errOpen
	}
	defer client.Close()

	repoMaterial, err := client.
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
	client, errOpen := ent.Open("mysql", "root:RooT27134668@tcp(localhost:3306)/dev_library")
	if errOpen != nil {
		return DTOBook{}, errOpen
	}
	defer client.Close()

	repoMaterial, repoBook, err := r.dtoToRepoBook(book)
	if err != nil {
		return DTOBook{}, err
	}

	_, err = client.
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

	log.Println("materials saved")

	return book, nil
}

func (r repository) UpdateBook(ctx context.Context, uniqueCode string, book DTOBook) (DTOBook, error) {
	client, errOpen := ent.Open("mysql", "root:RooT27134668@tcp(localhost:3306)/dev_library")
	if errOpen != nil {
		return DTOBook{}, errOpen
	}
	defer client.Close()

	repoMaterial, repoBook, err := r.dtoToRepoBook(book)
	if err != nil {
		return DTOBook{}, err
	}

	_, err = client.
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

	log.Println("materials updated")

	return book, nil
}

// Newspapers

func (r repository) GetNewspapers(_ context.Context) ([]DTONewspaper, error) {
	return nil, nil
}

func (r repository) GetNewspaperByCode(_ context.Context, uniqueCode string) (DTONewspaper, error) {
	return DTONewspaper{}, nil
}

func (r repository) AddNewspaper(_ context.Context, newspaper DTONewspaper) (DTONewspaper, error) {
	return DTONewspaper{}, nil
}

func (r repository) UpdateNewspaper(_ context.Context, uniqueCode string, newspaper DTONewspaper) (DTONewspaper, error) {
	return DTONewspaper{}, nil
}

// Magazines

func (r repository) GetMagazines(_ context.Context) ([]DTOMagazine, error) {
	return nil, nil
}

func (r repository) GetMagazineByCode(_ context.Context, uniqueCode string) (DTOMagazine, error) {
	return DTOMagazine{}, nil
}

func (r repository) AddMagazine(_ context.Context, magazine DTOMagazine) (DTOMagazine, error) {
	return DTOMagazine{}, nil
}

func (r repository) UpdateMagazine(_ context.Context, uniqueCode string, magazine DTOMagazine) (DTOMagazine, error) {
	return DTOMagazine{}, nil
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

func (r repository) repoBooksToDto(ms []*ent.Material, b []*ent.Book) ([]DTOBook, error) {
	var err error
	dtoBooks := make([]DTOBook, len(ms))

	for idx, m := range ms {
		dtoBooks[idx], err = r.repoBookToDto(m, b[idx])

		if err != nil {
			return nil, errors.New("invalid book object in repoBooksToDto")
		}
	}

	return dtoBooks, nil
}

func (r repository) dtoToRepoBook(m DTOBook) (*ent.Material, *ent.Book, error) {
	return &ent.Material{
			UniqueCode:     m.UniqueCode,
			Name:           m.Name,
			DateOfEmission: m.DateOfEmission,
			NumberOfPages:  m.NumberOfPages,
			MaterialType:   int(m.MaterialType),
		},
		&ent.Book{
			AuthorName: m.AuthorName,
			Genre:      m.Genre,
		},
		nil
}

func (r repository) dtoToRepoBooks(ms []DTOBook) ([]*ent.Material, []*ent.Book, error) {
	var err error
	materials := make([]*ent.Material, len(ms))
	books := make([]*ent.Book, len(ms))

	for idx, m := range ms {
		materials[idx], books[idx], err = r.dtoToRepoBook(m)
		if err != nil {
			return nil, nil, errors.New("invalid material object in dtoToRepoBooks")
		}

	}

	return materials, books, nil
}

package materials

import (
	"context"
	"errors"
	"excercise-library/ent"
	book2 "excercise-library/ent/book"
	magazine2 "excercise-library/ent/magazine"
	material2 "excercise-library/ent/material"
	newspaper2 "excercise-library/ent/newspaper"
	section2 "excercise-library/ent/section"
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
	repoBooks, err := r.client.
		Material.
		Query().
		Where(material2.MaterialTypeEQ(int(BookType))).
		WithBook().
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying books: %v", err)
	}

	log.Println("books returned")

	var books []DTOBook
	books, err = r.repoBooksToDto(repoBooks)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (r repository) GetBookByCode(ctx context.Context, uniqueCode string) (DTOBook, error) {
	repoBook, err := r.client.
		Material.
		Query().
		Where(material2.UniqueCodeEQ(uniqueCode)).
		Where(material2.MaterialTypeEQ(int(BookType))).
		WithBook().
		Only(ctx)
	if err != nil {
		return DTOBook{}, fmt.Errorf("failed querying book by code: %v", err)
	}

	log.Println("book returned")

	var book DTOBook
	book, err = r.repoBookToDto(repoBook)
	if err != nil {
		return DTOBook{}, err
	}

	return book, nil
}

func (r repository) AddBook(ctx context.Context, book DTOBook) (DTOBook, error) {
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return DTOBook{}, fmt.Errorf("starting a transaction: %v", err)
	}

	var repoMaterial *ent.Material
	repoMaterial, err = tx.
		Material.
		Create().
		SetUniqueCode(book.UniqueCode).
		SetName(book.Name).
		SetDateOfEmission(book.DateOfEmission).
		SetNumberOfPages(book.NumberOfPages).
		SetMaterialType(int(book.MaterialType)).
		Save(ctx)
	if err != nil {
		return DTOBook{}, rollback(tx, fmt.Errorf("failed adding book (material): %v", err))
	}

	_, err = tx.
		Book.
		Create().
		SetAuthorName(book.AuthorName).
		SetGenre(book.Genre).
		SetRelatedMaterialID(repoMaterial.ID).
		Save(ctx)
	if err != nil {
		return DTOBook{}, rollback(tx, fmt.Errorf("failed adding book (book): %v", err))
	}

	err = tx.Commit()
	if err != nil {
		return DTOBook{}, err
	}

	log.Println("book saved")

	return book, nil
}

func (r repository) UpdateBook(ctx context.Context, uniqueCode string, book DTOBook) (DTOBook, error) {
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return DTOBook{}, fmt.Errorf("starting a transaction: %v", err)
	}

	//var updatedCount int
	_, err = tx.
		Material.
		Update().
		Where(material2.UniqueCodeEQ(uniqueCode)).
		Where(material2.MaterialTypeEQ(int(BookType))).
		SetName(book.Name).
		SetDateOfEmission(book.DateOfEmission).
		SetNumberOfPages(book.NumberOfPages).
		SetMaterialType(int(book.MaterialType)).
		Save(ctx)
	if err != nil {
		return DTOBook{}, rollback(tx, fmt.Errorf("failed updating book (material): %v", err))
	}

	_, err = tx.
		Book.
		Update().
		Where(book2.HasRelatedMaterialWith(material2.UniqueCodeEQ(uniqueCode))).
		SetAuthorName(book.AuthorName).
		SetGenre(book.Genre).
		Save(ctx)
	if err != nil {
		return DTOBook{}, rollback(tx, fmt.Errorf("failed updating book (book): %v", err))
	}

	err = tx.Commit()
	if err != nil {
		return DTOBook{}, err
	}

	log.Println("book updated")

	return book, nil
}

// Newspapers

func (r repository) GetNewspapers(ctx context.Context) ([]DTONewspaper, error) {
	repoNewspapers, err := r.client.
		Material.
		Query().
		Where(material2.MaterialTypeEQ(int(NewspaperType))).
		WithNewspaper().
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying newspapers: %v", err)
	}

	log.Println("newspapers returned")

	var newspapers []DTONewspaper
	newspapers, err = r.repoNewspapersToDto(repoNewspapers)
	if err != nil {
		return nil, err
	}

	return newspapers, nil
}

func (r repository) GetNewspaperByCode(ctx context.Context, uniqueCode string) (DTONewspaper, error) {
	repoNewspaper, err := r.client.
		Material.
		Query().
		Where(material2.UniqueCodeEQ(uniqueCode)).
		Where(material2.MaterialTypeEQ(int(NewspaperType))).
		WithNewspaper().
		Only(ctx)
	if err != nil {
		return DTONewspaper{}, fmt.Errorf("failed querying newspaper by code: %v", err)
	}

	log.Println("newspaper returned")

	var newspaper DTONewspaper
	newspaper, err = r.repoNewspaperToDto(repoNewspaper)
	if err != nil {
		return DTONewspaper{}, err
	}

	return newspaper, nil
}

func (r repository) AddNewspaper(ctx context.Context, newspaper DTONewspaper) (DTONewspaper, error) {
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return DTONewspaper{}, fmt.Errorf("starting a transaction: %v", err)
	}

	var repoMaterial *ent.Material
	repoMaterial, err = tx.
		Material.
		Create().
		SetUniqueCode(newspaper.UniqueCode).
		SetName(newspaper.Name).
		SetDateOfEmission(newspaper.DateOfEmission).
		SetNumberOfPages(newspaper.NumberOfPages).
		SetMaterialType(int(newspaper.MaterialType)).
		Save(ctx)
	if err != nil {
		return DTONewspaper{}, rollback(tx, fmt.Errorf("failed adding newspaper (material): %v", err))
	}

	_, err = tx.
		Newspaper.
		Create().
		SetURL(newspaper.Url).
		SetRelatedMaterialID(repoMaterial.ID).
		Save(ctx)
	if err != nil {
		return DTONewspaper{}, rollback(tx, fmt.Errorf("failed adding newspaper (newspaper): %v", err))
	}

	err = tx.Commit()
	if err != nil {
		return DTONewspaper{}, err
	}

	log.Println("newspaper saved")

	return newspaper, nil
}

func (r repository) UpdateNewspaper(ctx context.Context, uniqueCode string, newspaper DTONewspaper) (DTONewspaper, error) {
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return DTONewspaper{}, fmt.Errorf("starting a transaction: %v", err)
	}

	//var updatedCount int
	_, err = tx.
		Material.
		Update().
		Where(material2.UniqueCodeEQ(uniqueCode)).
		Where(material2.MaterialTypeEQ(int(NewspaperType))).
		SetName(newspaper.Name).
		SetDateOfEmission(newspaper.DateOfEmission).
		SetNumberOfPages(newspaper.NumberOfPages).
		SetMaterialType(int(newspaper.MaterialType)).
		Save(ctx)
	if err != nil {
		return DTONewspaper{}, rollback(tx, fmt.Errorf("failed updating newspaper (material): %v", err))
	}

	_, err = tx.
		Newspaper.
		Update().
		Where(newspaper2.HasRelatedMaterialWith(material2.UniqueCodeEQ(uniqueCode))).
		SetURL(newspaper.Url).
		Save(ctx)
	if err != nil {
		return DTONewspaper{}, rollback(tx, fmt.Errorf("failed updating newspaper (newspaper): %v", err))
	}

	err = tx.Commit()
	if err != nil {
		return DTONewspaper{}, err
	}

	log.Println("newspaper updated")

	return newspaper, nil
}

// Magazines

func (r repository) GetMagazines(ctx context.Context) ([]DTOMagazine, error) {
	repoMagazines, err := r.client.
		Material.
		Query().
		Where(material2.MaterialTypeEQ(int(MagazineType))).
		WithMagazine(func(query *ent.MagazineQuery) { query.WithSection() }).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying magazines: %v", err)
	}

	log.Println("magazines returned")

	var magazines []DTOMagazine
	magazines, err = r.repoMagazinesToDto(repoMagazines)
	if err != nil {
		return nil, err
	}

	return magazines, nil
}

func (r repository) GetMagazineByCode(ctx context.Context, uniqueCode string) (DTOMagazine, error) {
	repoMagazine, err := r.getRepoMagazineByCode(ctx, uniqueCode)
	if err != nil {
		return DTOMagazine{}, fmt.Errorf("failed querying magazine by code: %v", err)
	}

	log.Println("magazine returned")

	var magazine DTOMagazine
	magazine, err = r.repoMagazineToDto(repoMagazine)
	if err != nil {
		return DTOMagazine{}, err
	}

	return magazine, nil
}

func (r repository) getRepoMagazineByCode(ctx context.Context, uniqueCode string) (*ent.Material, error) {
	repoMagazine, err := r.client.
		Material.
		Query().
		Where(material2.UniqueCodeEQ(uniqueCode)).
		Where(material2.MaterialTypeEQ(int(MagazineType))).
		WithMagazine(func(query *ent.MagazineQuery) { query.WithSection() }).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying magazine by code: %v", err)
	}

	return repoMagazine, nil
}

func (r repository) AddMagazine(ctx context.Context, magazine DTOMagazine) (DTOMagazine, error) {
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return DTOMagazine{}, fmt.Errorf("starting a transaction: %v", err)
	}

	var repoMaterial *ent.Material
	repoMaterial, err = tx.
		Material.
		Create().
		SetUniqueCode(magazine.UniqueCode).
		SetName(magazine.Name).
		SetDateOfEmission(magazine.DateOfEmission).
		SetNumberOfPages(magazine.NumberOfPages).
		SetMaterialType(int(magazine.MaterialType)).
		Save(ctx)
	if err != nil {
		return DTOMagazine{}, rollback(tx, fmt.Errorf("failed adding magazine (material): %v", err))
	}

	var repoMagazine *ent.Magazine
	repoMagazine, err = tx.
		Magazine.
		Create().
		SetURL(magazine.Url).
		SetRelatedMaterialID(repoMaterial.ID).
		Save(ctx)
	if err != nil {
		return DTOMagazine{}, rollback(tx, fmt.Errorf("failed adding magazine (magazine): %v", err))
	}

	for _, s := range magazine.Sections {
		_, err = tx.
			Section.
			Create().
			SetCode(s.Code).
			SetContent(s.Content).
			SetRelatedMagazineID(repoMagazine.ID).
			Save(ctx)
		if err != nil {
			return DTOMagazine{}, rollback(tx, fmt.Errorf("failed adding magazine (section): %v", err))
		}
	}

	err = tx.Commit()
	if err != nil {
		return DTOMagazine{}, err
	}

	log.Println("magazine saved")

	return magazine, nil
}

func (r repository) UpdateMagazine(ctx context.Context, uniqueCode string, magazine DTOMagazine) (DTOMagazine, error) {
	repoMagazine, err := r.getRepoMagazineByCode(ctx, uniqueCode)
	if err != nil {
		return DTOMagazine{}, fmt.Errorf("failed querying magazine by code: %v", err)
	}

	var tx *ent.Tx
	tx, err = r.client.Tx(ctx)
	if err != nil {
		return DTOMagazine{}, fmt.Errorf("starting a transaction: %v", err)
	}

	//var updatedCount int
	_, err = tx.
		Material.
		Update().
		Where(material2.UniqueCodeEQ(uniqueCode)).
		Where(material2.MaterialTypeEQ(int(NewspaperType))).
		SetName(magazine.Name).
		SetDateOfEmission(magazine.DateOfEmission).
		SetNumberOfPages(magazine.NumberOfPages).
		SetMaterialType(int(magazine.MaterialType)).
		Save(ctx)
	if err != nil {
		return DTOMagazine{}, rollback(tx, fmt.Errorf("failed updating magazine (material): %v", err))
	}

	_, err = tx.
		Magazine.
		Update().
		Where(magazine2.HasRelatedMaterialWith(material2.UniqueCodeEQ(uniqueCode))).
		SetURL(magazine.Url).
		Save(ctx)
	if err != nil {
		return DTOMagazine{}, rollback(tx, fmt.Errorf("failed updating magazine (magazine): %v", err))
	}

	_, err = tx.
		Section.
		Delete().
		Where(section2.HasRelatedMagazineWith(magazine2.HasRelatedMaterialWith(material2.UniqueCodeEQ(uniqueCode)))).
		Exec(ctx)
	if err != nil {
		return DTOMagazine{}, rollback(tx, fmt.Errorf("failed updating magazine (deleteing previous sections): %v", err))
	}

	for _, s := range magazine.Sections {
		_, err = tx.
			Section.
			Create().
			SetCode(s.Code).
			SetContent(s.Content).
			SetRelatedMagazineID(repoMagazine.Edges.Magazine.ID).
			Save(ctx)
		if err != nil {
			return DTOMagazine{}, rollback(tx, fmt.Errorf("failed adding magazine (section): %v", err))
		}
	}

	err = tx.Commit()
	if err != nil {
		return DTOMagazine{}, err
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

// Books

func (r repository) repoBookToDto(b *ent.Material) (DTOBook, error) {
	mt, err := r.intToMaterialType(b.MaterialType)
	if err != nil {
		return DTOBook{}, errors.New("invalid material type in repoBookToDto")
	}

	book := b.Edges.Book

	return DTOBook{
		DTOMaterial: DTOMaterial{
			UniqueCode:     b.UniqueCode,
			Name:           b.Name,
			DateOfEmission: b.DateOfEmission,
			NumberOfPages:  b.NumberOfPages,
			MaterialType:   mt,
		},
		AuthorName: book.AuthorName,
		Genre:      book.Genre,
	}, nil
}

func (r repository) repoBooksToDto(bs []*ent.Material) ([]DTOBook, error) {
	var err error
	dtoBooks := make([]DTOBook, len(bs))

	for idx, b := range bs {
		dtoBooks[idx], err = r.repoBookToDto(b)

		if err != nil {
			return nil, errors.New("invalid book object in repoBooksToDto")
		}
	}

	return dtoBooks, nil
}

// Newspaper

func (r repository) repoNewspaperToDto(n *ent.Material) (DTONewspaper, error) {
	mt, err := r.intToMaterialType(n.MaterialType)
	if err != nil {
		return DTONewspaper{}, errors.New("invalid material type in repoNewspaperToDto")
	}

	newspaper := n.Edges.Newspaper

	return DTONewspaper{
		DTOMaterial: DTOMaterial{
			UniqueCode:     n.UniqueCode,
			Name:           n.Name,
			DateOfEmission: n.DateOfEmission,
			NumberOfPages:  n.NumberOfPages,
			MaterialType:   mt,
		},
		Url: newspaper.URL,
	}, nil
}

func (r repository) repoNewspapersToDto(ns []*ent.Material) ([]DTONewspaper, error) {
	var err error
	dtoNewspapers := make([]DTONewspaper, len(ns))

	for idx, n := range ns {
		dtoNewspapers[idx], err = r.repoNewspaperToDto(n)

		if err != nil {
			return nil, errors.New("invalid newspaper object in repoNewspapersToDto")
		}
	}

	return dtoNewspapers, nil
}

// Magazine

func (r repository) repoMagazineToDto(m *ent.Material) (DTOMagazine, error) {
	mt, err := r.intToMaterialType(m.MaterialType)
	if err != nil {
		return DTOMagazine{}, errors.New("invalid material type in repoMagazineToDto")
	}

	magazine := m.Edges.Magazine
	sections := magazine.Edges.Section

	var dtoSections []DTOSection
	dtoSections, err = r.repoSectionsToDto(sections)
	if err != nil {
		return DTOMagazine{}, err
	}

	return DTOMagazine{
		DTOMaterial: DTOMaterial{
			UniqueCode:     m.UniqueCode,
			Name:           m.Name,
			DateOfEmission: m.DateOfEmission,
			NumberOfPages:  m.NumberOfPages,
			MaterialType:   mt,
		},
		Url:      magazine.URL,
		Sections: dtoSections,
	}, nil
}

func (r repository) repoMagazinesToDto(ms []*ent.Material) ([]DTOMagazine, error) {
	var err error
	dtoMagazines := make([]DTOMagazine, len(ms))

	for idx, m := range ms {
		dtoMagazines[idx], err = r.repoMagazineToDto(m)

		if err != nil {
			return nil, errors.New("invalid magazine object in repoMagazinesToDto")
		}
	}

	return dtoMagazines, nil
}

func (r repository) repoSectionToDto(s *ent.Section) (DTOSection, error) {
	return DTOSection{
		Code:    s.Code,
		Content: s.Content,
	}, nil
}

func (r repository) repoSectionsToDto(ss []*ent.Section) ([]DTOSection, error) {
	var err error
	dtoSections := make([]DTOSection, len(ss))

	for idx, s := range ss {
		dtoSections[idx], err = r.repoSectionToDto(s)

		if err != nil {
			return nil, errors.New("invalid section object in repoSectionsToDto")
		}
	}

	return dtoSections, nil
}

// rollback calls to tx.Rollback and wraps the given error with the rollback error if occurred.
func rollback(tx *ent.Tx, err error) error {
	if rollErr := tx.Rollback(); rollErr != nil {
		err = fmt.Errorf("%v: %v", err, rollErr)
	}
	return err
}

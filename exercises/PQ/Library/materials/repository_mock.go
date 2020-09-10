package materials

import (
	"context"
	"excercise-library/ent"
	"excercise-library/shared"
	"time"
)

var (
	MockMaterials = []DTOMaterial{
		{
			UniqueCode:     "Book_123",
			Name:           "Book Name 123",
			DateOfEmission: time.Time{},
			NumberOfPages:  12,
			MaterialType:   BookType,
		},
		{
			UniqueCode:     "Book_321",
			Name:           "Book Name 321",
			DateOfEmission: time.Time{},
			NumberOfPages:  34,
			MaterialType:   BookType,
		},
		{
			UniqueCode:     "Magazine_534",
			Name:           "Magazine Name 534",
			DateOfEmission: time.Time{},
			NumberOfPages:  36,
			MaterialType:   MagazineType,
		},
		{
			UniqueCode:     "Newspaper_432",
			Name:           "Newspaper Name 432",
			DateOfEmission: time.Time{},
			NumberOfPages:  72,
			MaterialType:   NewspaperType,
		},
		{
			UniqueCode:     "Magazine_153",
			Name:           "Magazine Name 153",
			DateOfEmission: time.Time{},
			NumberOfPages:  867,
			MaterialType:   MagazineType,
		},
	}
	MockBooks = []DTOBook{
		{
			DTOMaterial: DTOMaterial{
				UniqueCode:     "Book_123",
				Name:           "Book Name 123",
				DateOfEmission: time.Time{},
				NumberOfPages:  12,
				MaterialType:   BookType,
			},
			AuthorName: "AuthorName 123",
			Genre:      "Genre 123",
		},
		{
			DTOMaterial: DTOMaterial{
				UniqueCode:     "Book_321",
				Name:           "Book Name 321",
				DateOfEmission: time.Time{},
				NumberOfPages:  34,
				MaterialType:   BookType,
			},
			AuthorName: "AuthorName 321",
			Genre:      "Genre 321",
		},
	}
	MockMagazines = []DTOMagazine{
		{
			DTOMaterial: DTOMaterial{
				UniqueCode:     "Magazine_534",
				Name:           "Magazine Name 534",
				DateOfEmission: time.Time{},
				NumberOfPages:  36,
				MaterialType:   MagazineType,
			},
			Sections: []DTOSection{
				{
					Code:    "Code_534_1",
					Content: "Content_534_1",
				},
				{
					Code:    "Code_534_2",
					Content: "Content_534_2",
				},
				{
					Code:    "Code_534_3",
					Content: "Content_534_3",
				},
			},
			Url: "http://Magazine.Url534.com/",
		},
		{
			DTOMaterial: DTOMaterial{
				UniqueCode:     "Magazine_153",
				Name:           "Magazine Name 153",
				DateOfEmission: time.Time{},
				NumberOfPages:  867,
				MaterialType:   MagazineType,
			},
			Sections: []DTOSection{
				{
					Code:    "Code_153_1",
					Content: "Content_153_1",
				},
				{
					Code:    "Code_153_2",
					Content: "Content_153_2",
				},
			},
			Url: "http://Magazine.Url153.com/",
		},
	}
	MockNewspapers = []DTONewspaper{
		{
			DTOMaterial: DTOMaterial{UniqueCode: "Newspaper_432",
				Name:           "Newspaper Name 432",
				DateOfEmission: time.Time{},
				NumberOfPages:  72,
				MaterialType:   NewspaperType},
			Url: "http://Newspaper.Url432.com/",
		},
	}
)

type repositoryMock struct {
}

func NewRepositoryMock() Repository {
	return repositoryMock{}
}

func (r repositoryMock) GetClient(ctx context.Context) *ent.Client {
	if ctx.Value("GetClientError") == true {
		return nil
	}
	return nil
}

func (r repositoryMock) StartTx(ctx context.Context) (context.Context, error) {
	if ctx.Value("StartTxError") == true {
		return nil, shared.ErrDefault
	}
	return ctx, nil
}

func (r repositoryMock) Commit(ctx context.Context) error {
	if ctx.Value("CommitError") == true {
		return shared.ErrDefault
	}
	return nil
}

func (r repositoryMock) Rollback(ctx context.Context) error {
	if ctx.Value("RollbackError") == true {
		return shared.ErrDefault
	}
	return nil
}

func (r repositoryMock) GetMaterials(ctx context.Context) ([]DTOMaterial, error) {
	if ctx.Value("GetMaterialsError") == true {
		return nil, shared.ErrDefault
	}
	return MockMaterials, nil
}

func (r repositoryMock) GetMaterialByCode(ctx context.Context, _ string) (DTOMaterial, error) {
	if ctx.Value("GetMaterialByCodeError") == true {
		return DTOMaterial{}, shared.ErrDefault
	}
	return MockMaterials[0], nil
}

func (r repositoryMock) DeleteMaterial(ctx context.Context, _ string) error {
	if ctx.Value("DeleteMaterialError") == true {
		return shared.ErrDefault
	}
	return nil
}

func (r repositoryMock) GetBooks(ctx context.Context) ([]DTOBook, error) {
	if ctx.Value("GetBooksError") == true {
		return nil, shared.ErrDefault
	}
	return MockBooks, nil
}

func (r repositoryMock) GetBookByCode(ctx context.Context, _ string) (DTOBook, error) {
	if ctx.Value("GetBookByCodeError") == true {
		return DTOBook{}, shared.ErrDefault
	}
	return MockBooks[0], nil
}

func (r repositoryMock) AddBook(ctx context.Context, _ DTOBook) (DTOBook, error) {
	if ctx.Value("AddBookError") == true {
		return DTOBook{}, shared.ErrDefault
	}
	return MockBooks[0], nil
}

func (r repositoryMock) UpdateBook(ctx context.Context, _ string, _ DTOBook) (DTOBook, error) {
	if ctx.Value("UpdateBookError") == true {
		return DTOBook{}, shared.ErrDefault
	}
	return MockBooks[0], nil
}

func (r repositoryMock) GetNewspapers(ctx context.Context) ([]DTONewspaper, error) {
	if ctx.Value("GetNewspapersError") == true {
		return nil, shared.ErrDefault
	}
	return MockNewspapers, nil
}

func (r repositoryMock) GetNewspaperByCode(ctx context.Context, _ string) (DTONewspaper, error) {
	if ctx.Value("GetNewspaperByCodeError") == true {
		return DTONewspaper{}, shared.ErrDefault
	}
	return MockNewspapers[0], nil
}

func (r repositoryMock) AddNewspaper(ctx context.Context, _ DTONewspaper) (DTONewspaper, error) {
	if ctx.Value("AddNewspaperError") == true {
		return DTONewspaper{}, shared.ErrDefault
	}
	return MockNewspapers[0], nil
}

func (r repositoryMock) UpdateNewspaper(ctx context.Context, _ string, _ DTONewspaper) (DTONewspaper, error) {
	if ctx.Value("UpdateNewspaperError") == true {
		return DTONewspaper{}, shared.ErrDefault
	}
	return MockNewspapers[0], nil
}

func (r repositoryMock) GetMagazines(ctx context.Context) ([]DTOMagazine, error) {
	if ctx.Value("GetMagazinesError") == true {
		return nil, shared.ErrDefault
	}
	return MockMagazines, nil
}

func (r repositoryMock) GetMagazineByCode(ctx context.Context, _ string) (DTOMagazine, error) {
	if ctx.Value("GetMagazineByCodeError") == true {
		return DTOMagazine{}, shared.ErrDefault
	}
	return MockMagazines[0], nil
}

func (r repositoryMock) AddMagazine(ctx context.Context, _ DTOMagazine) (DTOMagazine, error) {
	if ctx.Value("AddMagazineError") == true {
		return DTOMagazine{}, shared.ErrDefault
	}
	return MockMagazines[0], nil
}

func (r repositoryMock) UpdateMagazine(ctx context.Context, _ string, _ DTOMagazine) (DTOMagazine, error) {
	if ctx.Value("UpdateMagazineError") == true {
		return DTOMagazine{}, shared.ErrDefault
	}
	return MockMagazines[0], nil
}

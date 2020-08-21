package material

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	uuid2 "github.com/google/uuid"
)

type Service interface {
	GetBooks(ctx context.Context) ([]Book, error)
	GetBookByCode(ctx context.Context, code string) (Book, error)
	AddBook(ctx context.Context, bookReq BookRequest) (Book, error)
	UpdateBook(ctx context.Context, code string, bookUpdate BookRequest) error
	DeleteBook(ctx context.Context, code string) error
}

type service struct {
	repository Repository
	logger     log.Logger
}

func (s service) GetBooks(ctx context.Context) ([]Book, error) {
	logger := log.With(s.logger, "method", "GetBooks")

	bookDTOs, err := s.repository.GetBooks(ctx)

	if err != nil {
		level.Error(logger).Log(err)
		return []Book{}, err
	}

	books := make([]Book, len(bookDTOs))
	for index, bookDTO := range bookDTOs {
		books[index] = Book{
			Material: Material{
				UniqueCode:    bookDTO.UniqueCode,
				Type:          bookDTO.Type,
				Name:          bookDTO.Name,
				EmissionDate:  bookDTO.EmissionDate,
				NumberOfPages: bookDTO.NumberOfPages,
			},
			Author: bookDTO.Author,
			Genre:  bookDTO.Genre,
		}
	}
	return books, nil
}

func (s service) GetBookByCode(ctx context.Context, code string) (Book, error) {
	logger := log.With(s.logger, "method", "GetBook by code")
	bookDTO, err := s.repository.GetBookByCode(ctx, code)
	if err != nil {
		level.Error(logger).Log(err)
		return Book{}, err
	}
	book := Book{
		Material: Material{
			UniqueCode:    bookDTO.UniqueCode,
			Type:          bookDTO.Type,
			Name:          bookDTO.Name,
			EmissionDate:  bookDTO.EmissionDate,
			NumberOfPages: bookDTO.NumberOfPages,
		},
		Author: bookDTO.Author,
		Genre:  bookDTO.Genre,
	}
	return book, nil
}

func (s service) AddBook(ctx context.Context, bookReq BookRequest) (Book, error) {
	logger := log.With(s.logger, "method", "Add book")

	newBook := BookDTO{
		MaterialDTO: MaterialDTO{
			UniqueCode:    uuid2.New().String(),
			Name:          bookReq.Name,
			NumberOfPages: bookReq.NumberOfPages,
			EmissionDate:  bookReq.EmissionDate,
			Type:          BookType,
		},
		Author: bookReq.Author,
		Genre:  bookReq.Genre,
	}
	bookDTO, err := s.repository.AddBook(ctx, newBook)
	if err != nil {
		level.Error(logger).Log(err)
		return Book{}, err
	}
	book := Book{
		Material: Material{
			UniqueCode:    bookDTO.UniqueCode,
			Type:          bookDTO.Type,
			Name:          bookDTO.Name,
			EmissionDate:  bookDTO.EmissionDate,
			NumberOfPages: bookDTO.NumberOfPages,
		},
		Author: bookDTO.Author,
		Genre:  bookDTO.Genre,
	}
	return book, nil
}

func (s service) UpdateBook(ctx context.Context, code string, bookUpdate BookRequest) error {
	logger := log.With(s.logger, "method", "Update book")

	updateBook := BookDTO{
		MaterialDTO: MaterialDTO{
			UniqueCode:    code,
			Name:          bookUpdate.Name,
			NumberOfPages: bookUpdate.NumberOfPages,
			EmissionDate:  bookUpdate.EmissionDate,
			Type:          BookType,
		},
		Author: bookUpdate.Author,
		Genre:  bookUpdate.Genre,
	}
	err := s.repository.UpdateBook(ctx, code, updateBook)
	if err != nil {
		level.Error(logger).Log(err)
		return err
	}
	return nil
}

func (s service) DeleteBook(ctx context.Context, code string) error {
	logger := log.With(s.logger, "method", "DeleteBook by code")
	err := s.repository.DeleteMaterial(ctx, code)
	if err != nil {
		level.Error(logger).Log(err)
	}
	return nil
}

func NewService(r Repository, logger log.Logger) Service {
	return &service{
		repository: r,
		logger:     logger,
	}
}

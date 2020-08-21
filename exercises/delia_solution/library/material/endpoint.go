package material

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetBooks      endpoint.Endpoint
	GetBookByCode endpoint.Endpoint
	AddBook       endpoint.Endpoint
	UpdateBook    endpoint.Endpoint
	DeleteBook    endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		GetBooks:      makeGetBooksEndpoint(s),
		GetBookByCode: makeGetBookByCodeEndpoint(s),
		AddBook:       makeAddBookEndpoint(s),
		UpdateBook:    makeUpdateBookEndpoint(s),
		DeleteBook:    makeDeleteBookEndpoint(s),
	}
}

// Delete Book Endpoint

func makeDeleteBookEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		bookCode := request.(string)
		err := s.DeleteBook(ctx, bookCode)
		return NoContentResponse{Err: err}, err
	}
}

// Update Book Endpoint

func makeUpdateBookEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateBookRequest)
		err := s.UpdateBook(ctx, req.Code, req.Book)
		return NoContentResponse{Err: err}, err
	}
}

// Add Book endpoint

func makeAddBookEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddBookRequest)
		newBook, err := s.AddBook(ctx, req.Book)
		return BookResponse{Err: err, Book: newBook}, err
	}
}

// Get Book by code endpoint

func makeGetBookByCodeEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(MaterialByCodeRequest)
		book, err := s.GetBookByCode(ctx, req.Id)
		return BookResponse{Book: book, Err: err}, err
	}
}

// Get Books endpoints

type getBooksResponse struct {
	Books []Book `json:"books,omitempty"`
	Err   error  `json:"err,omitempty"`
}

func makeGetBooksEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		books, err := s.GetBooks(ctx)
		return getBooksResponse{Books: books, Err: err}, err
	}
}

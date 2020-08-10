package library

import (
	"context"
	"github.com/go-kit/kit/log"
)

type Service interface {
	Method(ctx context.Context) (Library, error)
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

func (s service) Method(ctx context.Context) (Library, error) {
	result, err := s.repository.NewMethod(ctx)
	if err != nil {
		return Library{}, err
	}

	return s.dtoToLibrary(result), nil
}

func (s service) dtoToLibrary(f DTOLibrary) Library {
	return Library{
		Title:       f.Title,
		Description: f.Description,
	}
}

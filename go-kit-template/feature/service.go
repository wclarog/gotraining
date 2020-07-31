package feature

import (
	"context"
	"github.com/go-kit/kit/log"
)

type Service interface {
	Method(ctx context.Context) (Feature, error)
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

func (s service) Method(ctx context.Context) (Feature, error) {
	result, err := s.repository.NewMethod(ctx)
	if err != nil {
		return Feature{}, err
	}

	return s.dtoToFeature(result), nil
}

func (s service) dtoToFeature(f DTOFeature) Feature {
	return Feature{
		Title:       f.Title,
		Description: f.Description,
	}
}

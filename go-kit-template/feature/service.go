package feature

import (
	"context"
	"github.com/go-kit/kit/log"
	"go-kit-template/database"
	"go-kit-template/shared"
)

type Service interface {
	Method(ctx context.Context) (Feature, error)
	database.Transaction
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

func (s service) StartTx(ctx context.Context) (context.Context, error) {
	ctx, err := s.repository.StartTx(ctx)
	if err != nil {
		err = shared.NewApiError(shared.ErrDatabase.Error(), shared.Internal, "StartTx", shared.ServiceLevel, err)
	}
	return ctx, err
}

func (s service) Commit(ctx context.Context) error {
	err := s.repository.Commit(ctx)
	if err != nil {
		err = shared.NewApiError(shared.ErrDatabase.Error(), shared.Internal, "Commit", shared.ServiceLevel, err)
	}
	return err
}

func (s service) Rollback(ctx context.Context) error {
	err := s.repository.Rollback(ctx)
	if err != nil {
		err = shared.NewApiError(shared.ErrDatabase.Error(), shared.Internal, "Rollback", shared.ServiceLevel, err)
	}
	return err
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

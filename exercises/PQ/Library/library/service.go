package library

import (
	"context"
	"errors"
	"github.com/go-kit/kit/log"
)

type Service interface {
	GetMaterial(ctx context.Context, uniqueCode string) (material interface{}, err error)
	AddMaterial(ctx context.Context, material interface{}) error
	UpdateMaterial(ctx context.Context, material interface{}) error
	DeleteMaterial(ctx context.Context, uniqueCode string) error
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

func (s service) GetMaterial(ctx context.Context, uniqueCode string) (material interface{}, err error) {
	dtoResult, err := s.repository.GetMaterial(ctx, uniqueCode)

	if err != nil {
		return Material{}, err
	}

	return s.dtoToMaterial(dtoResult)
}

func (s service) AddMaterial(ctx context.Context, material interface{}) error {
	dtoMaterial, err := s.materialToDto(material)
	if err != nil {
		return err
	}

	_, err = s.repository.AddMaterial(ctx, dtoMaterial)

	return err
}

func (s service) UpdateMaterial(ctx context.Context, material interface{}) error {
	dtoMaterial, err := s.materialToDto(material)
	if err != nil {
		return err
	}

	_, err = s.repository.UpdateMaterial(ctx, dtoMaterial)

	return err
}

func (s service) DeleteMaterial(ctx context.Context, uniqueCode string) error {
	err := s.repository.DeleteMaterial(ctx, uniqueCode)

	return err
}

func (s service) materialToDto(m interface{}) (material interface{}, err error) {
	switch v := m.(type) {
	case Book:
		return DTOBook{
			DTOMaterial: DTOMaterial{
				UniqueCode:     v.UniqueCode,
				Name:           v.Name,
				DateOfEmission: v.DateOfEmission,
				NumberOfPages:  v.NumberOfPages,
			},
			AuthorName: v.AuthorName,
			Genre:      v.Genre,
		}, nil

	case Newspaper:
		return DTONewspaper{
			DTOMaterial: DTOMaterial{
				UniqueCode:     v.UniqueCode,
				Name:           v.Name,
				DateOfEmission: v.DateOfEmission,
				NumberOfPages:  v.NumberOfPages,
			},
			Url: v.Url,
		}, nil

	case Magazine:
		return DTOMagazine{
			DTOMaterial: DTOMaterial{
				UniqueCode:     v.UniqueCode,
				Name:           v.Name,
				DateOfEmission: v.DateOfEmission,
				NumberOfPages:  v.NumberOfPages,
			},
			Sections: v.Sections,
			Url:      v.Url,
		}, nil

	default:
		return Material{}, errors.New("invalid material object in materialToDto")
	}
}

func (s service) dtoToMaterial(m interface{}) (material interface{}, err error) {
	switch v := m.(type) {
	case DTOBook:
		return Book{
			Material: Material{
				UniqueCode:     v.UniqueCode,
				Name:           v.Name,
				DateOfEmission: v.DateOfEmission,
				NumberOfPages:  v.NumberOfPages,
			},
			AuthorName: v.AuthorName,
			Genre:      v.Genre,
		}, nil

	case DTONewspaper:
		return Newspaper{
			Material: Material{
				UniqueCode:     v.UniqueCode,
				Name:           v.Name,
				DateOfEmission: v.DateOfEmission,
				NumberOfPages:  v.NumberOfPages,
			},
			Url: v.Url,
		}, nil

	case DTOMagazine:
		return Magazine{
			Material: Material{
				UniqueCode:     v.UniqueCode,
				Name:           v.Name,
				DateOfEmission: v.DateOfEmission,
				NumberOfPages:  v.NumberOfPages,
			},
			Sections: v.Sections,
			Url:      v.Url,
		}, nil

	default:
		return Material{}, errors.New("invalid material object in dtoToMaterial")
	}
}

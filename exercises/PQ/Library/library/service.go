package library

import (
	"context"
	"errors"
	"github.com/go-kit/kit/log"
)

type Service interface {
	GetMaterials(ctx context.Context) (material []interface{}, err error)
	GetMaterialByCode(ctx context.Context, uniqueCode string) (material interface{}, err error)
	AddMaterial(ctx context.Context, material interface{}) (newMaterial interface{}, err error)
	UpdateMaterial(ctx context.Context, uniqueCode string, material interface{}) (newMaterial interface{}, err error)
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

func (s service) GetMaterials(ctx context.Context) (materials []interface{}, err error) {
	dtoResults, err := s.repository.GetMaterials(ctx)

	if err != nil {
		return nil, err
	}

	return s.dtoToMaterials(dtoResults)
}

func (s service) GetMaterialByCode(ctx context.Context, uniqueCode string) (material interface{}, err error) {
	dtoResult, err := s.repository.GetMaterialByCode(ctx, uniqueCode)

	if err != nil {
		return Material{}, err
	}

	return s.dtoToMaterial(dtoResult)
}

func (s service) AddMaterial(ctx context.Context, material interface{}) (newMaterial interface{}, err error) {
	var result interface{}

	dtoMaterial, err := s.materialToDto(material)
	if err != nil {
		return Material{}, err
	}

	result, err = s.repository.AddMaterial(ctx, dtoMaterial)
	if err != nil {
		return nil, err
	}

	newMaterial, err = s.dtoToMaterial(result)

	return newMaterial, err
}

func (s service) UpdateMaterial(ctx context.Context, uniqueCode string, material interface{}) (newMaterial interface{}, err error) {
	dtoMaterial, err := s.materialToDto(material)
	if err != nil {
		return nil, err
	}

	dtoMaterial, err = s.repository.UpdateMaterial(ctx, uniqueCode, dtoMaterial)
	if err != nil {
		return nil, err
	}

	newMaterial, err = s.dtoToMaterial(dtoMaterial)

	return newMaterial, err
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

func (s service) materialsToDto(ms []interface{}) (materials []interface{}, err error) {
	var errCast error
	dtoMaterials := make([]interface{}, len(ms))

	for idx, m := range ms {
		dtoMaterials[idx], errCast = s.materialToDto(m)

		if errCast != nil {
			return nil, errors.New("invalid material object in materialsToDtos")
		}
	}

	return dtoMaterials, nil
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

func (s service) dtoToMaterials(ms []interface{}) (materials []interface{}, err error) {
	var errCast error
	materials = make([]interface{}, len(ms))

	for idx, m := range ms {
		materials[idx], errCast = s.dtoToMaterial(m)

		if errCast != nil {
			return nil, errors.New("invalid material object in dtoToMaterials")
		}
	}

	return materials, nil
}

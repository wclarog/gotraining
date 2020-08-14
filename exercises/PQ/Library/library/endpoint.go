package library

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	AddMaterialEndpoint        endpoint.Endpoint
	UpdateMaterialEndpoint     endpoint.Endpoint
	DeleteMaterialEndpoint     endpoint.Endpoint
	GetMaterialsEndpoint       endpoint.Endpoint
	GetMaterialByCodeEndpoint  endpoint.Endpoint
	GetBooksEndpoint           endpoint.Endpoint
	GetBookByCodeEndpoint      endpoint.Endpoint
	GetMagazinesEndpoint       endpoint.Endpoint
	GetMagazineByCodeEndpoint  endpoint.Endpoint
	GetNewspapersEndpoint      endpoint.Endpoint
	GetNewspaperByCodeEndpoint endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		AddMaterialEndpoint:        makeAddMaterialEndpoint(s),
		UpdateMaterialEndpoint:     makeUpdateMaterialEndpoint(s),
		DeleteMaterialEndpoint:     makeDeleteMaterialEndpoint(s),
		GetMaterialsEndpoint:       makeGetMaterialsEndpoint(s),
		GetMaterialByCodeEndpoint:  makeGetMaterialByCodeEndpoint(s),
		GetBooksEndpoint:           makeGetBooksEndpoint(s),
		GetBookByCodeEndpoint:      makeGetBookByCodeEndpoint(s),
		GetMagazinesEndpoint:       makeGetMagazinesEndpoint(s),
		GetMagazineByCodeEndpoint:  makeGetMagazineByCodeEndpoint(s),
		GetNewspapersEndpoint:      makeGetNewspapersEndpoint(s),
		GetNewspaperByCodeEndpoint: makeGetNewspaperByCodeEndpoint(s),
	}
}

func makeAddMaterialEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(addMaterialRequest)

		switch req.MaterialType {
		case BookType:
			return s.AddBook(ctx, req.Book)

		case MagazineType:
			return s.AddMagazine(ctx, req.Magazine)

		case NewspaperType:
			return s.AddNewspaper(ctx, req.Newspaper)

		default:
			return nil, errors.New("invalid material object in AddMaterial")
		}
	}
}

func makeUpdateMaterialEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(updateMaterialRequest)

		switch req.MaterialType {
		case BookType:
			return s.UpdateBook(ctx, req.Code, req.Book)

		case MagazineType:
			return s.UpdateMagazine(ctx, req.Code, req.Magazine)

		case NewspaperType:
			return s.UpdateNewspaper(ctx, req.Code, req.Newspaper)

		default:
			return nil, errors.New("invalid material object in UpdateMaterial")
		}
	}
}

func makeDeleteMaterialEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(deleteMaterialRequest)
		err = s.DeleteMaterial(ctx, req.Code)
		return nil, err
	}
}

func makeGetMaterialsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		m, e := s.GetMaterials(ctx)
		return m, e
	}
}

func makeGetMaterialByCodeEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getByCodeRequest)
		return s.GetMaterialByCode(ctx, req.Code)
	}
}

func makeGetBooksEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return s.GetBooks(ctx)
	}
}

func makeGetBookByCodeEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getByCodeRequest)
		return s.GetBookByCode(ctx, req.Code)
	}
}

func makeGetMagazinesEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return s.GetMagazines(ctx)
	}
}

func makeGetMagazineByCodeEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getByCodeRequest)
		return s.GetMagazineByCode(ctx, req.Code)
	}
}

func makeGetNewspapersEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return s.GetNewspapers(ctx)
	}
}

func makeGetNewspaperByCodeEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getByCodeRequest)
		return s.GetNewspaperByCode(ctx, req.Code)
	}
}

// Request types

type addMaterialRequest struct {
	MaterialType MaterialType
	Book         Book
	Magazine     Magazine
	Newspaper    Newspaper
}

type updateMaterialRequest struct {
	Code         string
	MaterialType MaterialType
	Book         Book
	Magazine     Magazine
	Newspaper    Newspaper
}

type deleteMaterialRequest struct {
	Code string
}

type getAllRequest struct {
}

type getByCodeRequest struct {
	Code string
}

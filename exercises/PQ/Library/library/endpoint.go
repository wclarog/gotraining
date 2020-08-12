package library

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetMaterialsEndpoint      endpoint.Endpoint
	GetMaterialByCodeEndpoint endpoint.Endpoint
	AddMaterialEndpoint       endpoint.Endpoint
	UpdateMaterialEndpoint    endpoint.Endpoint
	DeleteMaterialEndpoint    endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		GetMaterialsEndpoint:      makeGetMaterialsEndpoint(s),
		GetMaterialByCodeEndpoint: makeGetMaterialByCodeEndpoint(s),
		AddMaterialEndpoint:       makeAddMaterialEndpoint(s),
		UpdateMaterialEndpoint:    makeUpdateMaterialEndpoint(s),
		DeleteMaterialEndpoint:    makeDeleteMaterialEndpoint(s),
	}
}

func makeGetMaterialsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		//_ := request.(getMaterialsRequest)
		m, e := s.GetMaterials(ctx)
		// return getMaterialsResponse{Materials: m, Err: e}, nil
		return m, e
	}
}

func makeGetMaterialByCodeEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getMaterialByCodeRequest)
		m, e := s.GetMaterialByCode(ctx, req.Code)
		// return getMaterialByCodeResponse{Material: m, Err: e}, nil
		return m, e
	}
}

func makeAddMaterialEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(addMaterialRequest)
		m, e := s.AddMaterial(ctx, req.Material)
		//return addMaterialResponse{Material: m, Err: e}, nil
		return m, e
	}
}

func makeUpdateMaterialEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(updateMaterialRequest)
		m, e := s.UpdateMaterial(ctx, req.Code, req.Material)
		//return updateMaterialResponse{Material: m, Err: e}, nil
		return m, e
	}
}

func makeDeleteMaterialEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(deleteMaterialRequest)
		err = s.DeleteMaterial(ctx, req.Code)
		//return deleteMaterialResponse{Err: e}, nil
		return nil, err
	}
}

type getMaterialsRequest struct {
}

type getMaterialsResponse struct {
	Materials []interface{} `json:"materials,omitempty"`
	Err       error         `json:"err,omitempty"`
}

func (r getMaterialsResponse) error() error { return r.Err }

type getMaterialByCodeRequest struct {
	Code string
}
type getMaterialByCodeResponse struct {
	Material interface{} `json:"material,omitempty"`
	Err      error       `json:"err,omitempty"`
}

func (r getMaterialByCodeResponse) error() error { return r.Err }

type addMaterialRequest struct {
	Material interface{}
}

type addMaterialResponse struct {
	Material interface{} `json:"material,omitempty"`
	Err      error       `json:"err,omitempty"`
}

func (r addMaterialResponse) error() error { return r.Err }

type updateMaterialRequest struct {
	Code     string
	Material interface{}
}

type updateMaterialResponse struct {
	Material interface{} `json:"material,omitempty"`
	Err      error       `json:"err,omitempty"`
}

func (r updateMaterialResponse) error() error { return r.Err }

type deleteMaterialRequest struct {
	Code string
}

type deleteMaterialResponse struct {
	Err error `json:"err,omitempty"`
}

func (r deleteMaterialResponse) error() error { return r.Err }

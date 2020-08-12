package library

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetMaterialEndpoint endpoint.Endpoint
	AddMaterialEndpoint endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		GetMaterialEndpoint: makeGetMaterialEndpoint(s),
		AddMaterialEndpoint: makeAddMaterialEndpoint(s),
	}
}

func makeGetMaterialEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getMaterialRequest)
		m, e := s.GetMaterial(ctx, req.Code)
		return getMaterialResponse{Material: m, Err: e}, nil
	}
}

func makeAddMaterialEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(postMaterialRequest)
		e := s.AddMaterial(ctx, req.Material)
		return postMaterialResponse{Err: e}, nil
	}
}

type getMaterialRequest struct {
	Code string
}

type getMaterialResponse struct {
	Material interface{} `json:"material,omitempty"`
	Err      error       `json:"err,omitempty"`
}

func (r getMaterialResponse) error() error { return r.Err }

type postMaterialRequest struct {
	Material interface{}
}

type postMaterialResponse struct {
	Err error `json:"err,omitempty"`
}

func (r postMaterialResponse) error() error { return r.Err }

package library

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	Method endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		Method: makeMethodEndpoint(s),
	}
}

func makeMethodEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		events, err := s.Method(ctx)
		return MethodResponse(events), err
	}
}

type (
	MethodResponse Library
)

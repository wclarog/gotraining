package event

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
		events, err := s.NextEvents(ctx)

		if err != nil {
			return "", err
		}

		return MethodResponse(events), nil
	}
}

type (
	MethodResponse []Event
)

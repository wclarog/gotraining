package feature

import (
	"context"
	"go-kit-template/database"
	"go-kit-template/shared"

	"github.com/go-kit/kit/endpoint"
)

func NewTXMiddleware(s Service, e Endpoints) Endpoints {
	return Endpoints{
		Method: newTXMiddleware(s)(e.Method),
	}
}

func newTXMiddleware(s Service) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			ctx, err = s.StartTx(ctx)
			if err != nil {
				return nil, err
			}
			defer func() {
				err = database.CommitOrRollback(ctx, s, err)
				if err != nil {
					err = shared.NewApiError(err.Error(), shared.Internal, "CommitOrRollback", shared.TxMiddlewareLevel, err)
				}
			}()

			response, err = next(ctx, request)
			return response, err
		}
	}
}

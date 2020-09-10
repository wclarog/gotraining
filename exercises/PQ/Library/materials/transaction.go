package materials

import (
	"context"
	"excercise-library/database"
	"excercise-library/shared"

	"github.com/go-kit/kit/endpoint"
)

func NewTXMiddleware(s Service, e Endpoints) Endpoints {
	return Endpoints{
		AddMaterialEndpoint:        newTXMiddleware(s)(e.AddMaterialEndpoint),
		UpdateMaterialEndpoint:     newTXMiddleware(s)(e.UpdateMaterialEndpoint),
		DeleteMaterialEndpoint:     newTXMiddleware(s)(e.DeleteMaterialEndpoint),
		GetMaterialsEndpoint:       e.GetMaterialsEndpoint,
		GetMaterialByCodeEndpoint:  e.GetMaterialByCodeEndpoint,
		GetBooksEndpoint:           e.GetBooksEndpoint,
		GetBookByCodeEndpoint:      e.GetBookByCodeEndpoint,
		GetMagazinesEndpoint:       e.GetMagazinesEndpoint,
		GetMagazineByCodeEndpoint:  e.GetMagazineByCodeEndpoint,
		GetNewspapersEndpoint:      e.GetNewspapersEndpoint,
		GetNewspaperByCodeEndpoint: e.GetNewspaperByCodeEndpoint,
	}
}

/*
func NewTXMiddleware(s Service, e Endpoints) Endpoints {
	return Endpoints{
		AddMaterialEndpoint:        newTXMiddleware(s)(e.AddMaterialEndpoint),
		UpdateMaterialEndpoint:     newTXMiddleware(s)(e.UpdateMaterialEndpoint),
		DeleteMaterialEndpoint:     newTXMiddleware(s)(e.DeleteMaterialEndpoint),
		GetMaterialsEndpoint:       newTXMiddleware(s)(e.GetMaterialsEndpoint),
		GetMaterialByCodeEndpoint:  newTXMiddleware(s)(e.GetMaterialByCodeEndpoint),
		GetBooksEndpoint:           newTXMiddleware(s)(e.GetBooksEndpoint),
		GetBookByCodeEndpoint:      newTXMiddleware(s)(e.GetBookByCodeEndpoint),
		GetMagazinesEndpoint:       newTXMiddleware(s)(e.GetMagazinesEndpoint),
		GetMagazineByCodeEndpoint:  newTXMiddleware(s)(e.GetMagazineByCodeEndpoint),
		GetNewspapersEndpoint:      newTXMiddleware(s)(e.GetNewspapersEndpoint),
		GetNewspaperByCodeEndpoint: newTXMiddleware(s)(e.GetNewspaperByCodeEndpoint),
	}
}
*/

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

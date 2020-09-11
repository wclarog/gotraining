package materials

import (
	"context"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

func NewHandler(_ context.Context, options []httptransport.ServerOption, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	errorEncoder := httptransport.ServerErrorEncoder(encodeErrorResponse)
	insertJWT := httptransport.ServerBefore(InsertJwtIntoContext())

	options = append(options, errorEncoder, insertJWT)

	r.Methods(http.MethodPost).Path("/library/material").Handler(httptransport.NewServer(
		endpoints.AddMaterialEndpoint,
		decodeAddMaterialRequest,
		encodeResponseCreated,
		options...,
	))

	r.Methods(http.MethodPut).Path("/library/material/{code}").Handler(httptransport.NewServer(
		endpoints.UpdateMaterialEndpoint,
		decodeUpdateMaterialRequest,
		encodeResponseNoContent,
		options...,
	))

	r.Methods(http.MethodDelete).Path("/library/material/{code}").Handler(httptransport.NewServer(
		endpoints.DeleteMaterialEndpoint,
		decodeDeleteMaterialRequest,
		encodeResponseNoContent,
		options...,
	))

	r.Methods(http.MethodGet).Path("/library/material").Handler(httptransport.NewServer(
		endpoints.GetMaterialsEndpoint,
		decodeGetAllRequest,
		encodeResponseOK,
		options...,
	))

	r.Methods(http.MethodGet).Path("/library/material/{code}").Handler(httptransport.NewServer(
		endpoints.GetMaterialByCodeEndpoint,
		decodeGetByCodeRequest,
		encodeResponseOK,
		options...,
	))

	r.Methods(http.MethodGet).Path("/library/book").Handler(httptransport.NewServer(
		endpoints.GetBooksEndpoint,
		decodeGetAllRequest,
		encodeResponseOK,
		options...,
	))

	r.Methods(http.MethodGet).Path("/library/book/{code}").Handler(httptransport.NewServer(
		endpoints.GetBookByCodeEndpoint,
		decodeGetByCodeRequest,
		encodeResponseOK,
		options...,
	))

	r.Methods(http.MethodGet).Path("/library/magazine").Handler(httptransport.NewServer(
		endpoints.GetMagazinesEndpoint,
		decodeGetAllRequest,
		encodeResponseOK,
		options...,
	))

	r.Methods(http.MethodGet).Path("/library/magazine/{code}").Handler(httptransport.NewServer(
		endpoints.GetMagazineByCodeEndpoint,
		decodeGetByCodeRequest,
		encodeResponseOK,
		options...,
	))

	r.Methods(http.MethodGet).Path("/library/newspaper").Handler(httptransport.NewServer(
		endpoints.GetNewspapersEndpoint,
		decodeGetAllRequest,
		encodeResponseOK,
		options...,
	))

	r.Methods(http.MethodGet).Path("/library/newspaper/{code}").Handler(httptransport.NewServer(
		endpoints.GetNewspaperByCodeEndpoint,
		decodeGetByCodeRequest,
		encodeResponseOK,
		options...,
	))

	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

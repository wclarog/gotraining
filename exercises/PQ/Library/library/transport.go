package library

import (
	"context"
	kitTransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

func NewHandler(_ context.Context, options []kitTransport.ServerOption, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	errorEncoder := kitTransport.ServerErrorEncoder(encodeErrorResponse)
	insertJWT := kitTransport.ServerBefore(InsertJwtIntoContext())

	options = append(options, errorEncoder, insertJWT)

	r.Methods(http.MethodPost).Path("/library/material").Handler(kitTransport.NewServer(
		endpoints.AddMaterialEndpoint,
		decodeAddMaterialRequest,
		encodeResponseCreated,
		options...,
	))

	r.Methods(http.MethodPut).Path("/library/material/{code}").Handler(kitTransport.NewServer(
		endpoints.UpdateMaterialEndpoint,
		decodeUpdateMaterialRequest,
		encodeResponseNoContent,
		options...,
	))

	r.Methods(http.MethodDelete).Path("/library/material/{code}").Handler(kitTransport.NewServer(
		endpoints.DeleteMaterialEndpoint,
		decodeDeleteMaterialRequest,
		encodeResponseNoContent,
		options...,
	))

	r.Methods(http.MethodGet).Path("/library/material/material").Handler(kitTransport.NewServer(
		endpoints.GetMaterialsEndpoint,
		decodeGetAllRequest,
		encodeResponseOK,
		options...,
	))

	r.Methods(http.MethodGet).Path("/library/material/material/{code}").Handler(kitTransport.NewServer(
		endpoints.GetMaterialByCodeEndpoint,
		decodeGetByCodeRequest,
		encodeResponseOK,
		options...,
	))

	r.Methods(http.MethodGet).Path("/library/material/book").Handler(kitTransport.NewServer(
		endpoints.GetBooksEndpoint,
		decodeGetAllRequest,
		encodeResponseOK,
		options...,
	))

	r.Methods(http.MethodGet).Path("/library/material/book/{code}").Handler(kitTransport.NewServer(
		endpoints.GetBookByCodeEndpoint,
		decodeGetByCodeRequest,
		encodeResponseOK,
		options...,
	))

	r.Methods(http.MethodGet).Path("/library/material/magazine").Handler(kitTransport.NewServer(
		endpoints.GetMagazinesEndpoint,
		decodeGetAllRequest,
		encodeResponseOK,
		options...,
	))

	r.Methods(http.MethodGet).Path("/library/material/magazine/{code}").Handler(kitTransport.NewServer(
		endpoints.GetMagazineByCodeEndpoint,
		decodeGetByCodeRequest,
		encodeResponseOK,
		options...,
	))

	r.Methods(http.MethodGet).Path("/library/material/newspaper").Handler(kitTransport.NewServer(
		endpoints.GetNewspapersEndpoint,
		decodeGetAllRequest,
		encodeResponseOK,
		options...,
	))

	r.Methods(http.MethodGet).Path("/library/material/newspaper/{code}").Handler(kitTransport.NewServer(
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

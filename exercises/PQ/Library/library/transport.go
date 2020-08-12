package library

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

	options = append(options, errorEncoder)

	r.Methods(http.MethodGet).Path("/library/materials").Handler(httptransport.NewServer(
		endpoints.GetMaterialsEndpoint,
		decodeGetMaterialsRequest,
		encodeResponse,
		options...,
	))

	r.Methods(http.MethodGet).Path("/library/materials/{code}").Handler(httptransport.NewServer(
		endpoints.GetMaterialByCodeEndpoint,
		decodeGetMaterialByCodeRequest,
		encodeResponse,
		options...,
	))

	r.Methods(http.MethodPost).Path("/library/materials").Handler(httptransport.NewServer(
		endpoints.AddMaterialEndpoint,
		decodeAddMaterialRequest,
		encodeResponse,
		options...,
	))

	r.Methods(http.MethodPut).Path("/library/materials/{code}").Handler(httptransport.NewServer(
		endpoints.UpdateMaterialEndpoint,
		decodeUpdateMaterialRequest,
		encodeResponse,
		options...,
	))

	r.Methods(http.MethodDelete).Path("/library/materials/{code}").Handler(httptransport.NewServer(
		endpoints.DeleteMaterialEndpoint,
		decodeDeleteMaterialRequest,
		encodeResponse,
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

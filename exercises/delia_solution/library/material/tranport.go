package material

import (
	"context"
	"encoding/json"
	"errors"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

type errorer interface {
	error() error
}

func NewHandler(_ context.Context, options []httptransport.ServerOption, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	errorEncoder := httptransport.ServerErrorEncoder(encodeErrorResponse)

	options = append(options, errorEncoder)

	r.Methods(http.MethodGet).Path("/books").Handler(httptransport.NewServer(
		endpoints.GetBooks,
		decodeListRequest,
		encodeResponse,
		options...,
	))

	r.Methods(http.MethodGet).Path("/books/{id}").Handler(httptransport.NewServer(
		endpoints.GetBookByCode,
		decodeMaterialByCodeRequest,
		encodeResponse,
		options...,
	))

	r.Methods(http.MethodPost).Path("/book").Handler(httptransport.NewServer(
		endpoints.AddBook,
		decodeAddBookRequest,
		encodeCreatedResponse,
		options...,
	))

	r.Methods(http.MethodPut).Path("/books/{id}").Handler(httptransport.NewServer(
		endpoints.UpdateBook,
		decodeUpdateBookRequest,
		encodeNoContentResponse,
		options...,
	))

	r.Methods(http.MethodDelete).Path("/books/{id}").Handler(httptransport.NewServer(
		endpoints.UpdateBook,
		decodeMaterialByCodeRequest,
		encodeNoContentResponse,
		options...,
	))

	return r
}

// Encode Responses

func encodeResponse(ctx context.Context, writer http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeErrorResponse(ctx, e.error(), writer)
		return nil
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(writer).Encode(response)
}

func encodeCreatedResponse(_ context.Context, rw http.ResponseWriter, response interface{}) error {
	rw.WriteHeader(http.StatusCreated)
	return json.NewEncoder(rw).Encode(response)
}

func encodeNoContentResponse(_ context.Context, rw http.ResponseWriter, response interface{}) error {
	rw.WriteHeader(http.StatusNoContent)
	return json.NewEncoder(rw).Encode(response)
}

func encodeErrorResponse(ctx context.Context, err error, w http.ResponseWriter) {
	/*w.WriteHeader(http.StatusInternalServerError)
	_ = json.NewEncoder(w).Encode(err)*/
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err.(type) {
	case *NotFoundError:
		w.WriteHeader(http.StatusNotFound)
	case *InvalidBookTypeError:
		w.WriteHeader(http.StatusExpectationFailed)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

// Decode

func decodeMaterialByCodeRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req MaterialByCodeRequest

	req.Id, err = decodeCodeParamRequest(r)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func decodeAddBookRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req AddBookRequest

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func decodeUpdateBookRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req UpdateBookRequest

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func decodeListRequest(_ context.Context, _ *http.Request) (request interface{}, err error) {
	return GetRequest{}, nil
}

func decodeCodeParamRequest(r *http.Request) (codeParam string, err error) {
	var ok bool

	vars := mux.Vars(r)
	codeParam, ok = vars["id"]
	if !ok {
		return "", errors.New("Expected parameter not provided.")
	}

	return codeParam, nil
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

package library

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

var (
	// ErrBadRouting is returned when an expected path variable is missing.
	// It always indicates programmer error.
	ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")
)

func encodeResponse(_ context.Context, rw http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(rw).Encode(response)
}

func encodeErrorResponse(_ context.Context, err error, rw http.ResponseWriter) {
	if strings.Index(err.Error(), "not found") >= 0 {
		rw.WriteHeader(http.StatusNotFound)
	} else {
		rw.WriteHeader(http.StatusInternalServerError)
	}

	_ = json.NewEncoder(rw).Encode(err.Error())
}

func decodeGetMaterialsRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	return getMaterialsRequest{}, nil
}

func decodeGetMaterialByCodeRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	code, ok := vars["code"]
	if !ok {
		return nil, ErrBadRouting
	}
	return getMaterialByCodeRequest{Code: code}, nil
}

func decodeAddMaterialRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req addMaterialRequest

	req.Material, err = decodeMaterialFromBody(r)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func decodeUpdateMaterialRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req updateMaterialRequest
	var ok bool
	vars := mux.Vars(r)
	req.Code, ok = vars["code"]
	if !ok {
		return updateMaterialRequest{}, ErrBadRouting
	}

	req.Material, err = decodeMaterialFromBody(r)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func decodeDeleteMaterialRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	code, ok := vars["code"]
	if !ok {
		return nil, ErrBadRouting
	}
	return deleteMaterialRequest{Code: code}, nil
}

func decodeMaterialFromBody(r *http.Request) (request interface{}, err error) {
	var book Book
	var magazine Magazine
	var newspaper Newspaper

	if err = json.NewDecoder(r.Body).Decode(&book); err == nil {
		return book, nil
	}

	if err = json.NewDecoder(r.Body).Decode(&magazine); err == nil {
		return magazine, nil
	}

	if err = json.NewDecoder(r.Body).Decode(&newspaper); err == nil {
		return newspaper, nil
	}

	return nil, errors.New("invalid format of request body (invalid material)")
}

/*
func decodeGetMaterialRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req string
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func decodeAddMaterialRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req Material
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}
*/

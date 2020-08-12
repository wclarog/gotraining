package library

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
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
	rw.WriteHeader(http.StatusInternalServerError)
	_ = json.NewEncoder(rw).Encode(err.Error())
}

func decodeGetMaterialRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	code, ok := vars["code"]
	if !ok {
		return nil, ErrBadRouting
	}
	return getMaterialRequest{Code: code}, nil
}

func decodeAddMaterialRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req postMaterialRequest
	if e := json.NewDecoder(r.Body).Decode(&req.Material); e != nil {
		return nil, e
	}
	return req, nil
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

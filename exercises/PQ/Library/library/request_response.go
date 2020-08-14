package library

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	// ErrBadRouting is returned when an expected path variable is missing.
	// It always indicates programmer error.
	ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")
)

func encodeResponseOK(_ context.Context, rw http.ResponseWriter, response interface{}) error {
	rw.WriteHeader(http.StatusCreated)
	return json.NewEncoder(rw).Encode(response)
}

func encodeResponseCreated(_ context.Context, rw http.ResponseWriter, response interface{}) error {
	rw.WriteHeader(http.StatusCreated)
	return json.NewEncoder(rw).Encode(response)
}

func encodeResponseNoContent(_ context.Context, rw http.ResponseWriter, response interface{}) error {
	rw.WriteHeader(http.StatusNoContent)
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

func decodeAddMaterialRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req addMaterialRequest
	var mat Material
	var bodyData []byte

	bodyData, err = ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(bytes.NewReader(bodyData)).Decode(&mat)
	if err != nil {
		return nil, err
	}

	req.MaterialType = mat.MaterialType

	switch req.MaterialType {
	case BookType:
		err = json.NewDecoder(bytes.NewReader(bodyData)).Decode(&req.Book)

	case MagazineType:
		err = json.NewDecoder(bytes.NewReader(bodyData)).Decode(&req.Magazine)

	case NewspaperType:
		err = json.NewDecoder(bytes.NewReader(bodyData)).Decode(&req.Newspaper)

	default:
		return nil, errors.New("invalid material object in decodeAddMaterialRequest")
	}

	if err != nil {
		return nil, err
	}

	return req, nil
}

func decodeUpdateMaterialRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req updateMaterialRequest
	var mat Material
	var ok bool
	var bodyData []byte

	vars := mux.Vars(r)
	req.Code, ok = vars["code"]
	if !ok {
		return nil, ErrBadRouting
	}

	bodyData, err = ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(bytes.NewReader(bodyData)).Decode(&mat)
	if err != nil {
		return nil, err
	}

	req.MaterialType = mat.MaterialType

	switch req.MaterialType {
	case BookType:
		err = json.NewDecoder(bytes.NewReader(bodyData)).Decode(&req.Book)

	case MagazineType:
		err = json.NewDecoder(bytes.NewReader(bodyData)).Decode(&req.Magazine)

	case NewspaperType:
		err = json.NewDecoder(bytes.NewReader(bodyData)).Decode(&req.Newspaper)

	default:
		return nil, errors.New("invalid material object in decodeAddMaterialRequest")
	}

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

func decodeGetAllRequest(_ context.Context, _ *http.Request) (request interface{}, err error) {
	return getAllRequest{}, nil
}

func decodeGetByCodeRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	code, ok := vars["code"]
	if !ok {
		return nil, ErrBadRouting
	}
	return getByCodeRequest{Code: code}, nil
}

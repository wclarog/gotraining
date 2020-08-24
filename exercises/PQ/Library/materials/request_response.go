package materials

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

// Request decoders

func decodeAddMaterialRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req addMaterialRequest

	req.materialBodyRequest, err = decodeMaterialBodyRequest(r)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func decodeUpdateMaterialRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req updateMaterialRequest

	req.materialBodyRequest, err = decodeMaterialBodyRequest(r)
	if err != nil {
		return nil, err
	}

	req.codeParamRequest, err = decodeCodeParamRequest(r)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func decodeDeleteMaterialRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req deleteMaterialRequest

	req.codeParamRequest, err = decodeCodeParamRequest(r)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func decodeGetAllRequest(_ context.Context, _ *http.Request) (request interface{}, err error) {
	return getAllRequest{}, nil
}

func decodeGetByCodeRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req getByCodeRequest

	req.codeParamRequest, err = decodeCodeParamRequest(r)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// Generic decoders

func decodeCodeParamRequest(r *http.Request) (codeParam codeParamRequest, err error) {
	var ok bool

	vars := mux.Vars(r)
	codeParam.Code, ok = vars["code"]
	if !ok {
		return codeParamRequest{}, ErrBadRouting
	}

	return codeParam, nil
}

func decodeMaterialBodyRequest(r *http.Request) (materialBody materialBodyRequest, err error) {
	var mat Material
	var bodyData []byte

	bodyData, err = ioutil.ReadAll(r.Body)
	if err != nil {
		return materialBodyRequest{}, err
	}

	err = json.NewDecoder(bytes.NewReader(bodyData)).Decode(&mat)
	if err != nil {
		return materialBodyRequest{}, err
	}

	materialBody.MaterialType = mat.MaterialType

	switch materialBody.MaterialType {
	case BookType:
		err = json.NewDecoder(bytes.NewReader(bodyData)).Decode(&materialBody.Book)

	case MagazineType:
		err = json.NewDecoder(bytes.NewReader(bodyData)).Decode(&materialBody.Magazine)

	case NewspaperType:
		err = json.NewDecoder(bytes.NewReader(bodyData)).Decode(&materialBody.Newspaper)

	default:
		return materialBodyRequest{}, errors.New("invalid material object in decodeAddMaterialRequest")
	}

	if err != nil {
		return materialBodyRequest{}, err
	}

	return materialBody, nil
}

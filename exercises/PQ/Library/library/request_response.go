package library

import (
	"context"
	"encoding/json"
	"net/http"
)

func encodeResponse(_ context.Context, rw http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(rw).Encode(response)
}

func encodeErrorResponse(_ context.Context, err error, rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusInternalServerError)
	_ = json.NewEncoder(rw).Encode(err.Error())
}

func decodeLibraryRequest(_ context.Context, _ *http.Request) (request interface{}, err error) {
	return
}

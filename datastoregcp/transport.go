package datastoregcp

import (
	"context"
	"encoding/json"
	"net/http"
	"cloud.google.com/go/datastore"
)

type getRequest struct {
	EntityName	string `json:"entityName"`
	Db	*datastore.Client	`json:"datastoreClient"`
}

type getResponse struct {
	Entity interface{} `json:"entity"`
	Err string `json:"err,omitempty"`
}

type storeRequest struct{
	EntityName	string `json:"entityName"`
	Entity	interface{} `json:"Entity"`
	Db	*datastore.Client	`json:"datastoreClient"`
}

type storeResponse struct {
	Status string `json:"status"`
	Err   string `json:"err,omitempty"`
}

func decodeGetRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req getRequest
	return req, nil
}

func decodeStoreRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req storeRequest
	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
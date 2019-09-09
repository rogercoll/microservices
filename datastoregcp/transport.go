package datastoregcp

import (
	"context"
	"encoding/json"
	"net/http"
	"cloud.google.com/go/datastore"
)

type getRequest struct {
	ProjID	string	`json:"projID"`
	EntityName	string `json:"entityName"`
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
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeStoreRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req storeRequest
	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
package datastoregcp

import (
	"context"
    "errors"
	"github.com/go-kit/kit/endpoint"
	"cloud.google.com/go/datastore"
)

type Endpoints struct {
	GetEndpoint			endpoint.Endpoint
	StoreEndpoint		endpoint.Endpoint
}


func MakeGetEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getRequest)
		d, err := srv.GetObject(ctx,req.Db,req.EntityName)
		if err != nil {
			return getResponse{d, err.Error()}, nil
		}
		return getResponse{d, ""}, nil
	}
}

func MakeStoreEndpoint(srv Service) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(storeRequest)
        s, err := srv.StoreObject(ctx,req.Db,req.EntityName,req.Entity)
        if err != nil {
            return storeResponse{s ,err.Error()}, nil
        }
        return storeResponse{s, ""}, nil
    }
}

func (e Endpoints) GetObject(ctx context.Context, db *datastore.Client, entity string) (interface{}, error) {
	req := getRequest{EntityName: entity, Db: db}
	resp, err := e.GetEndpoint(ctx, req)
	if err != nil {
		return "", err
	}
	getResp := resp.(getResponse)
	if getResp.Err != "" {
		return "", errors.New(getResp.Err)
	}
	return getResp.Entity, nil
}

func (e Endpoints) StoreObject(ctx context.Context, db *datastore.Client, entityName string, entity interface{}) (string, error) {
    req := storeRequest{EntityName: entityName, Entity: entity, Db: db}
    resp, err := e.StoreEndpoint(ctx, req)
    if err != nil {
        return "", err
    }
    storeResp := resp.(storeResponse)
    return storeResp.Status, nil
}
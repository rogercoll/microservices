package datastoregcp

import (
	"context"
	"cloud.google.com/go/datastore"
)

type Service interface {
	GetObject(ctx context.Context, db *datastore.Client, entity string) ([]interface{}, error)
	StoreObject(ctx context.Context, db *datastore.Client, entityName string, entity interface{}) (string, error)
}

type cDataStoreService struct {}

func NewService() Service {
	return cDataStoreService{}
}

func (cDataStoreService) GetObject(ctx context.Context, db *datastore.Client, entity string) ([]interface{}, error) {
	result := make([]interface{},0)
	q := datastore.NewQuery(entity)
	_, err := db.GetAll(ctx, q, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (cDataStoreService) StoreObject(ctx context.Context, db *datastore.Client, entityName string, entity interface{}) (string, error) {
	k := datastore.IncompleteKey(entityName, nil)
	k, err := db.Put(ctx, k, entity)
	if err != nil {
		return "", err
	}
	return "Object has been stored!", nil
}
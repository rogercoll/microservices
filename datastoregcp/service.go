package datastoregcp

import (
	"fmt"
	"context"
	"cloud.google.com/go/datastore"
)

type Service interface {
	GetObject(ctx context.Context, projID string, entity string) ([]interface{}, error)
	StoreObject(ctx context.Context, db *datastore.Client, entityName string, entity interface{}) (string, error)
}

type cDataStoreService struct {}

func NewService() Service {
	return cDataStoreService{}
}

func (cDataStoreService) GetObject(ctx context.Context, projID string, entity string) ([]interface{}, error) {
	result := make([]interface{},0)
	fmt.Println(projID)
	fmt.Println(entity)
	db, err := datastore.NewClient(ctx, projID)
	if err != nil {
		return nil, err
	}
	q := datastore.NewQuery(entity)
	_, err = db.GetAll(ctx, q, result)
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
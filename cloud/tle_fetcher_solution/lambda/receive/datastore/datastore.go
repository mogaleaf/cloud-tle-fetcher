package datastore

import (
	"tle_fetcher_solution/shared/db"
	"tle_fetcher_solution/shared/db/model"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type datastore struct {
	client *dynamodb.DynamoDB
}

type Datastore interface {
	GetRegisteredConnection(satellite string) ([]model.DBConnectedClient, error)
}

func NewDatastore(region string) (Datastore, error) {
	return &datastore{
		client: db.NewDynamoDB(region),
	}, nil
}

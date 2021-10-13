package datastore

import (
	"tle_fetcher_solution/lambda/fetch/model"
	"tle_fetcher_solution/shared/db"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type datastore struct {
	client *dynamodb.DynamoDB
}

type Datastore interface {
	InsertTle(tles map[string]*model.Tle) error
	GetSatelliteNames() ([]string, error)
}

func NewDatastore(region string) (Datastore, error) {
	return &datastore{
		client: db.NewDynamoDB(region),
	}, nil
}

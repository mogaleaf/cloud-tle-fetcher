package datastore

import (
	"tle_fetcher_solution/shared/db"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type datastore struct {
	client *dynamodb.DynamoDB
}

type Datastore interface {
	RegisterConnection(connectionID, appid, stage string, satellite []string) error
}

func NewDatastore(region string) (Datastore, error) {
	return &datastore{
		client: db.NewDynamoDB(region),
	}, nil
}

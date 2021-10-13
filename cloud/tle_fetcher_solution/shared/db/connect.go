package db

import (
	"tle_fetcher_solution/shared/session"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func NewDynamoDB(region string) *dynamodb.DynamoDB {
	return dynamodb.New(session.NewSession(region))
}

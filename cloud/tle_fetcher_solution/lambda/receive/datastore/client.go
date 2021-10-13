package datastore

import (
	"fmt"
	"tle_fetcher_solution/shared/db"
	"tle_fetcher_solution/shared/db/model"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

const (
	connection_table_name = "Connection"
)

func (d *datastore) GetRegisteredConnection(satellite string) ([]model.DBConnectedClient, error) {
	contains := expression.Contains(expression.Name("Satellites"), satellite)
	build, err := expression.NewBuilder().WithFilter(contains).Build()
	if err != nil {
		return nil, fmt.Errorf("cant filter: %w", err)
	}
	scan, err := d.client.Scan(&dynamodb.ScanInput{
		FilterExpression:          build.Filter(),
		ExpressionAttributeNames:  build.Names(),
		ExpressionAttributeValues: build.Values(),
		TableName:                 aws.String(connection_table_name),
	})
	err = db.EscapeNoRowError(err)
	if err != nil {
		return nil, fmt.Errorf("can't fin object: %w", err)
	}
	var connections []model.DBConnectedClient
	for _, item := range scan.Items {
		con := model.DBConnectedClient{}
		err = dynamodbattribute.UnmarshalMap(item, &con)
		if err != nil {
			return nil, fmt.Errorf("cant unmarshall: %w", err)
		}
		connections = append(connections, con)
	}
	return connections, nil
}

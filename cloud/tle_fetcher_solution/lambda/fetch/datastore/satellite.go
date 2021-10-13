package datastore

import (
	"fmt"
	"tle_fetcher_solution/shared/db/model"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const (
	sat_table_name = "Satellite"
)

func (db *datastore) GetSatelliteNames() ([]string, error) {
	var names []string
	scan, err := db.client.Scan(&dynamodb.ScanInput{
		TableName: aws.String(sat_table_name),
	})
	if err != nil {
		return nil, fmt.Errorf("can't select satellites: %w", err)
	}
	for _, item := range scan.Items {
		sat := model.Satellite{}
		err = dynamodbattribute.UnmarshalMap(item, &sat)
		if err != nil {
			return nil, fmt.Errorf("cant unmarshall: %w", err)
		}
		names = append(names, sat.SatelliteName)
	}
	return names, nil
}

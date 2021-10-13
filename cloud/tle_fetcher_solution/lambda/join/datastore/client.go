package datastore

import (
	"fmt"
	"time"
	"tle_fetcher_solution/shared/db"
	"tle_fetcher_solution/shared/db/model"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const (
	connection_table_name = "Connection"
	satellite_table_name  = "Satellite"
)

func (d *datastore) RegisterConnection(connectionID, appid, stage string, satellites []string) error {
	err := d.registerSatellites(satellites)
	if err != nil {
		return err
	}

	err = d.registerClient(connectionID, appid, stage, satellites)
	if err != nil {
		return err
	}
	return nil
}

func (d *datastore) registerClient(connectionID string, appid string, stage string, satellites []string) error {
	connection := model.DBConnectedClient{
		ConnectionID: connectionID,
		Satellites:   satellites,
		APIID:        appid,
		Stage:        stage,
		Expire:       time.Now().Add(24 * time.Hour).Unix(),
	}

	av, err := dynamodbattribute.MarshalMap(&connection)
	if err != nil {
		return fmt.Errorf("can't marshal:%w", err)
	}
	i := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(connection_table_name),
	}
	_, err = d.client.PutItem(i)
	if err != nil {
		return fmt.Errorf("can't insert:%w", err)
	}
	return nil
}

func (d *datastore) registerSatellites(satellites []string) error {
	for _, satellite := range satellites {
		satdb := model.Satellite{
			SatelliteName: satellite,
		}
		av, err := dynamodbattribute.MarshalMap(&satdb)
		if err != nil {
			return fmt.Errorf("can't marshal:%w", err)
		}
		i := &dynamodb.PutItemInput{
			Item:                av,
			TableName:           aws.String(satellite_table_name),
			ConditionExpression: aws.String("attribute_not_exists(SatelliteName)"),
		}
		_, err = d.client.PutItem(i)
		err = db.EscapeNoRowError(err)
		if err != nil {
			return fmt.Errorf("cant insert satellite: %w", err)
		}
	}
	return nil
}

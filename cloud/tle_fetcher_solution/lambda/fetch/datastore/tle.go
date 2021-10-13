package datastore

import (
	"fmt"
	"log"
	"strings"
	"time"
	"tle_fetcher_solution/lambda/fetch/model"
	"tle_fetcher_solution/shared/db"
	db_model "tle_fetcher_solution/shared/db/model"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const (
	tles_table_name = "Tle"
)

func (d *datastore) InsertTle(tles map[string]*model.Tle) error {
	for _, tle := range tles {
		satTle := db_model.SatelliteTLE{
			SatelliteLastID: fmt.Sprintf("%s_%d", strings.Trim(tle.TitleLine.Satname, " "), tle.LineOne.ElementSetNumber),
			TLE:             tle.Text,
			Satellite:       strings.Trim(tle.TitleLine.Satname, " "),
			Expire:          time.Now().Add(24 * 7 * time.Hour).Unix(),
		}
		av, err := dynamodbattribute.MarshalMap(satTle)
		if err != nil {
			log.Fatalf("Got error marshalling map: %s", err)
		}
		i := &dynamodb.PutItemInput{
			Item:                av,
			TableName:           aws.String(tles_table_name),
			ConditionExpression: aws.String("attribute_not_exists(SatelliteLastID)"),
		}
		_, err = d.client.PutItem(i)
		err = db.EscapeNoRowError(err)
		if err != nil {
			return fmt.Errorf("cant insert tle: %w", err)
		}
	}
	return nil
}

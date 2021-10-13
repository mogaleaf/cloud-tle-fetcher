package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"tle_fetcher_solution/lambda/receive/datastore"
	"tle_fetcher_solution/lambda/receive/model"
	"tle_fetcher_solution/lambda/receive/notify"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(e events.DynamoDBEvent) error {
	db, err := datastore.NewDatastore(os.Getenv("REGION"))
	if err != nil {
		return err
	}
	gw, err := notify.NewGateway(os.Getenv("REGION"))
	if err != nil {
		return err
	}
	for _, record := range e.Records {
		if record.EventName == "INSERT" {
			log.Printf("processing request data, event=%s, type=%s", record.EventID, record.EventName)
			tle := &model.SatelliteTLE{}
			sat, ok := record.Change.NewImage["Satellite"]
			if !ok {
				continue
			}
			tle.Satellite = sat.String()
			tleS, ok := record.Change.NewImage["TLE"]
			if !ok {
				continue
			}
			tle.TLE = tleS.String()
			connections, err := db.GetRegisteredConnection(tle.Satellite)
			if err != nil {
				return err
			}
			for _, c := range connections {
				err := gw.PostTle(c, tle)
				if err != nil {
					log.Printf(fmt.Sprintf("ERROR: %v", err))
				}
			}
		}
	}
	return nil
}

func main() {
	aws, _ := os.LookupEnv("AWS_LAMBDA")
	if aws != "" {
		lambda.Start(handler)
	} else {
		os.Setenv("REGION", "us-east-1")
		e := events.DynamoDBEvent{}
		err := json.Unmarshal([]byte(test), &e)
		if err != nil {
			println(err.Error())
		}
		err = handler(e)
		if err != nil {
			println(err.Error())
		}
	}
}

const test = `
	{
	"Records": [
	{
	"eventID": "c4ca4238a0b923820dcc509a6f75849b",
	"eventName": "INSERT",
	"eventVersion": "1.1",
	"eventSource": "aws:dynamodb",
	"awsRegion": "us-east-1",
	"dynamodb": {
	"Keys": {
	"Id": {
	"N": "101"
	}
	},
	"NewImage": {
	"Message": {
	"S": "New item!"
	},
	"Satellite": {
	"S": "HIBER-3"
	},
	"TLE": {
	"S": "tle"
	}
	},
	"ApproximateCreationDateTime": 1428537600,
	"SequenceNumber": "4421584500000000017450439091",
	"SizeBytes": 26,
	"StreamViewType": "NEW_AND_OLD_IMAGES"
	},
	"eventSourceARN": "arn:aws:dynamodb:us-east-1:123456789012:table/ExampleTableWithStream/stream/2015-06-27T00:48:05.899"
	},
	{
	"eventID": "c81e728d9d4c2f636f067f89cc14862c",
	"eventName": "MODIFY",
	"eventVersion": "1.1",
	"eventSource": "aws:dynamodb",
	"awsRegion": "us-east-1",
	"dynamodb": {
	"Keys": {
	"Id": {
	"N": "101"
	}
	},
	"NewImage": {
	"Message": {
	"S": "This item has changed"
	},
	"Id": {
	"N": "101"
	}
	},
	"OldImage": {
	"Message": {
	"S": "New item!"
	},
	"Id": {
	"N": "101"
	}
	},
	"ApproximateCreationDateTime": 1428537600,
	"SequenceNumber": "4421584500000000017450439092",
	"SizeBytes": 59,
	"StreamViewType": "NEW_AND_OLD_IMAGES"
	},
	"eventSourceARN": "arn:aws:dynamodb:us-east-1:123456789012:table/ExampleTableWithStream/stream/2015-06-27T00:48:05.899"
	},
	{
	"eventID": "eccbc87e4b5ce2fe28308fd9f2a7baf3",
	"eventName": "REMOVE",
	"eventVersion": "1.1",
	"eventSource": "aws:dynamodb",
	"awsRegion": "us-east-1",
	"dynamodb": {
	"Keys": {
	"Id": {
	"N": "101"
	}
	},
	"OldImage": {
	"Message": {
	"S": "This item has changed"
	},
	"Id": {
	"N": "101"
	}
	},
	"ApproximateCreationDateTime": 1428537600,
	"SequenceNumber": "4421584500000000017450439093",
	"SizeBytes": 38,
	"StreamViewType": "NEW_AND_OLD_IMAGES"
	},
	"eventSourceARN": "arn:aws:dynamodb:us-east-1:123456789012:table/ExampleTableWithStream/stream/2015-06-27T00:48:05.899"
	}
	]
	}`

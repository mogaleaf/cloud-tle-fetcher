package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"tle_fetcher_solution/lambda/join/datastore"
	"tle_fetcher_solution/shared/session"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(req events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println(fmt.Sprintf("Receiving request ConnectionID= %s", req.RequestContext.ConnectionID))
	requestSat := &session.JoinRequest{}
	err := json.Unmarshal([]byte(req.Body), requestSat)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, err
	}
	newDatastore, err := datastore.NewDatastore(os.Getenv("REGION"))
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, err
	}
	var names []string
	for _, sat := range requestSat.Satellites {
		names = append(names, sat.Name)
	}
	err = newDatastore.RegisterConnection(req.RequestContext.ConnectionID, req.RequestContext.APIID, req.RequestContext.Stage, names)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, err
	}
	return events.APIGatewayProxyResponse{StatusCode: http.StatusOK}, nil
}

func main() {
	aws, _ := os.LookupEnv("AWS_LAMBDA")
	if aws != "" {
		lambda.Start(handler)
	} else {
		//TODO
		req := &session.JoinRequest{
			Satellites: []session.Satellite{
				{Name: "HIBER-3"},
				{Name: "HIBER-4"},
			},
		}
		marshal, _ := json.Marshal(req)
		println(string(marshal))

		requestSat := &session.JoinRequest{}
		json.Unmarshal(marshal, requestSat)
		println(requestSat.Satellites[0].Name)

		newDatastore, _ := datastore.NewDatastore("us-east-1")
		newDatastore.RegisterConnection("test", "appid", "stage", []string{"CustomSat"})
		//println(err.Error())
		//can't insert:ConditionalCheckFailedException: The conditional request failed: wrapError
		//null
	}
}

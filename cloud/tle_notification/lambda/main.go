package main

import (
	"flag"
	"fmt"
	"os"
	"tle_manager/tle_notification/lambda/api"
	"tle_manager/tle_notification/lambda/datastore"
	event2 "tle_manager/tle_notification/lambda/event"

	"github.com/aws/aws-lambda-go/lambda"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
	aws, _ := os.LookupEnv("AWS_LAMBDA")
	if aws != "" {
		lambda.Start(RunMain)
	} else {
		os.Setenv("RDS_HOSTNAME", "datastore.ck4nzrl2hito.us-east-1.rds.amazonaws.com")
		os.Setenv("RDS_PORT", "5432")
		os.Setenv("RDS_USER", "...")
		os.Setenv("RDS_PASSWORD", "...")
		os.Setenv("RDS_SCHEMA", "datastore")
		os.Setenv("WS_ADDR", "localhost:8080")
		err := RunMain()
		if err != nil {
			println(fmt.Sprintf("%v", fmt.Errorf("problem: %w", err)))
		}
		//tle, err := api.FetchAndSaveTle()
		//if err != nil {
		//	println(fmt.Sprintf("%v",err))
		//} else{
		//	println(fmt.Sprintf("%v",tle))
		//}
	}

}

func RunMain() error {
	rdsHostname, _ := os.LookupEnv("RDS_HOSTNAME")
	rdsPort, _ := os.LookupEnv("RDS_PORT")
	rdsDatabase, _ := os.LookupEnv("RDS_SCHEMA")
	rdsUser, _ := os.LookupEnv("RDS_USER")
	rdsPassword, _ := os.LookupEnv("RDS_PASSWORD")
	addr, _ := os.LookupEnv("WS_ADDR")
	event := event2.NewSatelliteTleObserver()
	db, err := datastore.NewDatastore(rdsHostname, rdsUser, rdsPassword, rdsPort, rdsDatabase, event)
	if err != nil {
		return fmt.Errorf("can't connect datastore: %w", err)
	}

	print(db)
	go db.ListenNewTle()
	api.ServeClient(addr, event, db)

	return nil
}

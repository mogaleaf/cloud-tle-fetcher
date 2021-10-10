package main

import (
	"fmt"
	"os"
	"tle-fetcher/api"
	"tle-fetcher/datastore"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	aws, _ := os.LookupEnv("AWS_LAMBDA")
	if aws != "" {
		lambda.Start(RunMain)
	} else {
		os.Setenv("RDS_HOSTNAME", "datastore.ck4nzrl2hito.us-east-1.rds.amazonaws.com")
		os.Setenv("RDS_PORT", "5432")
		os.Setenv("RDS_USER", "....")
		os.Setenv("RDS_PASSWORD", "....")
		os.Setenv("RDS_SCHEMA", "datastore")
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
	db, err := datastore.Connect(rdsHostname, rdsUser, rdsPassword, rdsPort, rdsDatabase)
	if err != nil {
		return fmt.Errorf("can't connect datastore: %w", err)
	}

	tle, err := api.FetchAndSaveTle(db)
	//TODO
	println(tle)

	return err
}

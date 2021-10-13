package main

import (
	"fmt"
	"os"
	"tle_fetcher_solution/lambda/fetch/api"
	"tle_fetcher_solution/lambda/fetch/datastore"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	aws := os.Getenv("AWS_LAMBDA")
	if aws != "" {
		lambda.Start(RunMain)
	} else {
		os.Setenv("REGION", "us-east-1")
		err := RunMain()
		if err != nil {
			println(fmt.Sprintf("%v", fmt.Errorf("problem: %w", err)))
		}
	}

}

func RunMain() error {
	db, err := datastore.NewDatastore(os.Getenv("REGION"))
	if err != nil {
		return fmt.Errorf("can't connect datastore: %w", err)
	}

	err = api.FetchAndSaveTle(db)
	return err
}

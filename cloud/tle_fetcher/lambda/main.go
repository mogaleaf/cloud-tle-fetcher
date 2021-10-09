package main

import (
	"tle-fetcher/api"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(api.FetchAndSaveTle)
}

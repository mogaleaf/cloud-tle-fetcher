package main

import (
	"os"
	"tle-fetcher/api"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	aws, _ := os.LookupEnv("AWS_LAMBDA")
	if aws != "" {
		lambda.Start(api.FetchAndSaveTle)
	} else {
		//os.Setenv("RDS_HOSTNAME","datastore.ck4nzrl2hito.us-east-1.rds.amazonaws.com")
		//os.Setenv("RDS_PORT","5432")
		//os.Setenv("RDS_USER","....")
		//os.Setenv("RDS_PASSWORD",".....")
		//os.Setenv("RDS_SCHEMA","datastore")
		//tle, err := api.FetchAndSaveTle()
		//if err != nil {
		//	println(fmt.Sprintf("%v",err))
		//} else{
		//	println(fmt.Sprintf("%v",tle))
		//}
	}

}

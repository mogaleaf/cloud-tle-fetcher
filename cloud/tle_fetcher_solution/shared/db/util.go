package db

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func EscapeNoRowError(err error) error {
	if err == nil {
		return nil
	}
	if aerr, ok := err.(awserr.Error); ok {
		switch aerr.Code() {
		case dynamodb.ErrCodeConditionalCheckFailedException:
		default:
			return err
		}
	}
	return nil
}

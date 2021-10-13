package session

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

func NewSession(region string) *session.Session {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{Region: aws.String(region)},
	}))
	return sess
}

type JoinRequest struct {
	Satellites []Satellite `json:"satellites"`
}

type Satellite struct {
	Name string `json:"name"`
}

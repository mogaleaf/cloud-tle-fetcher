package notify

import (
	"fmt"
	internal "tle_fetcher_solution/lambda/receive/model"
	"tle_fetcher_solution/shared/db/model"
	"tle_fetcher_solution/shared/session"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/apigatewaymanagementapi"
)

type gateway struct {
	region string
}

type Gateway interface {
	PostTle(client model.DBConnectedClient, tle *internal.SatelliteTLE) error
}

func (g *gateway) newClient(appid, stage string) *apigatewaymanagementapi.ApiGatewayManagementApi {
	return apigatewaymanagementapi.New(session.NewSession(g.region), aws.NewConfig().WithEndpoint(fmt.Sprintf("%s.execute-api.us-east-1.amazonaws.com/%s", appid, stage)).WithRegion(g.region))
}

func NewGateway(region string) (Gateway, error) {
	g := &gateway{
		region: region,
	}
	return g, nil
}

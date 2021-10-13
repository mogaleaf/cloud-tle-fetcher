package notify

import (
	"encoding/json"
	"fmt"
	internal "tle_fetcher_solution/lambda/receive/model"
	"tle_fetcher_solution/shared/db/model"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/apigatewaymanagementapi"
)

func (p *gateway) PostTle(client model.DBConnectedClient, tle *internal.SatelliteTLE) error {
	marshal, err := json.Marshal(tle)
	if err != nil {
		return fmt.Errorf("cant marshal: %w", err)
	}
	_, err = p.newClient(client.APIID, client.Stage).PostToConnection(&apigatewaymanagementapi.PostToConnectionInput{
		ConnectionId: aws.String(client.ConnectionID),
		Data:         marshal,
	})
	if err != nil {
		return fmt.Errorf("cant send: %w", err)
	}
	return nil
}

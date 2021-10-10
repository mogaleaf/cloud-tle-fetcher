package datastore

import (
	"tle_manager/shared/db"
	"tle_manager/tle_notification/lambda/event"

	"github.com/go-pg/pg/v10"
)

type datastore struct {
	*pg.DB
	observer event.SatelliteTleObserver
}
type Datastore interface {
	SaveSatellites(names []string) error
	ListenNewTle() error
}

func NewDatastore(rdsHostname, rdsUser, rdsPassword, rdsPort, rdsDatabase string, observer event.SatelliteTleObserver) (Datastore, error) {
	connect, err := db.Connect(rdsHostname, rdsUser, rdsPassword, rdsPort, rdsDatabase)
	if err != nil {
		return nil, err
	}
	return &datastore{
		DB:       connect,
		observer: observer,
	}, nil
}

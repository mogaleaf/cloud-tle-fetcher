package datastore

import (
	"tle_manager/shared/db"
	"tle_manager/tle_fetcher/lambda/model"

	"github.com/go-pg/pg/v10"
)

type datastore struct {
	*pg.DB
}

type Datastore interface {
	InsertTle(tles map[string]*model.Tle) error
	GetSatelliteNames() ([]string, error)
}

func NewDatastore(rdsHostname, rdsUser, rdsPassword, rdsPort, rdsDatabase string) (Datastore, error) {
	connect, err := db.Connect(rdsHostname, rdsUser, rdsPassword, rdsPort, rdsDatabase)
	if err != nil {
		return nil, err
	}
	return &datastore{
		DB: connect,
	}, nil
}

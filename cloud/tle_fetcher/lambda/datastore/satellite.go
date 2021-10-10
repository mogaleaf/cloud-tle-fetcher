package datastore

import (
	"fmt"
	"tle_manager/shared/db/model"
)

func (db *datastore) GetSatelliteNames() ([]string, error) {
	var names []string
	if err := db.Model(&model.Satellite{}).
		Column("name").
		Select(&names); err != nil {
		return nil, fmt.Errorf("can't select satellites: %w", err)
	}
	return names, nil
}

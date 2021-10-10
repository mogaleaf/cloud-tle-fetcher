package datastore

import (
	"tle_manager/shared/db/model"
)

func (db *datastore) getLastTle(satname string) (*model.Tle, error) {
	tle := &model.Tle{}
	if err := db.Model(tle).
		Join("JOIN satellite on satellite.id = tle.satellite_id").
		Where("satellite.name=?", satname).
		Order("tle.last_id").
		Limit(1).
		Select(); err != nil {
		return nil, err
	}
	return tle, nil
}

package datastore

import (
	"fmt"
	"tle_manager/shared/db/model"
)

func (db *datastore) SaveSatellites(names []string) error {
	var sats []*model.Satellite
	for _, name := range names {
		satName := name
		sats = append(sats, &model.Satellite{
			Name: &satName,
		})
	}
	_, err := db.Model(&sats).OnConflict("do nothing").Insert()
	if err != nil {
		return fmt.Errorf("can't insert sats:%w", err)
	}
	return nil
}

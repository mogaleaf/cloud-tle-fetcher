package datastore

import (
	"fmt"
	"tle-fetcher/model"

	"github.com/go-pg/pg/v10"
)

func (db *datastore) InsertTle(tles map[string]*model.Tle) error {
	var dbTles []*Tle
	var names []string
	for satName, _ := range tles {
		names = append(names, satName)
	}
	var satellites []*Satellite
	if err := db.Model(&satellites).
		Where("satellite.name in (?)", pg.In(names)).
		Select(); err != nil {
		return fmt.Errorf("can't select satellites: %w", err)
	}
	mapSatByName := make(map[string]*Satellite)
	for _, sat := range satellites {
		mapSatByName[*sat.Name] = sat
	}

	for satName, tle := range tles {
		name := satName
		satellite, ok := mapSatByName[name]
		if !ok {
			continue
		}
		dbTles = append(dbTles, &Tle{
			SatelliteId: satellite.Id,
			Line1:       &tle.Lines[0],
			Line2:       &tle.Lines[1],
			Line3:       &tle.Lines[2],
		})
	}

	_, err := db.Model(&dbTles).Insert()
	if err != nil {
		return fmt.Errorf("can't save tle: %w", err)
	}
	return nil
}

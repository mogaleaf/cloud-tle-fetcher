package api

import (
	"tle_fetcher_solution/lambda/fetch/celestrak_fetcher"
	"tle_fetcher_solution/lambda/fetch/datastore"
)

func FetchAndSaveTle(db datastore.Datastore) error {
	sats, err := db.GetSatelliteNames()
	if err != nil {
		return err
	}
	if len(sats) == 0 {
		return nil
	}
	tles, _ := celestrak_fetcher.FetchTLEForSatellites(sats)
	err = db.InsertTle(tles)
	if err != nil {
		return err
	}
	return nil
}

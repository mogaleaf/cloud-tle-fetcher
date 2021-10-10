package api

import (
	"fmt"
	"tle_manager/tle_fetcher/lambda/celestrak_fetcher"
	"tle_manager/tle_fetcher/lambda/datastore"
)

func FetchAndSaveTle(db datastore.Datastore) (string, error) {
	//TODO from db get sat names

	sats, err := db.GetSatelliteNames()
	if err != nil {
		return "", err
	}
	if len(sats) == 0 {
		return "", nil
	}
	tles, _ := celestrak_fetcher.FetchTLEForSatellites(sats)
	//TODO save those tles
	fmt.Sprintf("%v", tles)

	err = db.InsertTle(tles)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v", tles), nil
}

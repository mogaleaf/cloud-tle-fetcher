package api

import (
	"fmt"
	"tle-fetcher/celestrak_fetcher"
	"tle-fetcher/datastore"
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

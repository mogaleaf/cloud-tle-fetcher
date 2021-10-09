package api

import (
	"fmt"
	"tle-fetcher/celestrak_fetcher"
	"tle-fetcher/datastore"
)

func FetchAndSaveTle() (string, error) {
	//TODO from db get sat names

	sats := []string{"HIBER-3", "HIBER-4"}
	tles, _ := celestrak_fetcher.FetchTLEForSatellites(sats)
	//TODO save those tles
	fmt.Sprintf("%v", tles)

	db, err := datastore.Connect()
	if err != nil {
		return "", err
	}
	err = datastore.InsertTle(db, tles)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v", tles), nil
}

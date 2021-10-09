package api

import (
	"fmt"
	"tle-fetcher/celestrak_fetcher"
)

func FetchAndSaveTle() (string, error) {
	//TODO from db get sat names

	sats := []string{"HIBER-3", "HIBER-4"}
	tles, _ := celestrak_fetcher.FetchTLEForSatellites(sats)
	//TODO save those tles
	fmt.Sprintf("%v", tles)
	return fmt.Sprintf("%v", tles), nil
}

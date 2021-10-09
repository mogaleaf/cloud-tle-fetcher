package celestrak_fetcher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchTLEForSatellite(t *testing.T) {
	tle, err := fetchTLEForSatellite("HIBER-4")
	assert.Nil(t, err)
	assert.NotEmpty(t, tle.Lines)
}

func TestFetchTLEForSatellites(t *testing.T) {
	tles, errs := FetchTLEForSatellites([]string{"HIBER-4", "HIBER-3", "Cecile-sat"})
	assert.Len(t, tles, 2)
	assert.Len(t, errs, 1)
}

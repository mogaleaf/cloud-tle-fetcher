package model

type SatelliteTLE struct {
	SatelliteLastID string `json:"-"`
	Satellite       string `json:"Satellite"`
	TLE             string `json:"Tle"`
	Expire          int64  `json:"-"`
}

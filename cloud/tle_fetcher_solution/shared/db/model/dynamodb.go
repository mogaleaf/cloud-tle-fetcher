package model

type DBConnectedClient struct {
	ConnectionID string
	APIID        string
	Stage        string
	Satellites   []string
	Expire       int64
}

type Satellite struct {
	SatelliteName string
}

type SatelliteTLE struct {
	SatelliteLastID string
	Satellite       string
	TLE             string
	Expire          int64
}

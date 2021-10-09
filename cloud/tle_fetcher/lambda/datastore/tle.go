package datastore

import (
	"tle-fetcher/model"

	"github.com/go-pg/pg/v10"
)

type Tle struct {
	tableName struct{} `pg:"tle,discard_unknown_columns"`

	SatelliteName *string `pg:"satellite_name"`
	Line1         *string `pg:"line1"`
	Line2         *string `pg:"line2"`
	Line3         *string `pg:"line3"`
}

func InsertTle(db *pg.DB, tles map[string]*model.Tle) error {
	var dbTles []*Tle
	for satName, tle := range tles {
		name := satName
		dbTles = append(dbTles, &Tle{
			SatelliteName: &name,
			Line1:         &tle.Lines[0],
			Line2:         &tle.Lines[1],
			Line3:         &tle.Lines[2],
		})
	}

	_, err := db.Model(&dbTles).Insert()
	return err
}

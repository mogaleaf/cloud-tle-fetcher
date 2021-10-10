package model

type Satellite struct {
	tableName struct{} `pg:"satellite,discard_unknown_columns"`

	Id   *uint32 `pg:"id,pk"`
	Name *string `pg:"name,unique"`
}

type Tle struct {
	tableName struct{} `pg:"tle,discard_unknown_columns"`

	Satellite   *Satellite `pg:"satellite_id,rel:belongs-to,join_fk:id"`
	SatelliteId *uint32    `pg:"satellite_id"`
	Line1       *string    `pg:"line1"`
	Line2       *string    `pg:"line2"`
	Line3       *string    `pg:"line3"`
	LastID      *uint32    `pg:"last_id"`
}

package datastore

import (
	"context"
	"text/template"
	"tle-fetcher/model"
	"tle-fetcher/util"

	"github.com/go-pg/pg/v10"
)

type datastore struct {
	*pg.DB
}

type Datastore interface {
	InsertTle(tles map[string]*model.Tle) error
	GetSatelliteNames() ([]string, error)
}

const (
	templatePostgresName = "postgres_template"
)

var (
	//TODO should be ssl
	templatePostgres = template.Must(template.New(templatePostgresName).Parse("postgres://{{ .User }}:{{ .Password }}@{{ .Host }}:{{ .Port }}/{{ .Db }}?sslmode=disable"))
)

func Connect(rdsHostname, rdsUser, rdsPassword, rdsPort, rdsDatabase string) (Datastore, error) {
	dbUrl, err := util.ExecTempl(templatePostgres, struct {
		User     string
		Password string
		Host     string
		Db       string
		Port     string
	}{
		User:     rdsUser,
		Password: rdsPassword,
		Host:     rdsHostname,
		Db:       rdsDatabase,
		Port:     rdsPort,
	})
	if err != nil {
		return nil, err
	}
	opt, err := pg.ParseURL(dbUrl)
	if err != nil {
		return nil, err
	}
	db := pg.Connect(opt)
	ctx := context.Background()

	if err := db.Ping(ctx); err != nil {
		return nil, err
	}

	d := datastore{
		db,
	}

	if err := d.createSchema(); err != nil {
		return nil, err
	}
	return &d, nil
}

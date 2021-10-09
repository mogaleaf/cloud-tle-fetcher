package datastore

import (
	"context"
	"os"
	"text/template"
	"tle-fetcher/util"

	"github.com/go-pg/pg/v10"
)

const (
	templatePostgresName = "postgres_template"
)

var (
	//TODO should be ssl
	templatePostgres = template.Must(template.New(templatePostgresName).Parse("postgres://{{ .User }}:{{ .Password }}@{{ .Host }}:{{ .Port }}/{{ .Db }}?sslmode=disable"))
)

func Connect() (*pg.DB, error) {
	rdsHostname, _ := os.LookupEnv("RDS_HOSTNAME")
	rdsPort, _ := os.LookupEnv("RDS_PORT")
	rdsDatabase, _ := os.LookupEnv("RDS_SCHEMA")
	rdsUser, _ := os.LookupEnv("RDS_USER")
	rdsPassword, _ := os.LookupEnv("RDS_PASSWORD")

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

	return db, nil
}

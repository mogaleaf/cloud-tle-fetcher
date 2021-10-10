package db

import (
	"context"
	"fmt"
	"text/template"
	"tle_manager/shared/db/model"
	"tle_manager/shared/util"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

const (
	templatePostgresName = "postgres_template"
)

var (
	//TODO should be ssl
	templatePostgres = template.Must(template.New(templatePostgresName).Parse("postgres://{{ .User }}:{{ .Password }}@{{ .Host }}:{{ .Port }}/{{ .Db }}?sslmode=disable"))
)

func Connect(rdsHostname, rdsUser, rdsPassword, rdsPort, rdsDatabase string) (*pg.DB, error) {
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

	if err := createSchema(db); err != nil {
		return nil, err
	}
	return db, nil
}

func createSchema(db *pg.DB) error {
	if err := db.Model(&model.Satellite{}).CreateTable(&orm.CreateTableOptions{
		IfNotExists:   true,
		FKConstraints: true,
	}); err != nil {
		return fmt.Errorf("can't create schema: %w", err)
	}
	if err := db.Model(&model.Tle{}).CreateTable(&orm.CreateTableOptions{
		IfNotExists:   true,
		FKConstraints: true,
	}); err != nil {
		return fmt.Errorf("can't create schema: %w", err)
	}
	return nil

}

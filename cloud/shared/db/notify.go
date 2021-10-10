package db

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

func Notify(db *pg.DB, channelName, message string) (orm.Result, error) {
	return db.Exec(fmt.Sprintf("NOTIFY %s,?", channelName), message)
}

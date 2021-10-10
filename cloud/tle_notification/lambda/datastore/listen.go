package datastore

import (
	"context"
	"fmt"
)

func (db *datastore) ListenNewTle() error {
	//TODO
	ln := db.Listen(context.Background(), "new_tle")
	defer ln.Close()

	ch := ln.Channel()
	for val := range ch {
		tle, _ := db.getLastTle(val.Payload)
		db.observer.NewTle(val.Payload, fmt.Sprintf("%s\n%s\n%s", *tle.Line1, *tle.Line2, *tle.Line3))
	}
	return nil
}

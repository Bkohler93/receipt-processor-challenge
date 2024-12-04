package main

import "database/sql"

type config struct {
	store store
}

func newConfig(database *sql.DB) config {
	store := newStore(database)
	return config{store: store}
}

package storage

import (
	"github.com/jehaby/webapp102/config"

	"github.com/jmoiron/sqlx"
)

func getTestDB() *sqlx.DB {
	c := config.DB{Conn: "user=postgres dbname=webapp port=65432 host=localhost sslmode=disable"} // TODO: test config
	db, err := NewDB(c)
	if err != nil {
		panic(err)
	}
	return db
}

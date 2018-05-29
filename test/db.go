package test

import (
	"database/sql"
	"io/ioutil"
	"log"

	"github.com/go-pg/pg"
	"github.com/jehaby/webapp102/config"
	"github.com/jehaby/webapp102/storage"
)

func getDBs(cfg config.DB) (*sql.DB, *pg.DB) {
	var (
		sqldb *sql.DB
		pgdb  *pg.DB
		err   error
	)

	sqldb, err = sql.Open("postgres", cfg.URL)
	if err != nil {
		log.Fatalf("couldn't open db: %v", err)
	}

	pgdb, err = storage.NewPGDB(cfg)
	if err != nil {
		log.Fatal("couldn't get pgdb", err)
	}

	return sqldb, pgdb

}

type db struct {
	*pg.DB
}

func (db *db) exec(query string) {
	if _, err := db.Exec(query); err != nil {
		log.Printf("got error from query: '%s' | err: %v", query, err)
	}
}

const clearQuery = `
DELETE FROM ads;
DELETE FROM products;
DELETE FROM users;
DELETE FROM brands;
DELETE FROM categories;
`

func seedQuery() string {
	res, err := ioutil.ReadFile("data/data.sql")
	if err != nil {
		log.Fatal(err)
	}
	return string(res)
}

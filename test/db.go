package test

import (
	"database/sql"
	"io/ioutil"
	"log"

	"github.com/go-pg/pg"
	"github.com/jehaby/webapp102/storage"
)

var (
	sqldb *sql.DB
	pgdb  *pg.DB
)

func init() {
	var err error
	sqldb, err = sql.Open("postgres", "postgres://postgres@localhost:65432/webapp?sslmode=disable")
	if err != nil {
		log.Fatalf("couldn't open db: %v", err)
	}

	conf := getConf()
	pgdb, err = storage.NewPGDB(conf.DB)
	if err != nil {
		log.Fatal("couldn't get pgdb", err)
	}
}

func GetPGDB() *pg.DB {
	if pgdb == nil {
		log.Fatal("dbi is nil")
	}
	return pgdb
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

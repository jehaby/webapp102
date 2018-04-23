package test

import (
	"io/ioutil"
	"log"

	"github.com/go-pg/pg"
)

type db struct {
	*pg.DB
}

func (db *db) exec(query string) {
	if _, err := db.Exec(query); err != nil {
		log.Panicf("couldn't execute query: '%s' | err: %v", query, err)
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

package test

import (
	"log"
	"os"
	"testing"

	"github.com/jehaby/webapp102/storage"

	"github.com/jehaby/webapp102/config"
)

func TestMain(m *testing.M) {

	log.Println("Started tests")
	var err error

	conf := getConf()

	pgdb, err := storage.NewPGDB(conf.DB)
	if err != nil {
		log.Fatal("couldn't get pgdb", err)
	}

	db := &db{pgdb}

	db.exec(clearQuery)
	db.exec(seedQuery())

	// TODO: call flag.Parse() here if TestMain uses flags
	os.Exit(func() int {
		defer func() {
			db.Close()
			log.Println("Finished tests")
		}()
		return m.Run()
	}())
}

func getConf() config.C {
	return config.C{
		DB: config.DB{
			User:     "postgres",
			Database: "webapp",
			Port:     "65432",
			Host:     "localhost",
		},
	}
}

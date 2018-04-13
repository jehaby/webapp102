package test

import (
	"log"
	"testing"

	"github.com/satori/go.uuid"

	"github.com/jehaby/webapp102/config"
	"github.com/jehaby/webapp102/storage"
)

func TestAdRepo_GetByUUID(t *testing.T) {

	dbConf := config.DB{
		Conn:     "user=postgres dbname=webapp port=65432 host=localhost sslmode=disable",
		User:     "postgres",
		Database: "webapp",
		Port:     "65432",
		Host:     "localhost",
	}

	db, err := storage.NewDB(dbConf)
	if err != nil {
		log.Panicf("couldn't open db: %v", err)
	}

	pgDB, err := storage.NewPGDB(dbConf)
	if err != nil {
		log.Fatal("error creating pg: ", err)
	}

	repo := storage.NewAdRepository(db, pgDB)

	res, err := repo.GetByUUID(uuid.FromStringOrNil("5df5b126-1fac-4fe1-a421-972ba56eb17b"))
	if err != nil {
		t.Fatal(err)
	}

	if res == nil {
		t.Fatalf("nil res")
	}
}

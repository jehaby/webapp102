package test

import (
	"log"
	"os"
	"testing"

	"github.com/jehaby/webapp102/test/data"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
)

func TestMain(m *testing.M) {
	log.Println("Started tests")

	db := &db{GetPGDB()}

	driver, err := postgres.WithInstance(sqldb, &postgres.Config{})
	mig, err := migrate.NewWithDatabaseInstance(
		"file://./../var/migrations",
		"postgres", driver)
	if err != nil {
		log.Fatal("migrate.New returned error", err)
	}

	if err = mig.Drop(); err != nil {
		log.Fatalf("mig.Down returned error: %v", err)
	}

	// TODO: PR to mattes migrate (mig.Drop doesn't do it); or use some other migration system

	db.exec("DROP TYPE USER_ROLE;")
	db.exec("DROP TYPE CURRENCY;")
	db.exec("DROP TYPE CONDITION;")

	if err = mig.Up(); err != nil {
		log.Fatalf("mig.Up returned error: %v", err)
	}

	db.exec(seedQuery())
	_, err = db.Model(&data.Ads).Insert()
	if err != nil {
		log.Panic(err)
	}

	// TODO: call flag.Parse() here if TestMain uses flags
	os.Exit(func() int {
		defer func() {
			db.Close()
			mig.Close()
			log.Println("Finished tests")
		}()
		return m.Run()
	}())
}

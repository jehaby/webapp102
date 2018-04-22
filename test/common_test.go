package test

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/jehaby/webapp102/storage"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"

	"github.com/jehaby/webapp102/config"
)

func TestMain(m *testing.M) {
	db2, err := sql.Open("postgres", "postgres://postgres@localhost:65432/webapp?sslmode=disable")
	if err != nil {
		log.Fatalf("couldn't open db: %v", err)
	}

	log.Println("Started tests")

	conf := getConf()

	pgdb, err := storage.NewPGDB(conf.DB)
	if err != nil {
		log.Fatal("couldn't get pgdb", err)
	}

	db := &db{pgdb}

	driver, err := postgres.WithInstance(db2, &postgres.Config{})
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
	_, err = db.Exec("DROP TYPE USER_ROLE;")
	if err != nil {
		log.Printf("error in: DROP TYPE USER_ROLE: %v", err)
	}
	_, err = db.Exec("DROP TYPE CURRENCY;")
	if err != nil {
		log.Printf("error in: DROP TYPE CURRENCY: %v", err)
	}

	if err = mig.Up(); err != nil {
		log.Fatalf("mig.Up returned error: %v", err)
	}

	db.exec(seedQuery())

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

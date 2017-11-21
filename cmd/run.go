package main

import (
	"context"

	"log"

	"github.com/jehaby/webapp102/config"
	"github.com/jehaby/webapp102/http"
	"github.com/jehaby/webapp102/service"
	"github.com/jehaby/webapp102/storage"
)

func main() {
	cfg := config.C{
		config.HTTP{Addr: ":8899"}, // TODO: config
		config.DB{Conn: "user=postgres dbname=webapp port=65432 host=localhost sslmode=disable"},
	}

	db, err := storage.NewDB(cfg)
	if err != nil {
		log.Panicf("couldn't open db: %v", err)
	}

	ur := storage.NewUserRepository(db)

	app := service.NewApp(cfg, ur)
	defer app.Logger.Sync()

	httpApp := http.NewApp(
		cfg,
		app,
	)

	httpApp.Start(context.TODO())

}

package main

import (
	"context"

	"github.com/jehaby/webapp102/config"
	"github.com/jehaby/webapp102/http"
	"github.com/jehaby/webapp102/service"
)

func main() {
	cfg := config.C{
		config.HTTP{
			Addr:   ":8899",
			Secret: "secret",
		}, // TODO: config
		config.DB{
			Conn:     "user=postgres dbname=webapp port=65432 host=localhost sslmode=disable",
			User:     "postgres",
			Database: "webapp",
			Port:     "65432",
			Host:     "localhost",
		},
	}

	app := service.NewApp(cfg)
	defer app.Logger.Sync()

	httpApp := http.NewApp(
		cfg,
		app,
	)

	httpApp.Start(context.TODO())
}

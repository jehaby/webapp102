package main

import (
	"context"
	"log"

	"github.com/jehaby/webapp102/config"
	"github.com/jehaby/webapp102/http"
	"github.com/jehaby/webapp102/service"
)

func main() {
	cfg, err := config.Get("dev")
	if err != nil {
		log.Fatalf("couldn't load condif: %v", err)
	}

	app := service.NewApp(cfg)
	defer app.Logger.Sync()

	httpApp := http.NewApp(
		cfg,
		app,
	)

	httpApp.Start(context.TODO())
}

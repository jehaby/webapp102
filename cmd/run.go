package main

import (
	"context"

	"log"

	"github.com/jehaby/webapp102/config"
	"github.com/jehaby/webapp102/http"
	"github.com/jehaby/webapp102/service"
	"github.com/jehaby/webapp102/storage"
)

//func handler(w http.ResponseWriter, r *http.Request) {
//w.Header().Set("Access-Control-Allow-Origin", "*")
//w.Write([]byte("all good"))
//}

func main() {
	//http.HandleFunc("/", handler)
	//http.ListenAndServe(":8899", nil)

	cfg := config.C{
		config.HTTP{Addr: ":8899"},
		config.DB{Conn: "user=postgres dbname=webapp port=65432 sslmode=disable"},
	}

	storage.NewMemory() // TODO: remove

	db, err := storage.NewDB(cfg)
	if err != nil {
		log.Panicf("couldn't open db: %v", err)
	}

	ur := storage.NewUserRepository(db)

	app := service.NewApp(cfg, ur)

	httpApp := http.NewApp(
		cfg,
		app,
	)

	httpApp.Start(context.TODO())

}

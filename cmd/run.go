package main

import (
	"context"

	"github.com/jehaby/webapp102/config"
	"github.com/jehaby/webapp102/http"
)

//func handler(w http.ResponseWriter, r *http.Request) {
//w.Header().Set("Access-Control-Allow-Origin", "*")
//w.Write([]byte("all good"))
//}

func main() {
	//http.HandleFunc("/", handler)
	//http.ListenAndServe(":8899", nil)

	app := http.NewApp(context.TODO(), config.C{
		config.HTTP{Addr: ":8899"},
	})

	app.Start()

}

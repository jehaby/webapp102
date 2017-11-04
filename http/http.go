package http

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"

	"io/ioutil"

	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/jehaby/webapp102/config"
)

type app struct {
	cfg config.C
}

func NewApp(ctx context.Context, c config.C) *app {
	return &app{c}
}

func (a *app) Start() {

	//	jwtauth.Authenticator()

	tokenAuth := jwtauth.New("HS256", []byte("secret"), nil)

	// For debugging/example purposes, we generate and print
	// a sample jwt token with claims `user_id:123` here:
	_, tokenString, _ := tokenAuth.Encode(jwtauth.Claims{"user_id": 123})
	fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)

	r := chi.NewRouter()

	r.Use(jwtauth.Verifier(tokenAuth))

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(30 * time.Second))

	r.Post("/auth", a.authHandler)
	r.Post("/register", a.authHandler)

	log.Fatal(http.ListenAndServe(a.cfg.HTTP.Addr, r))
}

func (a *app) authHandler(w http.ResponseWriter, r *http.Request) {
	spew.Dump(ioutil.ReadAll(r.Body))
}

func (a *app) registerHandler(w http.ResponseWriter, r *http.Request) {

}

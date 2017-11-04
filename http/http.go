package http

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"

	"time"

	"github.com/jehaby/webapp102/config"
	"github.com/jehaby/webapp102/storage"
)

type app struct {
	cfg config.C

	storage *storage.Memory
}

func NewApp(c config.C) *app {
	return &app{c, storage.New()}
}

func (a *app) Start(ctx context.Context) {

	tokenAuth := jwtauth.New("HS256", []byte("secret"), nil)



	// For debugging/example purposes, we generate and print
	// a sample jwt token with claims `user_id:123` here:
	_, tokenString, _ := tokenAuth.Encode(jwtauth.Claims{
		"login": "urf",
		"password": "111",
	})
	fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)

	r := chi.NewRouter()

	r.Use(jwtauth.Verifier(tokenAuth))

//	r.Use(jwtauth.Authenticator)

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(30 * time.Second))

	r.Post("/login", a.loginHandler)
	r.Post("/register", a.loginHandler)

	log.Fatal(http.ListenAndServe(a.cfg.HTTP.Addr, r))
}

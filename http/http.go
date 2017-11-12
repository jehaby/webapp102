package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth"

	"github.com/jehaby/webapp102/config"
	"github.com/jehaby/webapp102/service"
	"go.uber.org/zap"
	"gopkg.in/go-playground/validator.v9"
)

var tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)

type app struct {
	cfg config.C

	validator *validator.Validate

	app *service.App
}

func NewApp(c config.C, a *service.App) *app {
	return &app{
		cfg:       c,
		validator: validator.New(),
		app:       a,
	}
}

func (a *app) Start(ctx context.Context) {

	auth := a.authRouter()

	//	mr := a.baseRouter()
	//	mr.Use(jwtauth.Authenticator)

	b := a.baseRouter()

	b.Mount("/api/v0/auth", auth)
	//	b.Mount("/api/v0", mr)

	a.log().Fatal(http.ListenAndServe(a.cfg.HTTP.Addr, b))
}

func (a *app) baseRouter() chi.Router {
	r := chi.NewRouter()

	claims := jwtauth.Claims{
		"login":    "urf",
		"password": "111",
	}

	// For debugging/example purposes, we generate and print
	// a sample jwt token with claims `user_id:123` here:
	_, tokenString, _ := tokenAuth.Encode(claims)
	fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)

	secret := "sss"
	_, tokenString, _ = jwtauth.New("HS256", []byte(secret), nil).Encode(claims)
	fmt.Printf("DEBUG: secret: '%s' a sample jwt is %s\n\n", secret, tokenString)

	secret = "www"
	_, tokenString, _ = jwtauth.New("HS256", []byte("secret"), nil).Encode(claims)
	fmt.Printf("DEBUG: secret: '%s' a sample jwt is %s\n\n", secret, tokenString)

	//	r.Use(jwtauth.Verifier(tokenAuth))
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(a.getCorsMiddleware())

	r.Use(middleware.Timeout(30 * time.Second))

	return r
}

func (a *app) getCorsMiddleware() func(http.Handler) http.Handler {
	cors := cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	return cors.Handler
}

func (a *app) authRouter() chi.Router {
	r := chi.NewRouter()
	r.Post("/login/", a.loginHandler)
	r.Post("/register/", a.registerHandler)
	r.Post("/reset/", a.resetPasswordHandler)
	r.Get("/logout/", a.logoutHandler)
	// TODO: reset password
	//
	return r
}

func (a *app) log() *zap.SugaredLogger {
	return a.app.Logger
}

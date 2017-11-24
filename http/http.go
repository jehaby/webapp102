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

type app struct {
	cfg config.C

	validator *validator.Validate

	app *service.App

	jwtAuth *jwtauth.JwtAuth
}

func NewApp(c config.C, a *service.App) *app {
	return &app{
		cfg:       c,
		validator: validator.New(),
		app:       a,
		jwtAuth:   jwtauth.New("HS256", []byte(c.HTTP.Secret), nil),
	}
}

func (a *app) Start(ctx context.Context) {
	a.log().Fatal(http.ListenAndServe(a.cfg.HTTP.Addr, a.getRoutes()))
}

func (a *app) baseRouter() chi.Router {
	r := chi.NewRouter()

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

func (a *app) log() *zap.SugaredLogger {
	return a.app.Logger
}

func withErr(l *zap.SugaredLogger, err error) *zap.SugaredLogger {
	// TODO: how to print it without the stack?
	return l.With("err", fmt.Sprintf("%s", err))
}

// contextKey is a value for use with context.WithValue. It's used as
// a pointer so it fits in an interface{} without allocation. This technique
// for defining context keys was copied from Go 1.7's new use of context in net/http.
type contextKey struct {
	name string
}

func (k *contextKey) String() string {
	return "jwtauth context value " + k.name
}

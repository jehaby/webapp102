package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	"go.uber.org/zap"
	"gopkg.in/go-playground/validator.v9"

	"github.com/jehaby/webapp102/config"
	"github.com/jehaby/webapp102/pkg/log"
	"github.com/jehaby/webapp102/service"
	"github.com/jehaby/webapp102/service/auth"
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
		jwtAuth:   jwtauth.New("HS256", []byte(c.Auth.Secret), nil),
	}
}

func (a *app) Start(ctx context.Context) {
	a.log().Infow("service started", "addr", a.cfg.HTTP.Addr)

	render.Respond = loggingRespond(a.log())

	a.log().Fatal(http.ListenAndServe(a.cfg.HTTP.Addr, a.getRoutes()))
}

func (a *app) baseRouter() chi.Router {
	r := chi.NewRouter()

	r.Use(
		a.authCookieMiddleware,
		middleware.Logger,
		middleware.Recoverer,
		a.getCorsMiddleware(),
		middleware.Timeout(30*time.Second),
	)

	return r
}

// authCookieMiddleware adds jwt token to context, if "auth" cookie is present;
func (a *app) authCookieMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, _ := r.Cookie(authCookieName)
		if c != nil {
			r = r.WithContext(
				context.WithValue(r.Context(), auth.StrTokenCtxKey, c.Value),
			)
		}
		spew.Dump("printing cookies", r.Cookies(), c)
		next.ServeHTTP(w, r)
	})
}

func (a *app) getCorsMiddleware() func(http.Handler) http.Handler {
	// TODO: prod settings
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8080"}, // Use this to allow specific origin hosts
		// AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	return cors.Handler
}

func (a *app) log() *log.Logger {
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

func loggingRespond(l *log.Logger) func(w http.ResponseWriter, r *http.Request, v interface{}) {
	return func(w http.ResponseWriter, r *http.Request, v interface{}) {
		switch err := v.(type) {
		case error:
			// TODO: should get here
			l.Errorw("responding with error", "err", err)
		case ErrResponse:
			l.Errorw(err.ErrorText, "err", err.Err, "http code", err.HTTPStatusCode)
		}
		render.DefaultResponder(w, r, v)
	}
}

package http

import (
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/jehaby/webapp102/resolver"
	"github.com/jehaby/webapp102/schema"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
)

func (a *app) getRoutes() http.Handler {

	r := a.baseRouter()

	r.Route("/api/v0/auth", func(r chi.Router) {
		r.Post("/login/", a.loginHandler)
		r.Post("/register/", a.registerHandler)
		r.Post("/reset/", a.resetPasswordHandler)
		r.Get("/logout/", a.logoutHandler)
	})

	// TODO: if not prod
	r.HandleFunc("/gdebug", graphqlDebugHandler)

	rootResolver := resolver.NewRootResolver(a.app)
	graphqlSchema := graphql.MustParseSchema(schema.GetRootSchema(), rootResolver)
	r.Handle("/query", &relay.Handler{Schema: graphqlSchema})

	return r
}

func (a app) protectedRouter(r chi.Router) chi.Router {
	return r.With(
		jwtauth.Verify(a.jwtAuth, jwtauth.TokenFromHeader, jwtauth.TokenFromQuery),
		a.Authenticator,
	)
}

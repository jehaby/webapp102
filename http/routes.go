package http

import (
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/jehaby/webapp102/resolver"
	"github.com/jehaby/webapp102/schema"

	"github.com/go-chi/chi"
)

func (a *app) getRoutes() http.Handler {

	r := a.baseRouter()

	r.Route("/api/v0/auth", func(r chi.Router) {
		r.Post("/login/", a.loginHandler)
		r.Post("/register/", a.registerHandler)

		r.Post("/confirm/", a.confirmPasswordHandler)

		r.Get("/refresh/", a.refreshTokenHandler)
		r.With(a.authMiddleware).Get("/logout/", a.logoutHandler)

		r.Post("/resetRequest/", a.resetPasswordRequestHandler)
		r.Post("/resetAction/", a.resetPasswordActionHandler)
	})

	r.With(a.authMiddleware).Route("/api/v0/users", func(r chi.Router) {
		r.Get("/{uuid}/", a.userGetHandler)
		r.Put("/{uuid}/", a.userUpdateHandler)

		r.Post("/{uuid}/phones/", a.userPhonesCreateHandler)
		r.Delete("/{uuid}/phones/{phone_uuid}/", a.userPhonesDeleteHandler)
	})

	// TODO: if not prod
	r.HandleFunc("/gdebug", graphqlDebugHandler)

	rootResolver := resolver.NewRootResolver(a.app)
	graphqlSchema := graphql.MustParseSchema(schema.GetRootSchema(), rootResolver)
	r.Handle("/query", &relay.Handler{Schema: graphqlSchema})

	return r
}

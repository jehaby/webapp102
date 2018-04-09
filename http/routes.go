package http

import (
	"net/http"

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

	r.Route("/api/v0/categories/", func(r chi.Router) {
		r.Get("/", a.getCategories)
	})

	r.Route("/api/v0/ads/", func(r chi.Router) {
		r.Get("/", a.viewAdsHandler) // TODO: paginate middleware maybe
		a.protectedRouter(r).Post("/", a.createAdHandler)

		r.Route("/{UUID}", func(r chi.Router) {
			r.Use(a.adCtx)
			r.Get("/", a.viewAdHandler)
			a.protectedRouter(r).Put("/", a.editAdHandler)
			a.protectedRouter(r).Delete("/", a.deleteAdHandler)
		})
	})

	return r
}

func (a app) protectedRouter(r chi.Router) chi.Router {
	return r.With(
		jwtauth.Verify(a.jwtAuth, jwtauth.TokenFromHeader, jwtauth.TokenFromQuery),
		a.Authenticator,
	)
}

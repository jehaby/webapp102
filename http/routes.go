package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
)

func (a *app) getRoutes() http.Handler {

	r := a.baseRouter()

	// TODO: middlewares

	r.Route("/api/v0/auth", func(r chi.Router) {
		r.Post("/login/", a.loginHandler)
		r.Post("/register/", a.registerHandler)
		r.Post("/reset/", a.resetPasswordHandler)
		r.Get("/logout/", a.logoutHandler)
	})

	r.Route("/api/v0/ads/", func(r chi.Router) {
		r.Get("/", a.viewAdsHandler) // TODO: paginate middleware maybe
		r.With(jwtauth.Verifier(a.jwtAuth)).Post("/", a.createAdHandler) // TODO: with auth middleware

		r.Route("/{UUID}", func(r chi.Router) {
			r.Use(a.adCtx)
			r.Get("/", a.viewAdHandler)
			r.Put("/", a.editAdHandler)
			r.Delete("/", a.deleteAdHandler)
		})
	})

	return r
}

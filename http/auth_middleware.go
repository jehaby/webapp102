package http

import (
	"context"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-chi/render"
	"github.com/jehaby/webapp102/service"
	"github.com/jehaby/webapp102/service/auth"
)

// authCookieMiddleware adds jwt token to context, if "auth" cookie is present;
func (a *app) authCookieMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, _ := r.Cookie(authCookieName)
		if c != nil {
			r = r.WithContext(
				context.WithValue(r.Context(), auth.StrTokenCtxKey, c.Value),
			)
		}
		next.ServeHTTP(w, r)
	})
}

// authMiddleware does actual authentication, use for protected routes
func (a *app) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, err := service.AddUserToCtx(r.Context(), a.app.Service.Auth, a.app.Service.User)
		if err != nil {
			// TODO: might be application error; logging (but not always)
			spew.Dump(err)
			render.Render(w, r, errUnauthorized(err))
			return
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

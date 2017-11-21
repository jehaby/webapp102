package http

import "net/http"

func (a *app) viewAdsHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "view ads", http.StatusNotImplemented)
}

func (a *app) viewAdHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "view ad", http.StatusNotImplemented)
}

func (a *app) createAdHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "create ad", http.StatusNotImplemented)
}

func (a *app) editAdHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "edit", http.StatusNotImplemented)
}

func (a *app) deleteAdHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "delete", http.StatusNotImplemented)
}

func (a *app) adCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

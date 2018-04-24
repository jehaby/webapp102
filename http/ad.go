package http

import (
	"encoding/json"
	"net/http"

	"github.com/jehaby/webapp102/entity"
)

func (a *app) viewAdsHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "view ads", http.StatusNotImplemented)
}

type adResponse struct {
	Ad entity.Ad `json:"ad"`
}

func (a *app) viewAdHandler(w http.ResponseWriter, r *http.Request) {
	ad := r.Context().Value("ad").(*entity.Ad)

	encoder := json.NewEncoder(w)
	encoder.Encode(adResponse{*ad}) // TODO: close w?
}

func (a *app) createAdHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "edit", http.StatusNotImplemented)

}

func (a *app) editAdHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "edit", http.StatusNotImplemented)
}

func (a *app) deleteAdHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "edit", http.StatusNotImplemented)
}

func (a *app) adCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		http.Error(w, "edit", http.StatusNotImplemented)
		// next.ServeHTTP(w, r.WithContext(ctx))
	})
}

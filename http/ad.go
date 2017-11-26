package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/render"
	"github.com/jehaby/webapp102/entity"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
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
	logger := a.log().With("handler", "/ads/create") // TODO: middleware maybe

	request := struct {
		Name        string `validate:"required"`
		Description string `validate:"required,min=10"`
	}{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		withErr(logger, err).Infow("decoding json", "err", err)
		http.Error(w, "bad json", 400)
		return
	}

	if err = a.validator.Struct(request); err != nil {
		withErr(logger, err).Debugw("validation", "request", request)
		http.Error(w, err.Error(), 400)
		return
	}

	ad := entity.Ad{
		Name:        request.Name,
		Description: request.Description,
		User:        mustUserFromCtx(r.Context()), // TODO: check panicking ok
	}

	res, err := a.app.Ad.Repo.Create(ad)
	if err != nil {
		code, msg := 0, "" // TODO: refactor (use render, also see auth)
		if e, ok := errors.Cause(err).(*pq.Error); ok && e.Code.Name() == "unique violation" {
			// user or email already exists
			code, msg = 400, "unique violation"
			withErr(logger, err).Infow(msg, "user", ad)
		} else {
			code, msg = 500, "error from save"
			withErr(logger, err).Errorw(msg, "user", ad)
		}

		http.Error(w, msg, code)
		return
	}

	enc := json.NewEncoder(w) // TODO: should I close it?
	enc.Encode(adResponse{Ad: *res})
}

func (a *app) editAdHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "edit", http.StatusNotImplemented)
}

func (a *app) deleteAdHandler(w http.ResponseWriter, r *http.Request) {
	ad := r.Context().Value("ad").(*entity.Ad)

	if err := a.app.Ad.Repo.Delete(ad.UUID); err != nil {
		http.Error(w, "Service error", 500) // TODO: better error (logging, text)
		return
	}

	w.Write([]byte("deleted")) // TODO: better
}

func (a *app) adCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		spew.Dump("in ad ctx", chi.URLParam(r, "UUID"))
		uuid, err := uuid.FromString(chi.URLParam(r, "UUID"))
		if err != nil {
			render.Render(w, r, errInvalidRequest(err))
			return
		}

		ad, err := a.app.Ad.Repo.GetByUUID(uuid)
		if err != nil {
			render.Render(w, r, errNotFound(err)) // TODO: public and logging error text
			return
		}

		ctx := context.WithValue(r.Context(), "ad", ad)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

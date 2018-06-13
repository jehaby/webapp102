package http

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	uuid "github.com/satori/go.uuid"

	"github.com/jehaby/webapp102/service"
)

func (a *app) userGetHandler(w http.ResponseWriter, r *http.Request) {
	userUUID, err := uuid.FromString(chi.URLParam(r, "uuid"))
	if err != nil {
		render.Render(w, r, errInvalidRequest(err))
		return
	}

	user, err := a.app.Service.User.GetByUUID(userUUID)
	if err != nil {
		render.Render(w, r, a.createRendererErr(err))
		return
	}

	render.JSON(w, r, user)
}

func (a *app) userUpdateHandler(w http.ResponseWriter, r *http.Request) {
	loggedInUser := service.UserFromCtx(r.Context())
	if loggedInUser == nil {
		render.Render(w, r, errNotLoggedIn500)
		return
	}

	userUUID, err := uuid.FromString(chi.URLParam(r, "uuid"))
	if err != nil {
		render.Render(w, r, errInvalidRequest(err))
		return
	}

	req := struct {
		Email        *string
		DefaultPhone *uuid.UUID `json:"default_phone"`
	}{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		render.Render(w, r, errInvalidRequest(err))
		return
	}

	args := service.UserUpdateArgs{
		Email:        req.Email,
		DefaultPhone: req.DefaultPhone,
	}
	user, err := a.app.Service.User.Update(r.Context(), userUUID, args)
	if err != nil {
		render.Render(w, r, a.createRendererErr(err))
		return
	}

	// updating jwt cookie, because frontend uses it's values for showing some user's info in profile
	// (it might be bad)
	if loggedInUser.UUID == userUUID {
		tkn, err := a.app.Service.Auth.TokenFromUser(*loggedInUser, jwtExpirationTime)
		if err != nil {
			render.Render(w, r, err500(err))
			return
		}
		http.SetCookie(w, createJwtCookie(tkn, a.cfg.HTTP.SecureJWTCookie))
		w.Write([]byte(tkn))
	}

	render.JSON(w, r, user)
}

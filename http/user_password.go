package http

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
)

// resetPasswordRequestHandler is first step in resetting password;
// it finds user by name or email, creates token and sends it to the user
func (a *app) resetPasswordRequestHandler(w http.ResponseWriter, r *http.Request) {
	req := struct {
		Name string `validate:"required"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		render.Render(w, r, errInvalidRequest(err))
		return
	}
	err = a.validator.Struct(req)
	if err != nil {
		render.Render(w, r, errInvalidRequest(err))
		return
	}

	err = a.app.Service.User.ProcessPasswordResetRequest(r.Context(), req.Name)
	if err != nil {
		render.Render(w, r, a.createRendererErr(err))
		return
	}

	render.JSON(w, r, "ok")
}

// resetPasswordActionHandler is the second step in resetting password;
// it finds user by token (from email (and checks token not expired))
// processes new password and saves it to database
func (a *app) resetPasswordActionHandler(w http.ResponseWriter, r *http.Request) {
	req := struct {
		Token    string `validate:"required`
		Password string `validate:"required"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		render.Render(w, r, errInvalidRequest(err))
		return
	}
	err = a.validator.Struct(req)
	if err != nil {
		render.Render(w, r, errInvalidRequest(err))
		return
	}

	user, err := a.app.Service.User.ProcessPasswordResetAction(r.Context(), req.Token, req.Password)
	if err != nil {
		render.Render(w, r, a.createRendererErr(err))
		return
	}

	tkn, err := a.app.Service.Auth.TokenFromUser(user, jwtExpirationTime)
	if err != nil {
		// got error but password was resetted sucessfully; what to do? (it's very unlikely that this happens)
		render.Render(w, r, a.createRendererErr(err))
		return
	}
	http.SetCookie(w, createJwtCookie(tkn, a.cfg.HTTP.SecureJWTCookie))
	render.JSON(w, r, "ok")
}

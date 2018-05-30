package http

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"

	"github.com/jehaby/webapp102/entity"
	"github.com/jehaby/webapp102/service"
)

var (
	userCtxKey = &contextKey{"User"}
)

const (
	authCookieName    = "jwt"
	jwtExpirationTime = time.Hour * 24 * 7
)

func (a *app) loginHandler(w http.ResponseWriter, r *http.Request) {

	logger := a.log().With("handler", "login")

	request := struct {
		Name     string `validate:"required"`
		Password string `validate:"required,min=3"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		withErr(logger, err).Errorw("decoding json", "json") // TODO: request body
		http.Error(w, "got error unmarshaling json", 400)
		return
	}

	if err = a.validator.Struct(request); err != nil {
		withErr(logger, err).Infow("validation", "request", request)
		http.Error(w, err.Error(), 400)
		return
	}

	user, err := a.app.User.GetByNameOrEmail(request.Name)
	if err != nil {
		code, msg := 0, ""
		if errors.Cause(err) == sql.ErrNoRows {
			// TODO: this probably doesn't work now
			msg = "couldn't find user"
			code = 404
		} else {
			msg = "something bad happened"
			code = 500
		}
		withErr(logger, err).Infow(msg, "user", user, "err", err)
		http.Error(w, msg, code)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		// TODO: log maybe
		http.Error(w, "bad password", http.StatusUnauthorized)
		return
	}

	tkn, err := a.app.Service.Auth.TokenFromUser(user, jwtExpirationTime)
	if err != nil {
		withErr(logger, err).Errorw("encoding jwt")
		http.Error(w, "encoding jwt", 500)
		return
	}
	http.SetCookie(w, createJwtCookie(tkn, a.cfg.HTTP.SecureJWTCookie))
	w.Write([]byte(tkn))
}

func (a *app) registerHandler(w http.ResponseWriter, r *http.Request) {

	logger := a.log().With("handler", "register")

	request := struct {
		Name     string `validate:"required"`
		Email    string `validate:"required,email"`
		Password string `validate:"required,min=3"`
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

	// TODO: move somewhere (service layer?)
	pass, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		withErr(logger, err).Errorw("encrypting password", "err", err)
		http.Error(w, "encrypting password", 500)
	}

	user := entity.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: string(pass),
	}

	if err = a.app.User.Repo.Save(user); err != nil {
		code, msg := 0, ""
		if e, ok := errors.Cause(err).(*pq.Error); ok && e.Code.Name() == "unique violation" {
			// user or email already exists
			code, msg = 400, "unique violation"
			withErr(logger, err).Infow(msg, "user_name", user.Name, "user_email", user.Email)
		} else {
			code, msg = 500, "error from save"
			withErr(logger, err).Errorw(msg, "user_name", user.Name, "user_email", user.Email)
		}

		http.Error(w, msg, code)
		return
	}

	tkn, err := a.app.Service.Auth.TokenFromUser(user, jwtExpirationTime)
	if err != nil {
		withErr(logger, err).Errorw("encoding jwt")
		http.Error(w, "encoding jwt", 500)
		return
	}
	http.SetCookie(w, createJwtCookie(tkn, a.cfg.HTTP.SecureJWTCookie))
	w.Write([]byte(tkn))
}

func createJwtCookie(jwtToken string, secure bool) *http.Cookie {
	return &http.Cookie{
		Name:     authCookieName,
		Value:    jwtToken,
		HttpOnly: true,
		Secure:   secure,
		Path:     "/",
		MaxAge:   int(jwtExpirationTime.Seconds()),
	}
}

func (a *app) resetPasswordHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}

func (a *app) refreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	// check user ok
	// refresh token
	ctx, err := service.AddUserToCtx(r.Context(), a.app.Service.Auth, a.app.Service.User)
	if err != nil {
		// TODO: might be application error; logging (but not always)
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	user := service.UserFromCtx(ctx)
	tkn, err := a.app.Service.Auth.TokenFromUser(*user, jwtExpirationTime)
	if err != nil {
		a.app.Logger.WithError(err).Errorw("encoding jwt")
		http.Error(w, "encoding jwt", 500)
		return
	}

	http.SetCookie(w, createJwtCookie(tkn, a.cfg.HTTP.SecureJWTCookie))
	w.Write([]byte(tkn))
}

func (a *app) logoutHandler(w http.ResponseWriter, r *http.Request) {
	// check user logged in
	// update user in db
	// remove cookie
	ctx, err := service.AddUserToCtx(r.Context(), a.app.Service.Auth, a.app.Service.User)
	if err != nil {
		// TODO: might be application error; logging (but not always)
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	_, err = a.app.Service.User.Update(
		ctx,
		service.UserFromCtx(ctx).UUID,
		service.UserUpdateArgs{LastLogout: pointer.ToTime(time.Now())},
	)
	if err != nil {
		a.app.Logger.WithError(err).Errorw("couldn't update user")
		http.Error(w, "couldn't update", http.StatusInternalServerError)
		return
	}

	c := &http.Cookie{
		Name:     authCookieName,
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
	}
	http.SetCookie(w, c)
	w.Write([]byte("ok"))
}

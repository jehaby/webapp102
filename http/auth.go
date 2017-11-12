package http

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-chi/jwtauth"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"

	"github.com/jehaby/webapp102/entity"
)

func (a *app) loginHandler(w http.ResponseWriter, r *http.Request) {

	logger := a.log().With("handler", "login")

	request := struct {
		Name     string `validate:"required"`
		Password string `validate:"required,min=3"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		logger.Errorw("decoding json", "err", err, "json") // TODO: request body
		http.Error(w, "got error unmarshaling json", 400)
		return
	}

	if err = a.validator.Struct(request); err != nil {
		logger.Infow("validation", "err", err, "request", request)
		http.Error(w, err.Error(), 400)
		return
	}

	user, err := a.app.UR.GetByName(request.Name)
	if err != nil {
		code, msg := 0, ""
		if errors.Cause(err) == sql.ErrNoRows {
			msg = "couldn't find user"
			code = 404
		} else {
			msg = "something bad happened"
			code = 500
		}
		logger.Infow(msg, "user", user, "err", err)
		http.Error(w, msg, code)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		// TODO: log maybe
		http.Error(w, "bad password", http.StatusUnauthorized)
		return
	}

	_, tkn, err := tokenAuth.Encode(jwtauth.Claims{"name": user.Name, "email": user.Email})
	if err != nil {
		logger.Errorw("encoding jwt", "err", err, "user", user)
		http.Error(w, "encoding jwt", 500)
		return
	}

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
		logger.Infow("decoding json", "err", err)
		http.Error(w, "bad json", 400)
		return
	}

	if err = a.validator.Struct(request); err != nil {
		logger.Debugw("validation", "err", err, "request", request)
		http.Error(w, err.Error(), 400)
		return
	}

	// TODO: move somewhere (service layer?)
	pass, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Errorw("encrypting password", "err", err)
		http.Error(w, "encrypting password", 500)
	}

	user := entity.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: string(pass),
	}

	if err = a.app.UR.Save(user); err != nil {
		code, msg := 0, ""
		if e, ok := errors.Cause(err).(*pq.Error); ok && e.Code.Name() == "unique violation" {
			// user or email already exists
			code, msg = 400, "unique violation"
			logger.Infow(msg, "err", err, "user", user)
		} else {
			code, msg = 500, "error from save"
			logger.Errorw(msg, "err", err, "user", user)
		}

		http.Error(w, msg, code)
		return
	}

	claims := jwtauth.Claims{"name": user.Name, "email": user.Email}
	_, tkn, err := tokenAuth.Encode(claims)
	if err != nil {
		logger.Errorw("encoding jwt", "err", err, "claims", claims)
		http.Error(w, "encoding jwt", 500)
		return
	}

	w.Write([]byte(tkn))

}

func (a *app) resetPasswordHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}

func (a *app) logoutHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}

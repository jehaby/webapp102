package http

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/jehaby/webapp102/entity"
)

var (
	userCtxKey = &contextKey{"User"}
)

const jwtExpirationTime = 24 * time.Hour

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

	user, err := a.app.User.Repo.GetByName(request.Name)
	if err != nil {
		code, msg := 0, ""
		if errors.Cause(err) == sql.ErrNoRows {
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

	claims := jwtauth.Claims{"user": responsefromUser(*user)}.SetExpiryIn(jwtExpirationTime)
	_, tkn, err := a.jwtAuth.Encode(claims)
	if err != nil {
		withErr(logger, err).Errorw("encoding jwt", "user", user)
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
			withErr(logger, err).Infow(msg, "user", user)
		} else {
			code, msg = 500, "error from save"
			withErr(logger, err).Errorw(msg, "user", user)
		}

		http.Error(w, msg, code)
		return
	}

	claims := jwtauth.Claims{"user": responsefromUser(user)}.SetExpiryIn(jwtExpirationTime)
	_, tkn, err := a.jwtAuth.Encode(claims)
	if err != nil {
		withErr(logger, err).Errorw("encoding jwt", "claims", claims)
		http.Error(w, "encoding jwt", 500)
		return
	}

	w.Write([]byte(tkn))

}

type userResponse struct {
	UUID  uuid.UUID `json:"uuid"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

func responsefromUser(user entity.User) userResponse {
	return userResponse{
		UUID:  user.UUID,
		Name:  user.Name,
		Email: user.Email,
	}
}

func (a *app) resetPasswordHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}

func (a *app) logoutHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}

// Authenticator is an authentication middleware based on jwtauth.Authenticator
func (a *app) Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, claims, err := jwtauth.FromContext(r.Context())

		// TODO: metrics, logging
		if err != nil {
			http.Error(w, http.StatusText(401), 401)
			return
		}

		if token == nil || !token.Valid {
			http.Error(w, http.StatusText(401), 401)
			return
		}

		userUUID, ok := claims["user_uuid"]
		if !ok {
			http.Error(w, http.StatusText(401), 401)
			return
		}

		user, err := a.app.User.Repo.GetByUUID(uuid.FromStringOrNil(userUUID.(string))) // TODO: types
		if err != nil {
			http.Error(w, http.StatusText(401), 401)
			return
		}

		// Got user, pass it through
		ctx := context.WithValue(r.Context(), userCtxKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func userFromCtx(ctx context.Context) (*entity.User, error) {
	if user, ok := ctx.Value(userCtxKey).(*entity.User); ok {
		return user, nil
	}
	return nil, fmt.Errorf("userFromCtx: no valid user")
}

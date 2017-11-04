package http

import (
	"net/http"

	"fmt"

	"github.com/go-chi/jwtauth"
	"github.com/jehaby/webapp102/entity"
	"github.com/pkg/errors"
)

func (a *app) loginHandler(w http.ResponseWriter, r *http.Request) {

	token, claims, err := jwtauth.FromContext(r.Context())

	if err != nil {
		http.Error(w, http.StatusText(401), 401)
		return
	}

	if token == nil || !token.Valid {
		http.Error(w, http.StatusText(401), 401)
		return
	}

	creds, err := getCredentials(claims)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	user, err := a.storage.GetUser(creds)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	// TODO: generate some temporary jwt token, save it to some storage, pass to frontend

	w.Write([]byte(fmt.Sprintf("Hello, %s", user.Login)))
}

func getCredentials(c jwtauth.Claims) (entity.Credentials, error) {
	var res entity.Credentials
	var err error

	res.Login, err = getString(c, "login")
	if err != nil {
		return entity.Credentials{}, err
	}

	res.Password, err = getString(c, "password")
	if err != nil {
		return entity.Credentials{}, err
	}

	return res, nil
}

func getString(c jwtauth.Claims, key string) (string, error) {
	tmp, ok := c.Get(key)
	if !ok {
		return "", errors.Errorf("key '%s' not found in claims", key)
	}
	if val, ok := tmp.(string); ok {
		return val, nil
	}
	return "", errors.Errorf("key '%s' is not a string", key)

}

func (a *app) registerHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}

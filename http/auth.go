package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/jwtauth"
	"golang.org/x/crypto/bcrypt"

	"github.com/jehaby/webapp102/entity"
)

func (a *app) loginHandler(w http.ResponseWriter, r *http.Request) {

	request := struct {
		Name     string `validate:"required"`
		Password string `validate:"required,min=3"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Printf("got error unmarshaling json: %v", err)
		http.Error(w, "got error unmarshaling json", 400)
		return
	}

	if err = a.validator.Struct(request); err != nil {
		log.Printf("login: validaion error: %v", err)
		http.Error(w, err.Error(), 400)
		return
	}

	user, err := a.app.UR.GetByName(request.Name)
	if err != nil {
		log.Printf("couldnt find user: %v", err)
		http.Error(w, err.Error(), 404)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		log.Printf("bcrypt error: %v", err)
		http.Error(w, "bad password", http.StatusUnauthorized)
		return
	}

	_, tkn, err := tokenAuth.Encode(jwtauth.Claims{"name": user.Name, "email": user.Email})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write([]byte(tkn))
}

func (a *app) registerHandler(w http.ResponseWriter, r *http.Request) {
	request := struct {
		Name     string `validate:"required"`
		Email    string `validate:"required,email"`
		Password string `validate:"required,min=3"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Printf("got error unmarshaling json: %v", err)
		http.Error(w, "got error unmarshaling json", 400)
		return
	}

	if err = a.validator.Struct(request); err != nil {
		log.Printf("register: validaion error: %v", err)
		http.Error(w, err.Error(), 400)
		return
	}

	user := entity.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	if err = a.app.UR.Save(user); err != nil {
		log.Printf("couldnt save user: %v", err)
		http.Error(w, err.Error(), 500) // TODO: could be bad request if user already exists
		return
	}

	_, tkn, err := tokenAuth.Encode(jwtauth.Claims{"name": user.Name, "email": user.Email})
	if err != nil {
		http.Error(w, err.Error(), 500)
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

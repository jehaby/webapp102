package http

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/satori/go.uuid"

	"github.com/jehaby/webapp102/service"
)

func (a *app) userPhonesCreateHandler(w http.ResponseWriter, r *http.Request) {
	loggedInUser := service.UserFromCtx(r.Context())
	if loggedInUser == nil {
		render.Render(w, r, errNotLoggedIn500)
		return
	}

	if !loggedInUser.CanEdit(uuid.FromStringOrNil(chi.URLParam(r, "uuid"))) {
		render.Render(w, r, errUnauthorized(nil))
		return
	}

	req := service.CreatePhoneArgs{}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		render.Render(w, r, errInvalidRequest(err))
		return
	}

	phone, err := a.app.Service.User.CreatePhone(req)
	if err != nil {
		render.Render(w, r, a.createRendererErr(err))
		return
	}

	render.JSON(w, r, phone)
}

func (a *app) userPhonesDeleteHandler(w http.ResponseWriter, r *http.Request) {
	loggedInUser := service.UserFromCtx(r.Context())
	if loggedInUser == nil {
		render.Render(w, r, errNotLoggedIn500)
		return
	}

	args := service.DeletePhoneArgs{
		UserUUID:  uuid.FromStringOrNil(chi.URLParam(r, "uuid")),
		PhoneUUID: uuid.FromStringOrNil(chi.URLParam(r, "phone_uuid")),
	}

	if !loggedInUser.CanEdit(args.UserUUID) {
		render.Render(w, r, errUnauthorized(nil))
		return
	}

	user, err := a.app.Service.User.DeletePhone(args)
	if err != nil {
		render.Render(w, r, a.createRendererErr(err))
		return
	}

	render.JSON(w, r, user)
}

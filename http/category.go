package http

import (
	"net/http"

	"github.com/go-chi/chi/render"
	"github.com/jehaby/webapp102/entity"
)

type categoryRespList map[int64]*entity.Category

func (a *app) getCategories(w http.ResponseWriter, r *http.Request) {
	res, err := a.app.Ad.Repo.Category.GetAll()
	if err != nil {
		render.Render(w, r, err500(err))
	}

	resp := make(categoryRespList)
	for _, cat := range res {
		resp[cat.ID] = cat
	}

	render.JSON(w, r, resp)
}

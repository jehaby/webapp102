package resolver

import (
	"context"

	"github.com/jehaby/webapp102/entity"
)

type categoryGetter interface {
	GetAll() ([]*entity.Category, error)
}

func (r *Resolver) Categories(ctx context.Context) ([]*categoryResolver, error) {
	categories, err := r.app.Ad.Repo.Category.GetAll()
	if err != nil {
		return nil, err
	}

	res := make([]*categoryResolver, 0, len(categories))
	for i, _ := range categories {
		res = append(res, &categoryResolver{categories[i]})
	}
	return res, nil
}

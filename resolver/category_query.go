package resolver

import (
	"context"
)

func (r *Resolver) Categories(ctx context.Context) ([]*categoryResolver, error) {
	categories, err := r.app.Service.Category.GetAll()
	if err != nil {
		return nil, err
	}

	res := make([]*categoryResolver, 0, len(categories))
	for i, _ := range categories {
		res = append(res, &categoryResolver{categories[i]})
	}
	return res, nil
}

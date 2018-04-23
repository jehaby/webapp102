package resolver

import (
	"fmt"

	"github.com/graph-gophers/graphql-go"

	"github.com/jehaby/webapp102/entity"
)

type adResolver struct {
	ad *entity.Ad
}

func (r *adResolver) Uuid() graphql.ID {
	return graphql.ID(r.ad.UUID.String())
}

func (r *adResolver) Name() string {
	return r.ad.Name
}

func (r *adResolver) Description() string {
	return r.ad.Description
}

func (r *adResolver) Product() (*productResolver, error) {
	cr, err := newProductResolver(r.ad.Product)
	if err != nil {
		return nil, err
	}
	return cr, nil
}

func (r *adResolver) User() (*userResolver, error) {
	ur, err := newUserResolver(r.ad.User)
	if err != nil {
		return nil, err
	}
	return ur, nil
}

func (r *adResolver) Locality() (*localityResolver, error) {
	if r.ad.Locality == nil {
		return nil, fmt.Errorf("adResolver.Locality(): Locality is nil %v", r.ad)
	}
	return &localityResolver{*r.ad.Locality}, nil
}

func (r *adResolver) Price() int32 {
	return int32(r.ad.Price)
}

func (r *adResolver) Currency() entity.Currency {
	return r.ad.Currency
}

func (r *adResolver) CreatedAt() graphql.Time {
	return graphql.Time{r.ad.CreatedAt}
}

func (r *adResolver) UpdatedAt() *graphql.Time {
	if r.ad.UpdatedAt == nil {
		return nil
	}
	return &graphql.Time{*r.ad.UpdatedAt}
}

// TODO:
/* func (r *adResolver) Category() *categoryResolver {
	return
} */

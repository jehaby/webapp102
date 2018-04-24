package resolver

import (
	"fmt"

	"github.com/AlekSi/pointer"
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
	if r.ad.Product == nil {
		return nil, nil
	}
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

func (r *adResolver) Weight() *int32 {
	if r.ad.Weight == 0 {
		return nil
	}
	return pointer.ToInt32(int32(r.ad.Weight))
}

func (r *adResolver) Brand() *brandResolver {
	if r.ad.Brand == nil {
		return nil
	}
	return &brandResolver{*r.ad.Brand}
}

func (r *adResolver) Properties() *string {
	if r.ad.Properties == "" {
		return nil
	}
	return pointer.ToString(r.ad.Properties)
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

func (r *adResolver) Category() *categoryResolver {
	if r.ad.Category == nil {
		return nil
	}
	return &categoryResolver{r.ad.Category}
}

package resolver

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/jehaby/webapp102/entity"
)

type adResolver struct {
	ad *entity.Ad
}

func (r *adResolver) Uuid() *graphql.ID {
	var res *graphql.ID
	*res = graphql.ID(r.ad.UUID.String())
	return res
}

func (r *adResolver) Name() string {
	return r.ad.Name
}

func (r *adResolver) Description() string {
	return r.ad.Description
}

func (r *adResolver) Component() *componentResolver {
	// TODO:
	return &componentResolver{}
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

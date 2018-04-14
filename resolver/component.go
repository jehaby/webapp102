package resolver

import (
	"context"
	"errors"

	"github.com/graph-gophers/graphql-go"
	"github.com/jehaby/webapp102/entity"
)

func (r *Resolver) Components(ctx context.Context, args componentsArgs) (*componentResolver, error) {
	return &componentResolver{}, nil
}

type componentsArgs struct {
	Name *string
}

func newComponentResolver(e *entity.Component) (*componentResolver, error) {
	if e == nil {
		return nil, errors.New("newComponentResolver: passed nil entity")
	}
	return &componentResolver{e}, nil
}

type componentResolver struct {
	e *entity.Component
}

func (r *componentResolver) ID() graphql.ID {
	return graphql.ID(r.e.ID)
}

func (r *componentResolver) Name() string {
	return r.e.Name
}

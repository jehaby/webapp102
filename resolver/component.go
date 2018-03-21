package resolver

import (
	"context"

	"github.com/graph-gophers/graphql-go"
)

func (r *Resolver) Components(ctx context.Context, args componentsArgs) (*componentResolver, error) {
	return &componentResolver{}, nil
}

type componentsArgs struct {
	Name *string
}

type componentResolver struct{}

func (r *componentResolver) ID() graphql.ID {
	return graphql.ID("42")
}

func (r *componentResolver) Name() string {
	return "bikes"
}

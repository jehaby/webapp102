package resolver

import (
	"context"
	"strconv"

	graphql "github.com/graph-gophers/graphql-go"
)

func (r *Resolver) Component(ctx context.Context, args struct {
	ID graphql.ID
}) (*componentResolver, error) {
	id, err := strconv.ParseUint(string(args.ID), 10, 64)
	if err != nil {
		return nil, err
	}

	comp, err := r.app.Repo.Component.GetByID(uint32(id))
	if err != nil {
		return nil, err
	}
	return &componentResolver{comp}, nil
}

// func (r *Resolver) Components(ctx context.Context, args componentsArgs) (*componentResolver, error) {
// 	return &componentResolver{}, nil
// }

// type componentsArgs struct {
// 	Name *string
// }

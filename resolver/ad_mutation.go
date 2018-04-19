package resolver

import (
	"context"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/jehaby/webapp102/service"
)

type AdCreateInput struct {
	Name        string
	Description string
	UserUUID    graphql.ID
	ComponentID graphql.ID
	CategoryID  graphql.ID
}

func (r *Resolver) AdCreate(ctx context.Context, args struct {
	Input AdCreateInput
}) (*adResolver, error) {
	ad, err := r.app.Service.Ad.Create(service.AdCreateArgs{
		Name:        args.Input.Name,
		Description: args.Input.Description,
		UserUUID:    string(args.Input.UserUUID),
		ComponentID: string(args.Input.ComponentID),
	})
	if err != nil {
		return nil, err
	}

	return &adResolver{ad}, nil
}

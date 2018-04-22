package resolver

import (
	"context"

	"github.com/AlekSi/pointer"
	graphql "github.com/graph-gophers/graphql-go"
	uuid "github.com/satori/go.uuid"

	"github.com/jehaby/webapp102/entity"
	"github.com/jehaby/webapp102/service"
)

type AdCreateInput struct {
	Name        string
	Description string
	UserUUID    graphql.ID
	ComponentID graphql.ID
	Price       float64 // TODO: wtf?
	Currency    entity.Currency
}

func (r *Resolver) AdCreate(ctx context.Context, args struct {
	Input AdCreateInput
}) (*adResolver, error) {

	ad, err := r.app.Service.Ad.Create(service.AdCreateArgs{
		Name:        args.Input.Name,
		Description: args.Input.Description,
		UserUUID:    string(args.Input.UserUUID),
		ComponentID: string(args.Input.ComponentID),
		Price:       int64(args.Input.Price),
		Currency:    args.Input.Currency,
	})
	if err != nil {
		return nil, err
	}

	return &adResolver{ad}, nil
}

type AdUpdateInput struct {
	Name        *string
	Description *string
	ComponentID *graphql.ID
	Price       *float64 // TODO: graphql-float
	Currency    *entity.Currency
}

func (r *Resolver) AdUpdate(ctx context.Context, args struct {
	UUID  graphql.ID
	Input AdUpdateInput
}) (*adResolver, error) {
	uuid, err := uuid.FromString(string(args.UUID))
	if err != nil {
		return nil, err
	}

	serviceArgs := service.AdUpdateArgs{
		Name:        args.Input.Name,
		Description: args.Input.Description,
		Currency:    args.Input.Currency,
	}

	if args.Input.Price != nil {
		serviceArgs.Price = pointer.ToInt64(int64(*args.Input.Price))
	}
	if args.Input.ComponentID != nil {
		serviceArgs.ComponentID = pointer.ToString(string(*args.Input.ComponentID))
	}

	ad, err := r.app.Service.Ad.Update(uuid, serviceArgs)
	if err != nil {
		return nil, err
	}

	return &adResolver{ad}, nil
}

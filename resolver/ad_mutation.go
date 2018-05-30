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
	CategoryID  graphql.ID
	Condition   entity.Condition
	LocalityID  graphql.ID
	Price       float64 // TODO: wtf?
	Currency    entity.Currency
	Weight      *float64
	BrandID     *graphql.ID
	Properties  *string
}

func (r *Resolver) AdCreate(ctx context.Context, args struct {
	Input AdCreateInput
}) (*adResolver, error) {

	var err error
	ctx, err = service.AddUserToCtx(ctx, r.app.Service.Auth, r.app.Service.User)
	if err != nil {
		return nil, err
	}

	serviceArgs := service.AdCreateArgs{
		Name:        args.Input.Name,
		Description: args.Input.Description,
		CategoryID:  string(args.Input.CategoryID),
		Condition:   args.Input.Condition,
		LocalityID:  string(args.Input.LocalityID),
		Price:       int64(args.Input.Price),
		Currency:    args.Input.Currency,
		Properties:  args.Input.Properties,
	}

	if args.Input.Weight != nil {
		serviceArgs.Weight = pointer.ToInt64(int64(*args.Input.Weight))
	}
	if args.Input.BrandID != nil {
		serviceArgs.BrandID = pointer.ToString(string(*args.Input.BrandID))
	}

	ad, err := r.app.Service.Ad.Create(ctx, serviceArgs)
	if err != nil {
		return nil, err
	}

	return &adResolver{ad}, nil
}

type AdUpdateInput struct {
	Name        *string
	Description *string
	Condition   *entity.Condition
	CategoryID  *graphql.ID
	ProductID   *graphql.ID
	LocalityID  *graphql.ID
	Price       *float64 // TODO: graphql-float
	Currency    *entity.Currency
	Weight      *float64
	BrandID     *graphql.ID
	Properties  *string
}

func (r *Resolver) AdUpdate(ctx context.Context, args struct {
	UUID  graphql.ID
	Input AdUpdateInput
}) (*adResolver, error) {

	var err error
	ctx, err = service.AddUserToCtx(ctx, r.app.Service.Auth, r.app.Service.User)
	if err != nil {
		return nil, err
	}

	uuid, err := uuid.FromString(string(args.UUID))
	if err != nil {
		return nil, err
	}

	serviceArgs := service.AdUpdateArgs{
		Name:        args.Input.Name,
		Description: args.Input.Description,
		Condition:   args.Input.Condition,
		Currency:    args.Input.Currency,
		Properties:  args.Input.Properties,
	}

	if args.Input.CategoryID != nil {
		serviceArgs.CategoryID = pointer.ToString(string(*args.Input.CategoryID))
	}
	if args.Input.Price != nil {
		serviceArgs.Price = pointer.ToInt64(int64(*args.Input.Price))
	}
	if args.Input.ProductID != nil {
		serviceArgs.ProductID = pointer.ToString(string(*args.Input.ProductID))
	}
	if args.Input.LocalityID != nil {
		serviceArgs.LocalityID = pointer.ToString(string(*args.Input.LocalityID))
	}
	if args.Input.Weight != nil {
		serviceArgs.Weight = pointer.ToInt64(int64(*args.Input.Weight))
	}
	if args.Input.BrandID != nil {
		serviceArgs.BrandID = pointer.ToString(string(*args.Input.BrandID))
	}

	ad, err := r.app.Service.Ad.Update(ctx, uuid, serviceArgs)
	if err != nil {
		return nil, err
	}

	return &adResolver{ad}, nil
}

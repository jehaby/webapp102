package resolver

import (
	"context"
	"strconv"

	"github.com/AlekSi/pointer"
	graphql "github.com/graph-gophers/graphql-go"

	"github.com/jehaby/webapp102/entity"
	"github.com/jehaby/webapp102/service"
)

type createComponentInput struct {
	Name           string
	CategoryID     graphql.ID
	ManufacturerID graphql.ID
}

func (r *Resolver) CreateComponent(ctx context.Context, args *struct {
	Input createComponentInput
}) (*componentResolver, error) {
	comp, err := r.app.Service.Component.Create(service.CreateComponentArgs{
		args.Input.Name,
		string(args.Input.CategoryID),
		string(args.Input.ManufacturerID),
	})
	if err != nil {
		return nil, err
	}

	return &componentResolver{comp}, nil
}

func (r *Resolver) RemoveComponent(ctx context.Context, args *struct {
	ID graphql.ID
}) (*componentResolver, error) {
	id, err := strconv.ParseInt(string(args.ID), 10, 64)
	if err != nil {
		return nil, err
	}

	err = r.app.Service.Component.Remove(id)
	if err != nil {
		return nil, err
	}

	return &componentResolver{&entity.Component{ID: id}}, nil
}

type updateComponentInput struct {
	Name           *string
	CategoryID     *graphql.ID
	ManufacturerID *graphql.ID
}

func (r *Resolver) UpdateComponent(ctx context.Context, args *struct {
	ID    graphql.ID
	Input updateComponentInput
}) (*componentResolver, error) {
	id, err := strconv.ParseInt(string(args.ID), 10, 64)
	if err != nil {
		return nil, err
	}

	serviceArgs := service.UpdateComponentArgs{
		Name: args.Input.Name,
	}
	if args.Input.CategoryID != nil {
		serviceArgs.CategoryID = pointer.ToString(string(*args.Input.CategoryID))
	}
	if args.Input.ManufacturerID != nil {
		serviceArgs.ManufacturerID = pointer.ToString(string(*args.Input.ManufacturerID))
	}

	e, err := r.app.Service.Component.Update(id, serviceArgs)
	if err != nil {
		return nil, err
	}

	return &componentResolver{e}, nil
}

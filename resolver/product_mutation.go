package resolver

import (
	"context"
	"strconv"

	"github.com/AlekSi/pointer"
	graphql "github.com/graph-gophers/graphql-go"

	"github.com/jehaby/webapp102/entity"
	"github.com/jehaby/webapp102/service"
)

type createProductInput struct {
	Name           string
	CategoryID     graphql.ID
	BrandID graphql.ID
}

func (r *Resolver) CreateProduct(ctx context.Context, args *struct {
	Input createProductInput
}) (*productResolver, error) {
	product, err := r.app.Service.Product.Create(service.CreateProductArgs{
		args.Input.Name,
		string(args.Input.CategoryID),
		string(args.Input.BrandID),
	})
	if err != nil {
		return nil, err
	}

	return &productResolver{product}, nil
}

func (r *Resolver) RemoveProduct(ctx context.Context, args *struct {
	ID graphql.ID
}) (*productResolver, error) {
	id, err := strconv.ParseInt(string(args.ID), 10, 64)
	if err != nil {
		return nil, err
	}

	err = r.app.Service.Product.Remove(id)
	if err != nil {
		return nil, err
	}

	return &productResolver{&entity.Product{ID: id}}, nil
}

type updateProductInput struct {
	Name           *string
	CategoryID     *graphql.ID
	BrandID *graphql.ID
}

func (r *Resolver) UpdateProduct(ctx context.Context, args *struct {
	ID    graphql.ID
	Input updateProductInput
}) (*productResolver, error) {
	id, err := strconv.ParseInt(string(args.ID), 10, 64)
	if err != nil {
		return nil, err
	}

	serviceArgs := service.UpdateProductArgs{
		Name: args.Input.Name,
	}
	if args.Input.CategoryID != nil {
		serviceArgs.CategoryID = pointer.ToString(string(*args.Input.CategoryID))
	}
	if args.Input.BrandID != nil {
		serviceArgs.BrandID = pointer.ToString(string(*args.Input.BrandID))
	}

	e, err := r.app.Service.Product.Update(id, serviceArgs)
	if err != nil {
		return nil, err
	}

	return &productResolver{e}, nil
}

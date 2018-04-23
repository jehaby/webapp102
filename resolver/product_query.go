package resolver

import (
	"context"
	"strconv"

	graphql "github.com/graph-gophers/graphql-go"
)

func (r *Resolver) Product(ctx context.Context, args struct {
	ID graphql.ID
}) (*productResolver, error) {
	id, err := strconv.ParseInt(string(args.ID), 10, 64)
	if err != nil {
		return nil, err
	}

	product, err := r.app.Service.Product.GetByID(id)
	if err != nil {
		return nil, err
	}

	return &productResolver{product}, nil
}

package resolver

import (
	"context"
	"errors"

	"github.com/davecgh/go-spew/spew"
	graphql "github.com/graph-gophers/graphql-go"
)

func (r *Resolver) Properties(ctx context.Context, args *struct{ CategoryID graphql.ID }) (*[]*propertyResolver, error) {
	spew.Dump(args)

	if args == nil {
		return nil, errors.New("category must be set!")
	}

	properties, err := r.app.Service.Property.GetByCategory(ctx, string(args.CategoryID))
	if err != nil {
		return nil, err
	}

	res := make([]*propertyResolver, 0, len(properties))
	for i, _ := range properties {
		res = append(res, &propertyResolver{*properties[i]})
	}

	return &res, nil
}

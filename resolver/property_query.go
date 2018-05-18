package resolver

import (
	"context"
	"errors"
	"strconv"

	graphql "github.com/graph-gophers/graphql-go"
)

func (r *Resolver) Properties(ctx context.Context, args *struct{ CategoryID graphql.ID }) (*[]*propertyResolver, error) {

	if args == nil {
		return nil, errors.New("category must be set!")
	}
	cid, err := strconv.ParseInt(string(args.CategoryID), 10, 64)
	if err != nil {
		return nil, err
	}

	properties, err := r.app.Service.Property.GetByCategory(ctx, cid)
	if err != nil {
		return nil, err
	}

	res := make([]*propertyResolver, 0, len(properties))
	for i, _ := range properties {
		res = append(res, &propertyResolver{*properties[i]})
	}

	return &res, nil
}

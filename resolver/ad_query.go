package resolver

import (
	"context"

	"github.com/graph-gophers/graphql-go"
	"github.com/jehaby/webapp102/service"
	"github.com/satori/go.uuid"
)

func (r *Resolver) Ad(ctx context.Context, args struct {
	UUID graphql.ID
}) (*adResolver, error) {
	uuid, err := uuid.FromString(string(args.UUID))
	if err != nil {
		return nil, err
	}
	ad, err := r.app.Repo.Ad.GetByUUID(uuid)
	if err != nil {
		return nil, err
	}
	return &adResolver{ad}, nil
}

func (r *Resolver) Ads(ctx context.Context, args struct{ Args *service.AdsArgs }) (*adsConnectionResolver, error) {
	// TODO: validate enums (orderBy, currency)
	if args.Args == nil {
		args.Args = &service.AdsArgs{}
	}

	res, err := r.app.Service.Ad.Ads(ctx, *args.Args)
	if err != nil {
		return nil, err
	}

	return &adsConnectionResolver{
		ads:         res.Ads,
		hasNextPage: res.HasNextPage,
		orderBy:     args.Args.Order.OrderByOrDefault(),
		totalCount:  int32(res.TotalCount),
	}, nil
}

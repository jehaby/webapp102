package resolver

import (
	"context"

	"github.com/graph-gophers/graphql-go"
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

func (r *Resolver) Ads(ctx context.Context, args *struct{ Args *AdsArgs }) (*adsConnectionResolver, error) {

	// TODO: here you have to get all ads using adsArgs. You have to use AdsService somehow

	return &adsConnectionResolver{}, nil
}

type AdsArgs struct {
	First *int64
	// Name *string
}

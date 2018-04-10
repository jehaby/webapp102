package resolver

import "context"

func (r *Resolver) Ads(ctx context.Context, args *struct{ Args *AdsArgs }) (*adsConnectionResolver, error) {

	// TODO: here you have to get all ads using adsArgs. You have to use AdsService somehow

	return &adsConnectionResolver{}, nil
}

type AdsArgs struct {
	First *uint64
	// Name *string
}

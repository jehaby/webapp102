package resolver

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/jehaby/webapp102/entity"
)

type adsEdgeResolver struct {
	cursor *graphql.ID
	ad     *entity.Ad
}

func (aer *adsEdgeResolver) Cursor() *graphql.ID {
	return aer.cursor
}

func (aer *adsEdgeResolver) Node() *adResolver {
	return &adResolver{aer.ad}
}

package resolver

import (
	"github.com/jehaby/webapp102/entity"
	"github.com/jehaby/webapp102/service"
)

type adsConnectionResolver struct {
	ads         []*entity.Ad
	totalCount  int32
	hasNextPage bool
	orderBy     service.OrderBy
}

func (acr *adsConnectionResolver) TotalCount() int32 {
	return acr.totalCount
}

func (acr *adsConnectionResolver) Edges() *[]*adsEdgeResolver {
	edges := make([]*adsEdgeResolver, len(acr.ads))
	var ad *entity.Ad
	for i := range edges {
		edges[i] = &adsEdgeResolver{}
		if ad = acr.ads[i]; ad == nil {
			// log maybe?
			continue
		}
		edges[i].cursor = service.EncodeCursor(*ad, acr.orderBy)
		edges[i].ad = acr.ads[i]
	}
	return &edges
}

func (acr *adsConnectionResolver) PageInfo() *pageInfoResolver {
	res := &pageInfoResolver{
		hasNextPage: acr.hasNextPage, // TODO: next page
	}
	res.startCursor = service.EncodeCursor(*acr.ads[0], acr.orderBy)
	res.endCursor = service.EncodeCursor(*acr.ads[len(acr.ads)-1], acr.orderBy)
	return res
}

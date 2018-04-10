package resolver

import (
	"github.com/jehaby/webapp102/entity"
	"github.com/jehaby/webapp102/service"
)

type adsConnectionResolver struct {
	ads         []*entity.Ad
	totalCount  int32
	from        *string
	to          *string
	hasNextPage bool
}

func (acr *adsConnectionResolver) TotalCount() int32 {
	return acr.totalCount
}

func (acr *adsConnectionResolver) Edges() *[]*adsEdgeResolver {
	l := make([]*adsEdgeResolver, len(acr.ads))
	for i := range l {
		l[i] = &adsEdgeResolver{}
		adPtr := acr.ads[i]
		if adPtr == nil {
			continue
		}
		tmpCursor := acr.ads[i].UUID.String()
		l[i].cursor = service.EncodeCursor(&tmpCursor)
		l[i].ad = acr.ads[i]

		// TODO: try to make this shit less ugly
	}
	return &l
}

func (acr *adsConnectionResolver) PageInfo() *pageInfoResolver {
	return &pageInfoResolver{
		startCursor: service.EncodeCursor(acr.from),
		endCursor:   service.EncodeCursor(acr.to),
		hasNextPage: acr.hasNextPage, // TODO: next page
	}
}

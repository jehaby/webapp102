package resolver

import (
	"strconv"

	graphql "github.com/graph-gophers/graphql-go"

	"github.com/jehaby/webapp102/entity"
)

type localityResolver struct {
	e entity.Locality
}

func (r *localityResolver) ID() graphql.ID {
	return graphql.ID(strconv.Itoa(int(r.e.ID)))
}

func (r *localityResolver) Name() string {
	return r.e.Name
}

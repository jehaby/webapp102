package resolver

import (
	"strconv"

	"github.com/graph-gophers/graphql-go"

	"github.com/jehaby/webapp102/entity"
)

type brandResolver struct {
	e *entity.Brand
}

func (r *brandResolver) ID() graphql.ID {
	return graphql.ID(strconv.FormatUint(uint64(r.e.ID), 10))
}

func (r *brandResolver) Name() string {
	return r.e.Name
}

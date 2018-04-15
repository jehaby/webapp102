package resolver

import (
	"strconv"

	"github.com/graph-gophers/graphql-go"

	"github.com/jehaby/webapp102/entity"
)

type manufacturerResolver struct {
	e *entity.Manufacturer
}

func (r *manufacturerResolver) ID() graphql.ID {
	return graphql.ID(strconv.FormatUint(uint64(r.e.ID), 10))
}

func (r *manufacturerResolver) Name() string {
	return r.e.Name
}

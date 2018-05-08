package resolver

import (
	"strconv"

	graphql "github.com/graph-gophers/graphql-go"

	"github.com/jehaby/webapp102/entity"
)

type propertyResolver struct {
	e entity.Property
}

func (r *propertyResolver) ID() graphql.ID {
	return graphql.ID(strconv.FormatUint(uint64(r.e.ID), 10))
}

func (r *propertyResolver) Name() string {
	return r.e.Name
}

func (r *propertyResolver) Type() entity.PropertyType {
	return r.e.Type
}

func (r *propertyResolver) Values() []string {
	return r.e.Values
}

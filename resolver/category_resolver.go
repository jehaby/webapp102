package resolver

import (
	"strconv"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/jehaby/webapp102/entity"
)

type categoryResolver struct {
	e *entity.Category
}

func (r *categoryResolver) ID() graphql.ID {
	return graphql.ID(strconv.FormatUint(uint64(r.e.ID), 10))
}

func (r *categoryResolver) Name() string {
	return r.e.Name
}

func (r *categoryResolver) Path() []int32 {
	res := make([]int32, 0, len(r.e.Path))
	for i, _ := range r.e.Path {
		res = append(res, int32(r.e.Path[i]))
	}
	return res
}

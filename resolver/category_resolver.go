package resolver

import (
	"github.com/jehaby/webapp102/entity"
)

type categoryResolver struct {
	category *entity.Category
}

func (r *categoryResolver) ID() int32 {
	return int32(r.category.ID)
}

func (r *categoryResolver) Name() string {
	return r.category.Name
}

func (r *categoryResolver) Path() []int32 {
	res := make([]int32, 0, len(r.category.Path))
	for i, _ := range r.category.Path {
		res = append(res, int32(r.category.Path[i]))
	}
	return res
}

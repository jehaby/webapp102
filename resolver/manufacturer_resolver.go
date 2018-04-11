package resolver

import (
	"github.com/jehaby/webapp102/entity"
)

type manufacturerResolver struct {
	manufacturer *entity.Manufacturer
}

func (r *manufacturerResolver) Id() int32 {
	return int32(r.manufacturer.ID)
}

func (r *manufacturerResolver) Name() string {
	return r.manufacturer.Name
}

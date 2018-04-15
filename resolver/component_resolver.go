package resolver

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/graph-gophers/graphql-go"
	"github.com/jehaby/webapp102/entity"
)

func newComponentResolver(e *entity.Component) (*componentResolver, error) {
	if e == nil {
		return nil, errors.New("newComponentResolver: passed nil entity")
	}
	return &componentResolver{e}, nil
}

type componentResolver struct {
	e *entity.Component
}

func (r *componentResolver) ID() graphql.ID {
	return graphql.ID(strconv.FormatUint(uint64(r.e.ID), 10))
}

func (r *componentResolver) Name() string {
	return r.e.Name
}

func (r *componentResolver) Manufacturer() (*manufacturerResolver, error) {
	if r.e.Manufacturer == nil {
		return nil, fmt.Errorf("componentResolver.Manufacturer(): Manufacturer is nil %v", r.e)
	}
	return &manufacturerResolver{r.e.Manufacturer}, nil
}

func (r *componentResolver) Category() (*categoryResolver, error) {
	if r.e.Category == nil {
		return nil, fmt.Errorf("componentResolver.Category(): Category is nil %v", r.e)
	}
	return &categoryResolver{r.e.Category}, nil

}

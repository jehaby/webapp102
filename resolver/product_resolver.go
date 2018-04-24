package resolver

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/graph-gophers/graphql-go"
	"github.com/jehaby/webapp102/entity"
)

func newProductResolver(e *entity.Product) (*productResolver, error) {
	if e == nil {
		return nil, errors.New("newProductResolver: passed nil entity")
	}
	return &productResolver{e}, nil
}

type productResolver struct {
	e *entity.Product
}

func (r *productResolver) ID() graphql.ID {
	return graphql.ID(strconv.FormatUint(uint64(r.e.ID), 10))
}

func (r *productResolver) Name() string {
	return r.e.Name
}

func (r *productResolver) Brand() (*brandResolver, error) {
	if r.e.Brand == nil {
		return nil, fmt.Errorf("productResolver.Brand(): Brand is nil %v", r.e)
	}
	return &brandResolver{*r.e.Brand}, nil
}

func (r *productResolver) Category() (*categoryResolver, error) {
	if r.e.Category == nil {
		return nil, fmt.Errorf("productResolver.Category(): Category is nil %v", r.e)
	}
	return &categoryResolver{r.e.Category}, nil

}

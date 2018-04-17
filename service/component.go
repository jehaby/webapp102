package service

import (
	"strconv"

	"github.com/go-pg/pg"
	"github.com/pkg/errors"
	validator "gopkg.in/go-playground/validator.v9"

	"github.com/jehaby/webapp102/entity"
)

type ComponentService struct {
	db  *pg.DB
	val *validator.Validate
}

func NewComponentService(
	db *pg.DB,
	val *validator.Validate,
) *ComponentService {
	return &ComponentService{
		db:  db,
		val: val,
	}
}

func (cs *ComponentService) GetByID(id int64) (*entity.Component, error) {
	c := &entity.Component{}
	err := cs.db.Model(c).
		Relation("Manufacturer").
		Relation("Category").
		Where("component.id = ?", id).
		First()

	if err != nil {
		return nil, errors.Wrapf(err, "ComponentService.GetByID: %d", id)
	}
	return c, nil
}

type CreateComponentArgs struct {
	Name           string `validate:"required,min=2"`
	CategoryID     string `validate:"required,numeric,min=1"`
	ManufacturerID string `validate:"required,numeric,min=1"`
}

func (cs *ComponentService) Create(args CreateComponentArgs) (*entity.Component, error) {
	if err := cs.val.Struct(args); err != nil {
		return nil, err
	}

	cid, _ := strconv.ParseInt(args.CategoryID, 10, 64)
	mid, _ := strconv.ParseInt(args.ManufacturerID, 10, 64)

	e := &entity.Component{
		Name:           args.Name,
		CategoryID:     cid,
		ManufacturerID: mid,
	}

	_, err := cs.db.Model(e).Insert()
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (cs *ComponentService) Remove(id int64) error {
	// TODO: check no ads references it
	err := cs.db.Delete(&entity.Component{ID: id})
	return err
}

type UpdateComponentArgs struct {
	Name           *string `validate:"min=2"`
	CategoryID     *string `validate:"numeric,min=1"`
	ManufacturerID *string `validate:"numeric,min=1"`
}

func (cs *ComponentService) Update(id int64, args UpdateComponentArgs) (*entity.Component, error) {

	comp := &entity.Component{}
	err := cs.db.Model(comp).
		Where("component.id = ?", id).
		First()

	if err != nil {
		return nil, errors.Wrapf(err, "ComponentService.Update: %d", id)
	}

	if args.Name != nil {
		comp.Name = *args.Name
	}
	if args.CategoryID != nil {
		id, err := strconv.ParseInt(*args.CategoryID, 10, 64)
		if err != nil {
			return nil, err
		}
		comp.CategoryID = id
	}
	if args.ManufacturerID != nil {
		id, err := strconv.ParseInt(*args.ManufacturerID, 10, 64)
		if err != nil {
			return nil, err
		}
		comp.ManufacturerID = id
	}

	if err := cs.db.Update(comp); err != nil {
		return nil, err
	}

	return comp, nil
}

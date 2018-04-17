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

func (cs *ComponentService) GetByID(id uint32) (*entity.Component, error) {
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

	cid, _ := strconv.ParseUint(args.CategoryID, 10, 32)
	mid, _ := strconv.ParseUint(args.ManufacturerID, 10, 32)

	e := &entity.Component{
		Name:           args.Name,
		CategoryID:     int64(cid),
		ManufacturerID: uint16(mid),
	}

	_, err := cs.db.Model(e).Insert()
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (cs *ComponentService) Remove(id int64) error {
	// TODO: check no ads references it
	err := cs.db.Delete(&entity.Component{ID: uint32(id)})
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
		id, err := strconv.Atoi(*args.CategoryID)
		if err != nil {
			return nil, err
		}
		comp.CategoryID = int64(id)
	}
	if args.ManufacturerID != nil {
		id, err := strconv.Atoi(*args.ManufacturerID)
		if err != nil {
			return nil, err
		}
		comp.ManufacturerID = uint16(id)
	}

	if err := cs.db.Update(comp); err != nil {
		return nil, err
	}

	return comp, nil
}

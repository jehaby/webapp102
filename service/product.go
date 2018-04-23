package service

import (
	"strconv"

	"github.com/go-pg/pg"
	"github.com/pkg/errors"
	validator "gopkg.in/go-playground/validator.v9"

	"github.com/jehaby/webapp102/entity"
)

type ProductService struct {
	db  *pg.DB
	val *validator.Validate
}

func NewProductService(
	db *pg.DB,
	val *validator.Validate,
) *ProductService {
	return &ProductService{
		db:  db,
		val: val,
	}
}

func (cs *ProductService) GetByID(id int64) (*entity.Product, error) {
	c := &entity.Product{}
	err := cs.db.Model(c).
		Relation("Manufacturer").
		Relation("Category").
		Where("product.id = ?", id).
		First()

	if err != nil {
		return nil, errors.Wrapf(err, "ProductService.GetByID: %d", id)
	}
	return c, nil
}

type CreateProductArgs struct {
	Name           string `validate:"required,min=2"`
	CategoryID     string `validate:"required,numeric,min=1"`
	ManufacturerID string `validate:"required,numeric,min=1"`
}

func (cs *ProductService) Create(args CreateProductArgs) (*entity.Product, error) {
	if err := cs.val.Struct(args); err != nil {
		return nil, err
	}

	cid, _ := strconv.ParseInt(args.CategoryID, 10, 64)
	mid, _ := strconv.ParseInt(args.ManufacturerID, 10, 64)

	e := &entity.Product{
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

func (cs *ProductService) Remove(id int64) error {
	// TODO: check no ads references it
	err := cs.db.Delete(&entity.Product{ID: id})
	return err
}

type UpdateProductArgs struct {
	Name           *string `validate:"omitempty,min=2"`
	CategoryID     *string `validate:"omitempty,numeric,min=1"`
	ManufacturerID *string `validate:"omitempty,numeric,min=1"`
}

func (cs *ProductService) Update(id int64, args UpdateProductArgs) (*entity.Product, error) {
	var err error

	if err = cs.val.Struct(args); err != nil {
		return nil, err
	}

	// TODO: wherePK
	product := &entity.Product{}
	err = cs.db.Model(product).
		Where("product.id = ?", id).
		First()

	if err != nil {
		return nil, errors.Wrapf(err, "ProductService.Update: %d", id)
	}

	if args.Name != nil {
		product.Name = *args.Name
	}
	if args.CategoryID != nil {
		id, err := strconv.ParseInt(*args.CategoryID, 10, 64)
		if err != nil {
			return nil, err
		}
		product.CategoryID = id
	}
	if args.ManufacturerID != nil {
		id, err := strconv.ParseInt(*args.ManufacturerID, 10, 64)
		if err != nil {
			return nil, err
		}
		product.ManufacturerID = id
	}

	if err := cs.db.Update(product); err != nil {
		return nil, err
	}

	return product, nil
}

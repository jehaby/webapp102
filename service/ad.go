package service

import (
	"strconv"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"

	"github.com/go-pg/pg"
	"github.com/jmoiron/sqlx"
	validator "gopkg.in/go-playground/validator.v9"

	"github.com/jehaby/webapp102/entity"
	"github.com/jehaby/webapp102/storage"
)

type AdService struct {
	Repo adServiceRepos
	db   *pg.DB
	val  *validator.Validate
}

type adServiceRepos struct {
	Ad       *storage.AdRepository
	Category *storage.CategoryRepository
}

func newAdService(db *sqlx.DB, pgdb *pg.DB, val *validator.Validate) *AdService {
	return &AdService{
		Repo: adServiceRepos{
			Ad:       storage.NewAdRepository(db, pgdb),
			Category: storage.NewCategoryRepository(db),
		},
		db:  pgdb,
		val: val,
	}
}

type AdCreateArgs struct {
	Name        string          `validate:"required,min=2"`
	Description string          `validate:"required,min=5"`
	UserUUID    string          `validate:"required"` // TODO: uuid validattion?
	ProductID   string          `validate:"required,numeric,min=1"`
	LocalityID  string          `validate:"required,numeric,min=1"`
	Price       int64           `validate:"required,min=0"`
	Currency    entity.Currency `validate:"required"`
}

func (as *AdService) Create(args AdCreateArgs) (*entity.Ad, error) {
	if err := as.val.Struct(args); err != nil {
		return nil, err
	}

	// TODO: transaction!
	productID, _ := strconv.ParseInt(args.ProductID, 10, 64)
	localityID, _ := strconv.ParseInt(args.LocalityID, 10, 64)
	userUUID, err := uuid.FromString(args.UserUUID)
	if err != nil {
		return nil, errors.Wrapf(err, "AdService.Create")
	}

	product := &entity.Product{ID: productID}
	err = as.db.Model(product).Column("category_id").WherePK().First()
	if err != nil {
		return nil, errors.Wrapf(err, "TODO")
	}

	ad := &entity.Ad{
		UUID:        uuid.NewV4(),
		CreatedAt:   time.Now(),
		Name:        args.Name,
		Description: args.Description,
		CategoryID:  product.CategoryID,
		ProductID:   productID,
		LocalityID:  localityID,
		UserUUID:    userUUID,
		Price:       args.Price,
		Currency:    args.Currency,
	}

	_, err = as.db.Model(ad).Insert()
	if err != nil {
		return nil, errors.Wrapf(err, "TODO")
	}

	return ad, nil
}

type AdUpdateArgs struct {
	Name        *string `validate:"omitempty,min=2"`
	Description *string `validate:"omitempty,min=5"`
	ProductID   *string `validate:"omitempty,numeric,min=1"`
	LocalityID  *string `validate:"omitempty,numeric,min=1"`
	Price       *int64  `validate:"omitempty,min=1"`
	Currency    *entity.Currency
}

func (as *AdService) Update(uuid uuid.UUID, args AdUpdateArgs) (*entity.Ad, error) {
	var err error

	if err = as.val.Struct(args); err != nil {
		return nil, err
	}

	// TODO: transaction!
	ad := &entity.Ad{UUID: uuid}
	err = as.db.Model(ad).
		Relation("Product").
		Relation("Locality").
		WherePK().First()
	if err != nil {
		return nil, err
	}

	if args.Name != nil {
		ad.Name = *args.Name
	}
	if args.Description != nil {
		ad.Description = *args.Description
	}
	if args.ProductID != nil {
		product := &entity.Product{}
		err = as.db.Model(product).Where("product.id = ?", *args.ProductID).First()
		if err != nil {
			return nil, errors.Wrapf(err, "product with ID (%v) not found in db", *args.ProductID)
		}
		ad.ProductID = product.ID
		ad.Product = product
	}
	if args.LocalityID != nil {
		loc := &entity.Locality{}
		err = as.db.Model(loc).Where("locality.id = ?", *args.LocalityID).First()
		if err != nil {
			return nil, errors.Wrapf(err, "locality with ID (%v) not found in db", *args.LocalityID)
		}
		ad.LocalityID = loc.ID
		ad.Locality = loc
	}

	if args.Price != nil {
		ad.Price = *args.Price
	}
	if args.Currency != nil {
		ad.Currency = *args.Currency
	}

	ad.UpdatedAt = pointer.ToTime(time.Now())

	if err = as.db.Update(ad); err != nil {
		return nil, err
	}

	return ad, nil
}

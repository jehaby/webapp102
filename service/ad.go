package service

import (
	"strconv"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"

	"github.com/go-pg/pg"
	validator "gopkg.in/go-playground/validator.v9"

	"github.com/jehaby/webapp102/entity"
	"github.com/jehaby/webapp102/pkg/log"
)

type AdService struct {
	db              *pg.DB
	val             *validator.Validate
	categoryService *CategoryService
	log             *log.Logger
}

func NewAdService(
	pgdb *pg.DB,
	val *validator.Validate,
	categoryService *CategoryService,
	log *log.Logger,
) *AdService {
	return &AdService{
		db:              pgdb,
		val:             val,
		categoryService: categoryService,
		log:             log,
	}
}

type AdCreateArgs struct {
	Name        string           `validate:"required,min=2"`
	Description string           `validate:"required,min=5"`
	Condition   entity.Condition `validate:"required"`
	UserUUID    string           `validate:"required"` // TODO: uuid validattion?
	CategoryID  string           `validate:"required,numeric,min=1"`
	ProductID   *string          `validate:"omitempty,numeric,min=1"`
	LocalityID  string           `validate:"required,numeric,min=1"`
	Price       int64            `validate:"required,min=0"`
	Currency    entity.Currency  `validate:"required"`
	Weight      *int64           `validate:"omitempty,min=1"`
	BrandID     *string          `validate:"omitempty,numeric,min=1"`
	Properties  *string
}

func (as *AdService) Create(args AdCreateArgs) (*entity.Ad, error) {
	if err := as.val.Struct(args); err != nil {
		return nil, err
	}

	userUUID, err := uuid.FromString(args.UserUUID)
	if err != nil {
		return nil, errors.Wrapf(err, "AdService.Create")
	}

	// TODO: transaction!
	// productID, _ := strconv.ParseInt(args.ProductID, 10, 64)
	localityID, _ := strconv.ParseInt(args.LocalityID, 10, 64)
	categoryID, _ := strconv.ParseInt(args.CategoryID, 10, 64)

	// product := &entity.Product{ID: productID}
	// err = as.db.Model(product).Column("category_id").WherePK().First()
	// if err != nil {
	// 	return nil, errors.Wrapf(err, "TODO")
	// }

	ad := &entity.Ad{
		UUID:        uuid.NewV4(),
		CreatedAt:   time.Now(),
		Name:        args.Name,
		Description: args.Description,
		Condition:   args.Condition,
		CategoryID:  categoryID,
		// ProductID:   productID,
		LocalityID: localityID,
		UserUUID:   userUUID,
		Price:      args.Price,
		Currency:   args.Currency,
	}

	if args.Weight != nil {
		ad.Weight = *args.Weight
	}
	if args.BrandID != nil {
		bid, _ := strconv.ParseInt(*args.BrandID, 10, 64)
		ad.BrandID = bid
	}
	if args.Properties != nil {
		ad.Properties = *args.Properties
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
	Condition   *entity.Condition
	CategoryID  *string `validate:"omitempty,numeric,min=1"`
	ProductID   *string `validate:"omitempty,numeric,min=1"`
	LocalityID  *string `validate:"omitempty,numeric,min=1"`
	Price       *int64  `validate:"omitempty,min=1"`
	BrandID     *string `validate:"omitempty,min=1"`
	Weight      *int64  `validate:"omitempty,min=0"`
	Properties  *string
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
		Relation("Category").
		Relation("Locality").
		Relation("Brand").
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
	if args.Condition != nil {
		ad.Condition = *args.Condition
	}
	if args.CategoryID != nil {
		cid, _ := strconv.ParseInt(*args.CategoryID, 10, 64)
		cat := as.categoryService.GetByID(cid)
		if cat == nil {
			return nil, errors.Errorf("category with id (%d) not found", cid)
		}
		ad.CategoryID = cat.ID
		ad.Category = cat
	}
	if args.LocalityID != nil {
		// TODO: locality service, cache, or just rely on postgres FK constraint
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
	if args.Weight != nil {
		ad.Weight = *args.Weight
	}
	if args.BrandID != nil {
		brand := &entity.Brand{}
		err = as.db.Model(brand).Where("brand.id = ?", *args.BrandID).First()
		if err != nil {
			return nil, errors.Wrapf(err, "brand with ID (%v) not found in db", *args.LocalityID)
		}
		ad.BrandID = brand.ID
		ad.Brand = brand
	}
	if args.Properties != nil {
		// TODO: much more complicated logic, separate service
		// Have to check if such property-value allowed for given category; log all new property/values
		ad.Properties = *args.Properties
	}

	ad.UpdatedAt = pointer.ToTime(time.Now())

	if err = as.db.Update(ad); err != nil {
		return nil, err
	}

	return ad, nil
}

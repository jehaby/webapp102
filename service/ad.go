package service

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/go-pg/pg"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"

	"github.com/jehaby/webapp102/entity"
	"github.com/jehaby/webapp102/pkg/log"
	"github.com/jehaby/webapp102/pkg/slices"
	"github.com/jehaby/webapp102/pkg/validator"
)

type AdService struct {
	db              *pg.DB
	val             *validator.Validate
	categoryService *CategoryService
	propertyService *PropertyService
	log             *log.Logger
}

func NewAdService(
	pgdb *pg.DB,
	val *validator.Validate,
	categoryService *CategoryService,
	propertyService *PropertyService,
	log *log.Logger,
) *AdService {
	return &AdService{
		db:              pgdb,
		val:             val,
		categoryService: categoryService,
		propertyService: propertyService,
		log:             log,
	}
}

type AdCreateArgs struct {
	Name        string           `validate:"required,min=2"`
	Description string           `validate:"required,min=5"`
	Condition   entity.Condition `validate:"required"`
	CategoryID  string           `validate:"required,digital_id"`
	LocalityID  string           `validate:"required,digital_id"`
	Price       int64            `validate:"required,min=0"`
	Currency    entity.Currency  `validate:"required"`
	Weight      *int64           `validate:"omitempty,min=1"`
	BrandID     *string          `validate:"omitempty,digital_id"`
	Properties  *string
}

func (as *AdService) Create(ctx context.Context, args AdCreateArgs) (*entity.Ad, error) {
	err := as.val.Struct(args)
	if err != nil {
		return nil, err
	}

	authorizedUser := UserFromCtx(ctx)
	if authorizedUser == nil {
		return nil, ErrNotAuthorized
	}

	localityID, _ := strconv.ParseInt(args.LocalityID, 10, 64)
	categoryID, _ := strconv.ParseInt(args.CategoryID, 10, 64)

	ad := &entity.Ad{
		UUID:        uuid.NewV4(),
		CreatedAt:   time.Now(),
		Name:        args.Name,
		Description: args.Description,
		Condition:   args.Condition,
		CategoryID:  categoryID,
		LocalityID:  localityID,
		UserUUID:    authorizedUser.UUID,
		Price:       args.Price,
		Currency:    args.Currency,
	}

	if args.Weight != nil {
		ad.Weight = *args.Weight
	}
	if args.BrandID != nil {
		bid, _ := strconv.ParseInt(*args.BrandID, 10, 64)
		ad.BrandID = bid
	}
	ad.Properties, err = as.checkProperties(ctx, args.Properties, categoryID)
	if err != nil {
		return nil, err
	}

	_, err = as.db.Model(ad).Insert()
	if err != nil {
		// TODO: better errors
		return nil, errors.Wrapf(err, "TODO")
	}

	return ad, nil
}

func (as *AdService) checkProperties(ctx context.Context, rawAdProps *string, catID int64) (string, error) {
	if rawAdProps == nil {
		return "", nil
	}
	categoryProperties, err := as.propertyService.GetByCategory(ctx, catID)
	if err != nil {
		return "", errors.Wrapf(err, "checkProperties: couldn't get properties by category (%d)", catID)
		// log or fail?
	}

	// TODO: several props?
	adProps := map[string]string{}
	err = json.Unmarshal([]byte(*rawAdProps), &adProps)
	if err != nil {
		return "", errors.Wrapf(err, "checkProperties: couldn't unmarshal rawAdProps (%s)", rawAdProps)
	}

	resMap := map[string]string{}
	for _, prop := range categoryProperties {
		val, ok := adProps[prop.Name]
		if !ok {
			if prop.Required {
				// TODO: user input error (bad request)
				return "", errors.Errorf("checkProperties: property %s is required", prop.Name)
			}
			continue
		}

		if !slices.StringInSlice(prop.Values, val) {
			// TODO: more info? metric?
			as.log.Warnw("unknown property value",
				"property_id", prop.ID,
				"property_value", val,
			)
		}

		resMap[prop.Name] = val
		delete(adProps, prop.Name)
	}

	for name, val := range adProps {
		as.log.Warnw("unknown property",
			"property_name", name,
			"property_value", val,
		)
	}

	res, err := json.Marshal(resMap)
	if err != nil {
		return "", errors.Wrapf(err, "checkProperties: couldn't marshal res (%v)", resMap)
	}
	return string(res), nil
}

type AdUpdateArgs struct {
	Name        *string `validate:"omitempty,min=2"`
	Description *string `validate:"omitempty,min=5"`
	Condition   *entity.Condition
	CategoryID  *string `validate:"omitempty,digital_id"`
	LocalityID  *string `validate:"omitempty,digital_id"`
	Price       *int64  `validate:"omitempty,min=1"`
	BrandID     *string `validate:"omitempty,digital_id"`
	Weight      *int64  `validate:"omitempty,min=0"`
	Properties  *string
	Currency    *entity.Currency
}

func (as *AdService) Update(ctx context.Context, uuid uuid.UUID, args AdUpdateArgs) (*entity.Ad, error) {
	var err error

	if err = as.val.Struct(args); err != nil {
		return nil, err
	}

	// TODO: transaction!
	ad := &entity.Ad{UUID: uuid}

	as.db.RunInTransaction(func(tx *pg.Tx) error {
		err = tx.Model(ad).
			Relation("Category").
			Relation("Locality").
			Relation("Brand").
			Relation("User").
			WherePK().First()
		if err != nil {
			return checkPgNotFoundErr(err)
		}

		authorizedUser := UserFromCtx(ctx)
		if !authorizedUser.CanEdit(ad.User.UUID) {
			return ErrNotAuthorized
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
		if args.LocalityID != nil {
			// TODO: locality service, cache, or *just rely on postgres FK constraint*
			loc := &entity.Locality{}
			err = tx.Model(loc).Where("locality.id = ?", *args.LocalityID).First()
			if err != nil {
				// TODO: err msg?
				return errors.Wrapf(err, "locality with ID (%v) not found in db", *args.LocalityID)
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
			err = tx.Model(brand).Where("brand.id = ?", *args.BrandID).First()
			if err != nil {
				// TODO: err message
				return errors.Wrapf(err, "brand with ID (%v) not found in db", *args.LocalityID)
			}
			ad.Brand = brand
			ad.BrandID = brand.ID
		}
		if args.Properties != nil {
			ad.Properties, err = as.checkProperties(ctx, args.Properties, ad.CategoryID)
			if err != nil {
				return err
			}
			ad.Properties = *args.Properties
		}

		ad.UpdatedAt = pointer.ToTime(time.Now())

		if err = tx.Update(ad); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return ad, nil
}

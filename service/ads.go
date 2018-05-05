package service

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/orm"
	"github.com/satori/go.uuid"

	"github.com/jehaby/webapp102/entity"
	"github.com/jehaby/webapp102/pkg/nums"
	"github.com/pkg/errors"
)

type OrderBy string

const (
	OrderByDate   = "DATE"
	OrderByPrice  = "PRICE"
	OrderByWeight = "WEIGHT"
)

var orderByMap = map[string]string{
	OrderByDate:   "created_at", // take into account updated_at
	OrderByPrice:  "price",      // currencies
	OrderByWeight: "weight",
}

func (o OrderBy) DBColumn() (string, error) {
	res, ok := orderByMap[string(o)]
	if !ok {
		return "", errors.Errorf("wrong order by value (%s)", o)
	}
	return res, nil
}

var defaultOrderBy = OrderArg{
	OrderBy:   OrderByDate,
	Direction: DirectionDesc,
}

type Direction string

const (
	DirectionAsc  = "ASC"
	DirectionDesc = "DESC"
)

type AdsArgs struct {
	First *int32 `validate:"omitempty,min=1"`
	After *string

	Order *OrderArg

	CategoryID *int32 `validate:"omitempty,min=1"`
	LocalityID *int64 `validate:"omitempty,min=1"`

	Price  *PriceArg
	Weight *struct{ Min, Max *int32 }
	// Name *string
}

type PriceArg struct {
	Currency entity.Currency
	Min, Max *int32
}

type OrderArg struct {
	OrderBy   OrderBy
	Direction Direction
}

func (o *OrderArg) OrderByOrDefault() OrderBy {
	if o != nil {
		return o.OrderBy
	}
	return defaultOrderBy.OrderBy
}

func (o OrderArg) DBString() (string, error) {
	orderBy, err := o.OrderBy.DBColumn()
	if err != nil {
		orderBy = string(defaultOrderBy.OrderBy)
	}
	return fmt.Sprintf("%s %s", orderBy, o.Direction), err
}

const (
	defaultLimit = 100
	maxLimit     = 500
)

type AdsResult struct {
	Ads         []*entity.Ad
	TotalCount  int
	HasNextPage bool
}

func (as *AdService) Ads(ctx context.Context, args AdsArgs) (AdsResult, error) {
	res := AdsResult{}
	if err := as.val.Struct(args); err != nil {
		return res, nil
	}

	limit := int(nums.PtrToInt32OrDefault(args.First, defaultLimit))
	if limit > maxLimit {
		// TODO: user warning (error?)
		limit = maxLimit
	}

	ads := make([]*entity.Ad, 0, limit)

	query := as.db.Model(&ads)

	// TODO: allow arg (only for administrator...)
	query = query.Where("deleted_at is NULL")

	if args.CategoryID != nil {
		query = query.Where("category_id = ?", *args.CategoryID)
	}
	if args.LocalityID != nil {
		query = query.Where("locality_id = ?", args.LocalityID)
	}

	if args.Weight != nil {
		if args.Weight.Min != nil {
			query = query.Where("weight >= ?", *args.Weight.Min)
		}
		if args.Weight.Max != nil {
			query = query.Where("weight <= ?", *args.Weight.Max)
		}
	}

	if args.Price != nil {
		// TODO: real price logic (currencies)
		if args.Price.Min != nil {
			query = query.Where("price >= ?", *args.Price.Min)
		}
		if args.Price.Max != nil {
			query = query.Where("price <= ?", *args.Price.Max)
		}
	}

	order := defaultOrderBy
	if args.Order != nil {
		order = *args.Order
	}
	orderBy, err := order.DBString()
	if err != nil {
		// TODO: log error
	}
	query = query.Order(orderBy).Order(fmt.Sprintf("id %s", order.Direction))

	// might be expensive!
	res.TotalCount, err = query.Count()
	if err != nil {
		return res, errors.Wrap(err, "pg error getting count")
	}

	if args.After == nil {
		// no pagination
		err = query.Limit(int(limit)).Select()
		if err != nil {
			return res, errors.Wrap(err, "pg error getting ads without pagination")
		}

		if len(res.Ads) < res.TotalCount {
			res.HasNextPage = true
		}
		res.Ads = ads
		return res, nil
	}

	// pagination
	decCursor, err := DecodeCursor(*args.After, order)
	if err != nil {
		return res, errors.Wrap(err, "couldn't decode cursor")
	}

	err = query.WhereGroup(
		func(q *orm.Query) (*orm.Query, error) {
			q = q.Where(fmt.Sprintf("%s = ?", decCursor.field), decCursor.value).
				Where(paginationIDCondition(order.Direction, decCursor.uuid))
			return q, nil
		}).
		WhereOr(fmt.Sprintf("%s %s ?", decCursor.field, getSign(order.Direction)), decCursor.value).
		Limit(limit + 1).
		Select()

	if err != nil {
		// TODO: logging
		return res, errors.Wrap(err, "pg error getting ads with pagination")
	}

	if cnt := len(ads); cnt <= limit {
		res.Ads = ads
	} else if cnt > limit {
		res.HasNextPage = true
		res.Ads = ads[0 : cnt-1]
	}

	return res, nil
}

func paginationIDCondition(d Direction, uid uuid.UUID) string {
	return fmt.Sprintf("id %s (SELECT id FROM ads WHERE uuid = '%s')", getSign(d), uid)
}

func getSign(d Direction) string {
	if d == DirectionDesc {
		return "<"
	}
	return ">"
}

package service

import (
	"context"
	"fmt"

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

type Direction string

const (
	DirectionAsc  = "ASC"
	DirectionDesc = "DESC"
)

type AdsArgs struct {
	First *int64
	After *string
	Count *int32

	Order *OrderArg

	CategoryID *int32
	LocalityID *int64

	Price  *PriceArg
	Weight *struct{ Min, Max *int64 }

	// Name *string
}

type PriceArg struct {
	Currency entity.Currency
	Min, Max *int64
}

type OrderArg struct {
	OrderBy   OrderBy
	Direction Direction
}

func (o OrderArg) DBString() (string, error) {
	orderBy, err := o.OrderBy.DBColumn()
	if err != nil {
		orderBy = string(defaultOrderBy.OrderBy)
	}
	return fmt.Sprintf("%s %s", orderBy, o.Direction), err
}

const defaultCount = 100

var defaultOrderBy = OrderArg{
	OrderBy:   OrderByDate,
	Direction: DirectionDesc,
}

func (as *AdService) Ads(ctx context.Context, args AdsArgs) ([]*entity.Ad, error) {
	count := nums.PtrToInt32OrDefault(args.Count, defaultCount)
	ads := make([]*entity.Ad, 0, count)

	query := as.db.Model(&ads)

	if args.CategoryID != nil {
		query = query.Where("category_id = ?", *args.CategoryID)
	}
	if args.LocalityID != nil {
		query = query.Where("locality_id = ?", args.LocalityID)
	}

	if args.Weight != nil {
		if args.Weight.Min != nil {
			query = query.Where("weight >= ", *args.Weight.Min)
		}
		if args.Weight.Max != nil {
			query = query.Where("weight <= ", *args.Weight.Max)
		}
	}

	if args.Price != nil {
		// TODO: real price logic (currencies)
		if args.Price.Min != nil {
			query = query.Where("price >= ", *args.Price.Min)
		}
		if args.Price.Max != nil {
			query = query.Where("price <= ", *args.Price.Max)
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
	err = query.Order(orderBy).Limit(int(count)).Select()
	if err != nil {
		return nil, err
	}

	return ads, nil
}

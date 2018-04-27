package service

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"time"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"

	"github.com/jehaby/webapp102/entity"
)

func EncodeCursor(e entity.Ad, ob OrderBy) *graphql.ID {
	var strVal string
	switch ob {
	case OrderByDate:
		strVal = e.CreatedAt.Format(time.RFC3339Nano)
	case OrderByPrice:
		strVal = strconv.FormatInt(e.Price, 10)
	case OrderByWeight:
		strVal = strconv.FormatInt(e.Weight, 10)
	}

	rawCursor := fmt.Sprintf("%s|%s", e.UUID, strVal)
	res := graphql.ID(base64.StdEncoding.EncodeToString([]byte(rawCursor)))
	return &res
}

type decodedCursor struct {
	field string
	value interface{}
	uuid  uuid.UUID
}

func DecodeCursor(cursor string, ob OrderArg) (decodedCursor, error) {
	res := decodedCursor{}
	b, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return res, errors.Wrap(err, "couldn't decode cursor")
	}

	raw := strings.Split(string(b), "|")
	if len(raw) != 2 {
		return res, errors.Errorf("bad cursor (no '|' char)")
	}
	rawUUID, rawValue := raw[0], raw[1]
	res.uuid, err = uuid.FromString(rawUUID)
	if err != nil {
		return res, errors.Wrap(err, "bad uuid in cursor")
	}

	switch ob.OrderBy {
	case OrderByDate:
		res.field = "date"
		var err error
		res.value, err = time.Parse(time.RFC3339, rawValue)
		if err != nil {
			return res, errors.Wrap(err, "bad date in cursor")
		}
	case OrderByPrice:
		// TODO: currencies!
		res.field = "price"
		res.value, err = strconv.ParseInt(rawValue, 10, 64)
		if err != nil {
			return res, errors.Wrap(err, "bad price in cursor")
		}
	case OrderByWeight:
		res.field = "weight"
		res.value, err = strconv.ParseInt(rawValue, 10, 64)
		if err != nil {
			return res, errors.Wrap(err, "bad weight in cursor")
		}
	}
	return res, nil
}

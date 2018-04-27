package service

import (
	"fmt"
	"testing"
	"time"

	"github.com/satori/go.uuid"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDecodeCursor(t *testing.T) {

	ttime, _ := time.Parse(time.RFC3339Nano, "2018-04-26T10:31:46.572997068+03:00")

	data := []struct {
		cursor string
		ob     OrderArg
		res    decodedCursor
		err    error
	}{
		{
			cursor: "NTFkZDgwMzQtMzYxOC00NDgzLWFiMTctZGE3MmE2ZGNhOTdhfDIwMTgtMDQtMjZUMTA6MzE6NDYuNTcyOTk3MDY4KzAzOjAw",
			ob: OrderArg{
				OrderBy:   OrderByDate,
				Direction: DirectionAsc,
			},
			res: decodedCursor{
				field: "date",
				value: ttime,
				uuid:  uuid.FromStringOrNil("51dd8034-3618-4483-ab17-da72a6dca97a"),
			},
		},
		{
			cursor: "NTFkZDgwMzQtMzYxOC00NDgzLWFiMTctZGE3MmE2ZGNhOTdhfDMwMA==",
			ob: OrderArg{
				OrderBy:   OrderByPrice,
				Direction: DirectionDesc,
			},
			res: decodedCursor{
				field: "price",
				value: int64(300),
				uuid:  uuid.FromStringOrNil("51dd8034-3618-4483-ab17-da72a6dca97a"),
			},
		},
	}

	for i, tc := range data {
		Convey(fmt.Sprintf("DecodeCursor testcase #%d", i+1), t, func() {
			res, err := DecodeCursor(tc.cursor, tc.ob)
			So(err, ShouldEqual, tc.err)
			So(res.field, ShouldEqual, tc.res.field)
			So(res.value, ShouldEqual, tc.res.value)
			So(res.uuid, ShouldEqual, tc.res.uuid)
		})

	}

}

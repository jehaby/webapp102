package test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/jehaby/webapp102/entity"
)

func TestManufacturersQuery(t *testing.T) {
	Convey("Query manufacturers", t, func() {
		query := `{
			"query": "{
				manufacturers() {id name}
			}"
		}`

		res := &struct {
			Data struct {
				Manufacturers []entity.Manufacturer
			}
			Errors []graphErr
		}{}

		queryGraphql(query, res, func() {
			Convey("Response shouldn't contain error", func() {
				So(res.Errors, ShouldBeNil)
				Convey("Manufacturers list isn't empty", func() {
					So(len(res.Data.Manufacturers), ShouldBeGreaterThan, 1)
				})
			})
		})

	})
}

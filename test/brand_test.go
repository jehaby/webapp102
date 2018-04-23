package test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/jehaby/webapp102/entity"
)

func TestBrandsQuery(t *testing.T) {
	Convey("Query brands", t, func() {
		query := `{
			"query": "{
				brands() {id name}
			}"
		}`

		res := &struct {
			Data struct {
				Brands []entity.Brand
			}
			Errors []graphErr
		}{}

		queryGraphql(query, res, func() {
			Convey("Response shouldn't contain error", func() {
				So(res.Errors, ShouldBeNil)
				Convey("Brands list isn't empty", func() {
					So(len(res.Data.Brands), ShouldBeGreaterThan, 1)
				})
			})
		})

	})
}

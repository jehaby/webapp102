package test

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/jehaby/webapp102/entity"
)

type componentQueryResp struct {
	Data struct {
		Component entity.Component
	}
	Errors []graphErr
}

func TestComponentQuery(t *testing.T) {
	expected := entity.Component{
		ID:   1,
		Name: "CN-HG95",
		Manufacturer: &entity.Manufacturer{
			ID: 1,
		},
		Category: &entity.Category{
			ID: 50,
		},
	}

	Convey("Given component in database", t, func() {
		// TODO: !

		Convey("Query existing component", func() {
			res := &componentQueryResp{}
			queryGraphql(componentQuery(expected.ID), res, func() {
				Convey("Response shouldn't contain error", func() {
					So(res.Errors, ShouldBeNil)
					Convey("Response have valid data", func() {
						So(res.Data.Component.ID, ShouldEqual, expected.ID)
						So(res.Data.Component.Name, ShouldEqual, expected.Name)

						So(res.Data.Component.Manufacturer, ShouldNotBeNil)
						So(res.Data.Component.Manufacturer.ID, ShouldEqual, expected.Manufacturer.ID)

						So(res.Data.Component.Category, ShouldNotBeNil)
						So(res.Data.Component.Category.ID, ShouldEqual, expected.Category.ID)
					})
				})
			})
		})

	})

}

func componentQuery(id uint32) string {
	return fmt.Sprintf(`{
		"query": "{
			component(id: %d) {
				id
				name
				category {
					id
				}
				manufacturer {
					id
				}
			}
		}"
	}`, id)
}

package test

import (
	"fmt"
	"strconv"
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

type componentUpdateResp struct {
	Data struct {
		UpdateComponent entity.Component
	}
	Errors []graphErr
}

func TestComponentCRUD(t *testing.T) {

	Convey("Creating component successfull", t, func() {
		createRes := &struct {
			Data struct {
				CreateComponent struct {
					ID string
				}
			}
			Errors []graphErr
		}{}

		e := entity.Component{
			Name:           "test component",
			CategoryID:     1,
			ManufacturerID: 1,
		}

		res := &componentQueryResp{}

		queryGraphql(createComponentMutation(e), createRes, func() {

			id, err := strconv.ParseInt(createRes.Data.CreateComponent.ID, 10, 64)
			So(err, ShouldBeNil)
			So(id, ShouldBeGreaterThan, 0)

			// TODO: checkDatabase

			queryGraphql(componentQuery(id), res, func() {
				Convey("Quering ID should return our info", func() {
					So(res.Errors, ShouldBeNil)
					So(res.Data.Component.ID, ShouldResemble, id)
					So(res.Data.Component.Name, ShouldEqual, e.Name)
					So(res.Data.Component.Category.ID, ShouldResemble, e.CategoryID)
					So(res.Data.Component.Manufacturer.ID, ShouldResemble, e.ManufacturerID)

					e.Name = "updated component"
					e.ID = id
					res := &componentUpdateResp{}

					queryGraphql(updateComponentMutation(e), res, func() {
						Convey("Update component ", func() {

							So(res.Errors, ShouldBeNil)
							So(res.Data.UpdateComponent.ID, ShouldEqual, uint32(id))
							So(res.Data.UpdateComponent.Name, ShouldEqual, e.Name)

							// TODO: check database
						})
					})
				})
			})
		})
		// TODO: delete query

	})
}

func componentQuery(id int64) string {
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

func createComponentMutation(e entity.Component) string {
	return fmt.Sprintf(`{
		"query":"
			mutation {
				createComponent(
					input: {
						name: \"%s\",
						categoryId: \"%d\",
						manufacturerId: %d,					
					}
				) {
					id
				}
			}"
		}`, e.Name, e.CategoryID, e.ManufacturerID)
}

func updateComponentMutation(e entity.Component) string {
	return fmt.Sprintf(`{
		"query":"
			mutation {
				updateComponent(
					id: \"%d\",
					input: {
						name: \"%s\",
					}
				) {
					id
					name
				}
			}"
		}`, e.ID, e.Name)
}

func removeComponentMutation(id int64) string {
	return fmt.Sprintf(`{
		"query":"
			mutation {
				removeComponent(
					id: \"%d\",
				) {
					id
				}
			}"
		}`, id)
}

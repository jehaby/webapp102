package test

import (
	"testing"

	"github.com/satori/go.uuid"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/jehaby/webapp102/entity"
)

func TestAdQuery(t *testing.T) {

	var expected = entity.Ad{
		UUID:        uuid.FromStringOrNil("5df5b126-1fac-4fe1-a421-972ba56eb17b"),
		Name:        "cool chain",
		Description: "very very cool chain bro",
		UserUUID:    uuid.FromStringOrNil("e12087ab-23b9-4d97-8b61-e7016e4e956b"),
		Component: &entity.Component{
			Name: "CN-HG54",
		},
	}

	Convey("Check database contains ad", t, func() {
		So(true, ShouldBeTrue)
		Convey("Query ad", func() {
			query := `{
				"query": "{
					ad(uuid:\"5df5b126-1fac-4fe1-a421-972ba56eb17b\") {
						uuid
						name
						description
						user {
							uuid
							name
						}
						component {
							name
						}
					}
				}"
			}`

			res := &struct {
				Data struct {
					Ad entity.Ad
				}
				Errors []graphErr
			}{}

			queryGraphql(query, res, func() {
				Convey("Response shouldn't contain error", func() {
					So(res.Errors, ShouldBeNil)
					Convey("Response have valid data", func() {
						So(res.Data.Ad.UUID, ShouldEqual, expected.UUID)
						So(res.Data.Ad.Description, ShouldEqual, expected.Description)
						So(res.Data.Ad.User, ShouldNotBeNil)
						So(res.Data.Ad.User.UUID, ShouldEqual, expected.UserUUID)
						So(res.Data.Ad.Component.Name, ShouldEqual, expected.Component.Name)
					})
				})
			})

		})
	})
}

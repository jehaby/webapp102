package test

import (
	"fmt"
	"testing"

	"github.com/satori/go.uuid"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/jehaby/webapp102/entity"
)

type adQueryResp struct {
	Data struct {
		Ad entity.Ad
	}
	Errors []graphErr
}

func TestAdQuery(t *testing.T) {
	expected := entity.Ad{
		UUID:        uuid.FromStringOrNil("5df5b126-1fac-4fe1-a421-972ba56eb17b"),
		Name:        "cool chain",
		Description: "very very cool chain bro",
		UserUUID:    uuid.FromStringOrNil("e12087ab-23b9-4d97-8b61-e7016e4e956b"),
		Component: &entity.Component{
			Name: "CN-HG54",
		},
	}

	Convey("Check database contains ad", t, func() {

		// TODO:

		Convey("Query existing ad", func() {

			res := &adQueryResp{}
			queryGraphql(adQuery(expected.UUID), res, func() {
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

	Convey("Query non-existing ad", t, func() {
		uuid := uuid.FromStringOrNil("5df5b126-1fac-4fe1-a421-972ba56eb666")
		res := &adQueryResp{}
		queryGraphql(adQuery(uuid), res, func() {
			// So(res.Data.Ad.UUID, ShouldEqual, uuid.Nil)
			So(res.Errors, ShouldNotBeNil)
			//  TODO: more assertions
		})
	})
}

func adQuery(u uuid.UUID) string {
	return fmt.Sprintf(`{
		"query": "{
			ad(uuid:\"%s\") {
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
	}`, u)
}

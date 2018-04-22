package test

import (
	"fmt"
	"testing"
	"time"

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

func TestAdCRUD(t *testing.T) {
	Convey("Ad creation successful", t, func() {
		createRes := &struct {
			Data struct {
				AdCreate struct {
					UUID      uuid.UUID
					Name      string
					CreatedAt time.Time
				}
			}
			Errors []graphErr
		}{}

		newAd := entity.Ad{
			Name:        "new test ad",
			Description: "some description",
			ComponentID: 1,
			LocalityID:  1,
			UserUUID:    testUser.UUID,
			Price:       50000,
			Currency:    entity.CurrencyRUB,
		}

		queryGraphql(mutationAdCreate(newAd), createRes, func() {
			So(createRes.Errors, ShouldBeNil)
			So(createRes.Data.AdCreate.UUID, ShouldNotEqual, uuid.Nil)
			So(createRes.Data.AdCreate.CreatedAt.Before(time.Now()), ShouldBeTrue)

			uuid := createRes.Data.AdCreate.UUID
			res := &adQueryResp{}
			queryGraphql(adQuery(uuid), res, func() {
				Convey("Querying new ad by UUID should work fine", func() {
					So(res.Errors, ShouldBeNil)
					So(res.Data.Ad.UUID, ShouldEqual, uuid)
					So(res.Data.Ad.Name, ShouldEqual, newAd.Name)
					So(res.Data.Ad.Description, ShouldEqual, newAd.Description)
					So(res.Data.Ad.Component.ID, ShouldEqual, newAd.ComponentID)
					So(res.Data.Ad.Locality.ID, ShouldEqual, newAd.LocalityID)
					So(res.Data.Ad.Price, ShouldEqual, newAd.Price)
					So(res.Data.Ad.Currency, ShouldEqual, newAd.Currency)

					newAd.UUID = uuid
					newAd.Name = "updated ad"
					newAd.Description = "updated description"
					newAd.ComponentID = 2
					newAd.LocalityID = 2
					newAd.Price = 9999999

					res := &struct {
						Data struct {
							AdUpdate entity.Ad
						}
						Errors []graphErr
					}{}
					queryGraphql(mutationAdUpdate(newAd), res, func() {
						Convey("Ad update successful", func() {
							So(res.Errors, ShouldBeNil)
							So(res.Data.AdUpdate.UUID, ShouldEqual, newAd.UUID)
							So(res.Data.AdUpdate.Name, ShouldEqual, newAd.Name)
							So(res.Data.AdUpdate.Description, ShouldEqual, newAd.Description)
							So(res.Data.AdUpdate.Component.ID, ShouldEqual, newAd.ComponentID)
							So(res.Data.AdUpdate.Locality.ID, ShouldEqual, newAd.LocalityID)
							So(res.Data.AdUpdate.Price, ShouldEqual, newAd.Price)

							So(res.Data.AdUpdate.UpdatedAt, ShouldNotBeNil)
							So(res.Data.AdUpdate.UpdatedAt.After(res.Data.AdUpdate.CreatedAt), ShouldBeTrue)
						})
					})
				})
			})
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
				price
				currency
				user {
					uuid
				}
				component {
					id
					name
				}
				locality {
					id
					name
				}
			}
		}"
	}`, u)
}

func mutationAdCreate(ad entity.Ad) string {
	return fmt.Sprintf(`{
		"query":"
			mutation {
				adCreate(
					input: {
						name: \"%s\",
						description: \"%s\",
						userUUID: \"%s\", 
						componentId: %d,
						localityId: %d,
						price: %d,
						currency: %s,
					}
				) {
					uuid
					name
					createdAt
					price
					currency
				}
			}
		"
	}`, ad.Name, ad.Description, ad.UserUUID.String(), ad.ComponentID, ad.LocalityID, ad.Price, ad.Currency)
}

func mutationAdUpdate(ad entity.Ad) string {
	return fmt.Sprintf(`{
		"query":"
			mutation {
				adUpdate(
					uuid:  \"%s\",
					input: {
						name: \"%s\",
						description: \"%s\",
						componentId: %d,
						localityId: %d,						
						price: %d,
					}
				) {
					uuid
					name
					description
					price
					component {
						id
					}
					locality {
						id
					}
					createdAt
					updatedAt
				}
			}
		"
	}`, ad.UUID, ad.Name, ad.Description, ad.ComponentID, ad.LocalityID, ad.Price)
}

package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/satori/go.uuid"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/jehaby/webapp102/entity"
	"github.com/jehaby/webapp102/test/data"
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

		props := `{"speed": "10"}`
		newAd := entity.Ad{
			Name:        "Shimano chain wtf",
			Description: "some description",
			CategoryID:  data.Categories.Chain,
			LocalityID:  1,
			UserUUID:    data.TestUser.UUID,
			Condition:   entity.ConditionNew,
			Price:       50000,
			Currency:    entity.CurrencyRUB,
			BrandID:     data.Brands.Shimano,
			Weight:      250,
			Properties:  escapeQuotes(props),
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
					So(res.Data.Ad.Condition, ShouldEqual, newAd.Condition)
					So(res.Data.Ad.Category.ID, ShouldEqual, newAd.CategoryID)
					So(res.Data.Ad.Locality.ID, ShouldEqual, newAd.LocalityID)
					So(res.Data.Ad.Price, ShouldEqual, newAd.Price)
					So(res.Data.Ad.Currency, ShouldEqual, newAd.Currency)
					So(res.Data.Ad.Brand.ID, ShouldEqual, newAd.BrandID)
					So(res.Data.Ad.Weight, ShouldEqual, newAd.Weight)
					So(res.Data.Ad.Properties, ShouldEqual, props)

					newAd.UUID = uuid
					newAd.Name = "updated ad"
					newAd.Description = "updated description"
					newAd.Condition = entity.ConditionMalfunctioned
					newAd.CategoryID = data.Categories.Chain
					newAd.LocalityID = 2
					newAd.Price = 9999999
					newAd.Weight = 329
					newAd.BrandID = data.Brands.SRAM
					props := `{"speed": "10", "material": "steel"}`
					newAd.Properties = escapeQuotes(props)

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
							So(res.Data.AdUpdate.Condition, ShouldEqual, newAd.Condition)
							So(res.Data.AdUpdate.Category.ID, ShouldEqual, newAd.CategoryID)
							So(res.Data.AdUpdate.Locality.ID, ShouldEqual, newAd.LocalityID)
							So(res.Data.AdUpdate.Price, ShouldEqual, newAd.Price)
							So(res.Data.AdUpdate.Weight, ShouldEqual, newAd.Weight)
							So(res.Data.AdUpdate.Brand.ID, ShouldEqual, newAd.BrandID)
							So(res.Data.AdUpdate.Properties, ShouldEqual, props)

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
				condition
				price
				currency
				category {
					id
				}
				user {
					uuid
				}
				locality {
					id
					name
				}
				brand {
					id
				}
				weight
				properties
			}
		}"
	}`, u)
}

func mutationAdCreate(ad entity.Ad) string {
	return fmt.Sprintf(`{
		"query":"
			mutation ($input: AdCreateInput!) {
				adCreate(input: $input) {
					uuid
					name
					createdAt
					price
					currency
				}
			}
		",
		"variables": {
			"input": {
				"name": "%s",
				"description": "%s",
				"categoryId": "%d",
				"userUUID": "%s",
				"condition": "%s",
				"localityId": "%d",
				"price": %d,
				"currency": "%s",
				"brandId": "%d",
				"weight": %d,
				"properties": "%s"
			}
		}
	}`, ad.Name, ad.Description, ad.CategoryID, ad.UserUUID.String(), ad.Condition, ad.LocalityID, ad.Price, ad.Currency, ad.BrandID, ad.Weight, ad.Properties)
}

func mutationAdUpdate(ad entity.Ad) string {
	return fmt.Sprintf(`{
		"query":"
			mutation ($input: AdUpdateInput!) {
				adUpdate(uuid: \"%s\", input: $input) {
					uuid
					name
					description
					condition
					category {
						id
					}
					price
					locality {
						id
					}
					weight
					brand {
						id
					}
					properties
					createdAt
					updatedAt
				}
			}
		",
		"variables": {
			"input": {
				"name": "%s",
				"description": "%s",
				"condition": "%s",
				"categoryId": "%d",
				"localityId": "%d",
				"price": %d,
				"brandId": "%d",
				"weight": %d,
				"properties": "%s"
			}
		}
	}`, ad.UUID, ad.Name, ad.Description, ad.Condition, ad.CategoryID, ad.LocalityID, ad.Price, ad.BrandID, ad.Weight, ad.Properties)
}

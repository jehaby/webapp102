package test

import (
	"fmt"
	"strconv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/jehaby/webapp102/entity"
)

type productQueryResp struct {
	Data struct {
		Product entity.Product
	}
	Errors []graphErr
}

func TestProductQuery(t *testing.T) {
	expected := entity.Product{
		ID:   1,
		Name: "CN-HG95",
		Brand: &entity.Brand{
			ID: 1,
		},
		Category: &entity.Category{
			ID: 50,
		},
	}

	Convey("Given product in database", t, func() {
		// TODO: !

		Convey("Query existing product", func() {
			res := &productQueryResp{}
			queryGraphql(productQuery(expected.ID), res, func() {
				Convey("Response shouldn't contain error", func() {
					So(res.Errors, ShouldBeNil)
					Convey("Response have valid data", func() {
						So(res.Data.Product.ID, ShouldEqual, expected.ID)
						So(res.Data.Product.Name, ShouldEqual, expected.Name)

						So(res.Data.Product.Brand, ShouldNotBeNil)
						So(res.Data.Product.Brand.ID, ShouldEqual, expected.Brand.ID)

						So(res.Data.Product.Category, ShouldNotBeNil)
						So(res.Data.Product.Category.ID, ShouldEqual, expected.Category.ID)

					})
				})
			})
		})

	})

}

type productUpdateResp struct {
	Data struct {
		UpdateProduct entity.Product
	}
	Errors []graphErr
}

func TestProductCRUD(t *testing.T) {

	Convey("Product creation successful", t, func() {
		createRes := &struct {
			Data struct {
				CreateProduct struct {
					ID string
				}
			}
			Errors []graphErr
		}{}

		e := entity.Product{
			Name:           "test product",
			CategoryID:     1,
			BrandID: 1,
		}

		queryGraphql(createProductMutation(e), createRes, func() {
			So(createRes.Errors, ShouldBeNil)
			id, err := strconv.ParseInt(createRes.Data.CreateProduct.ID, 10, 64)
			So(err, ShouldBeNil)
			So(id, ShouldBeGreaterThan, 0)

			// TODO: checkDatabase

			res := &productQueryResp{}
			queryGraphql(productQuery(id), res, func() {
				Convey("Quering ID should return our info", func() {
					So(res.Errors, ShouldBeNil)
					So(res.Data.Product.ID, ShouldResemble, id)
					So(res.Data.Product.Name, ShouldEqual, e.Name)
					So(res.Data.Product.Category.ID, ShouldResemble, e.CategoryID)
					So(res.Data.Product.Brand.ID, ShouldResemble, e.BrandID)

					e.Name = "updated product"
					e.ID = id

					res := &productUpdateResp{}
					queryGraphql(updateProductMutation(e), res, func() {
						Convey("Update product ", func() {

							So(res.Errors, ShouldBeNil)
							So(res.Data.UpdateProduct.ID, ShouldEqual, uint32(id))
							So(res.Data.UpdateProduct.Name, ShouldEqual, e.Name)

							// TODO: check database
						})
					})
				})
			})
		})
		// TODO: delete query

	})
}

func productQuery(id int64) string {
	return fmt.Sprintf(`{
		"query": "{
			product(id: %d) {
				id
				name
				category {
					id
				}
				brand {
					id
				}
			}
		}"
	}`, id)
}

func createProductMutation(e entity.Product) string {
	return fmt.Sprintf(`{
		"query":"
			mutation {
				createProduct(
					input: {
						name: \"%s\",
						categoryId: \"%d\",
						brandId: %d,					
					}
				) {
					id
				}
			}"
		}`, e.Name, e.CategoryID, e.BrandID)
}

func updateProductMutation(e entity.Product) string {
	return fmt.Sprintf(`{
		"query":"
			mutation {
				updateProduct(
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

func removeProductMutation(id int64) string {
	return fmt.Sprintf(`{
		"query":"
			mutation {
				removeProduct(
					id: \"%d\",
				) {
					id
				}
			}"
		}`, id)
}

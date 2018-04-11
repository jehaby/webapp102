package test

import (
	"bytes"
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/jehaby/webapp102/entity"
)

func TestManufacturersQuery(t *testing.T) {

	Convey("Query manufacturers", t, func() {
		query := `{"query": "{manufacturers() {id name}}"}`
		resp, err := http.Post(graphqlAddr(), jsonContentType, bytes.NewBufferString(query))

		Convey("Error should be nil", func() {
			So(err, ShouldBeNil)

			Convey("HTTP status should be ok", func() {
				So(resp.StatusCode, ShouldEqual, http.StatusOK)

				res := struct {
					Data struct {
						Manufacturers []entity.Manufacturer
					}
					Errors []graphErr
				}{}

				unmarshalBody(resp.Body, &res)

				Convey("Response shouldn't contain error", func() {
					So(res.Errors, ShouldBeNil)
				})

				Convey("Manufacturers list isn't empty", func() {
					So(len(res.Data.Manufacturers), ShouldBeGreaterThan, 1)
				})

			})
		})
	})
}

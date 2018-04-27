package test

import (
	"fmt"
	"testing"

	"github.com/AlekSi/pointer"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/jehaby/webapp102/entity"
	"github.com/jehaby/webapp102/service"
)

type adsQueryResp struct {
	Data struct {
		Ads struct {
			TotalCount int
			PageInfo   struct {
				StartCursor, EndCursor string
				HasNextPage            bool
			}
			Edges []struct {
				Cursor string
				Node   entity.Ad
			}
		}
	}
	errors []graphErr
}

func TestAds(t *testing.T) {

	data := []struct {
		args       service.AdsArgs
		totalCount int
	}{
		{
			args:       service.AdsArgs{First: pointer.ToInt32(5)},
			totalCount: 6,
		},
	}

	for _, tc := range data {

		res := &adsQueryResp{}
		Convey("", t, func() {
			queryGraphql(adsQuery(tc.args), res, func() {
				Convey("Response data is correct", func() {
					So(res.errors, ShouldBeNil)
					So(len(res.Data.Ads.Edges), ShouldEqual, 5)
					// So(res.Data.Ads.TotalCount, ShouldEqual, tc.totalCount)
					So(res.Data.Ads.PageInfo.HasNextPage, ShouldEqual, true)
				})
			})

		})

	}
}

func adsQuery(a service.AdsArgs) string {
	return fmt.Sprintf(`{
		"query":"
			query ($args: AdsArgs) {
				ads(args: $args) {
					totalCount
					pageInfo {
						startCursor
						endCursor
						hasNextPage
					}
					edges {
						cursor
						node {
							uuid
						}
					}
				}
			}
		",
		"variables": {
			"args": {
				"first": %d
			}
		}
	}`, *a.First)
}

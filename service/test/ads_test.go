package test

import (
	"context"
	"testing"
	"time"

	"github.com/AlekSi/pointer"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/jehaby/webapp102/entity"
	"github.com/jehaby/webapp102/service"
	"github.com/jehaby/webapp102/test/data"
)

type tdata struct {
	desc        string
	args        service.AdsArgs
	ads         []*entity.Ad
	len         int
	totalCount  int
	hasNextPage bool
	err         error
}

func TestAdsSimple(t *testing.T) {
	data := []tdata{
		{
			desc: "count",
			args: service.AdsArgs{
				First: pointer.ToInt32(3),
			},
			len: 3,
			// toalCount: 3,
			hasNextPage: true,
		},
	}

	for _, tc := range data {
		Convey(tc.desc, t, func() {
			res, err := adsService.Ads(context.TODO(), tc.args)
			So(err, ShouldResemble, tc.err)
			So(len(res.Ads), ShouldEqual, tc.len)
			So(res.HasNextPage, ShouldEqual, tc.hasNextPage)
		})

	}
}

func TestAdsSortingByDate(t *testing.T) {
	td := []tdata{
		{
			desc: "Simple asc",
			args: service.AdsArgs{
				Order: &service.OrderArg{OrderBy: service.OrderByDate, Direction: service.DirectionAsc},
			},
		},
		{
			desc: "Filter by category asc",
			args: service.AdsArgs{
				CategoryID: pointer.ToInt32(int32(data.Categories.Chain)),
				Order:      &service.OrderArg{OrderBy: service.OrderByDate, Direction: service.DirectionAsc},
			},
		},
		{
			desc: "Filter by category desc",
			args: service.AdsArgs{
				CategoryID: pointer.ToInt32(int32(data.Categories.Chain)),
				Order:      &service.OrderArg{OrderBy: service.OrderByDate, Direction: service.DirectionDesc},
			},
		},
	}

	for _, tc := range td {
		Convey(tc.desc, t, func() {
			res, err := adsService.Ads(context.TODO(), tc.args)
			So(err, ShouldResemble, tc.err)

			previousDate, cond := time.Time{}, ShouldHappenOnOrBefore
			if tc.args.Order.Direction == service.DirectionAsc {
				previousDate, cond = time.Now().Add(time.Hour*1000000), ShouldHappenOnOrAfter
			}

			for _, ad := range res.Ads {
				So(previousDate, cond, previousDate)
				previousDate = ad.CreatedAt
			}
		})
	}

}

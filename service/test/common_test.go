package test

import (
	"os"
	"testing"

	"github.com/jehaby/webapp102/pkg/validator"
	"github.com/jehaby/webapp102/service"
	"github.com/jehaby/webapp102/test"
)

var adsService *service.AdService

func TestMain(m *testing.M) {
	adsService = service.NewAdService(
		test.GetPGDB(),
		validator.New(),
		nil,
		nil,
	)

	os.Exit(func() int {
		return m.Run()
	}())
}

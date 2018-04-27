package test

import (
	"os"
	"testing"

	validator "gopkg.in/go-playground/validator.v9"

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

package geo_test

import (
	"testing"

	"github.com/ultradns/ultradns-go-sdk/internal/testing/integration"
	"github.com/ultradns/ultradns-go-sdk/pkg/dirgroup/geo"
)

const serviceErrorString = "DirGroupGeo service configuration failed"

func TestNewSuccess(t *testing.T) {
	conf := integration.GetConfig()

	if _, err := geo.New(conf); err != nil {
		t.Fatal(err)
	}
}

func TestNewError(t *testing.T) {
	conf := integration.GetConfig()
	conf.Password = ""

	if _, err := geo.New(conf); err.Error() != "DirGroupGeo service configuration failed: Missing required parameters: [ password ]" {
		t.Fatal(err)
	}
}

func TestGetError(t *testing.T) {
	if _, err := geo.Get(nil); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

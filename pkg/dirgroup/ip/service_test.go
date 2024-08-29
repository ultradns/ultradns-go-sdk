package ip_test

import (
	"testing"

	"github.com/ultradns/ultradns-go-sdk/internal/testing/integration"
	"github.com/ultradns/ultradns-go-sdk/pkg/dirgroup/ip"
)

const serviceErrorString = "DirGroupIP service configuration failed"

func TestNewSuccess(t *testing.T) {
	conf := integration.GetConfig()

	if _, err := ip.New(conf); err != nil {
		t.Fatal(err)
	}
}

func TestNewError(t *testing.T) {
	conf := integration.GetConfig()
	conf.Password = ""

	if _, err := ip.New(conf); err.Error() != "DirGroupIP service configuration failed: Missing required parameters: [ password ]" {
		t.Fatal(err)
	}
}

func TestGetError(t *testing.T) {
	if _, err := ip.Get(nil); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

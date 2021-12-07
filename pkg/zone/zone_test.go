package zone_test

import (
	"testing"

	"github.com/ultradns/ultradns-go-sdk/pkg/test"
	"github.com/ultradns/ultradns-go-sdk/pkg/zone"
)

func TestNewSuccess(t *testing.T) {
	conf := test.GetConfig()
	_, err := zone.New(conf)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewError(t *testing.T) {
	conf := test.GetConfig()
	conf.Password = ""
	_, err := zone.New(conf)

	if err.Error() != "password is required to create a client" {
		t.Fatal(err)
	}
}

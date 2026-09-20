package webforward_test

import (
	"testing"

	"github.com/ultradns/ultradns-go-sdk/internal/testing/integration"
	"github.com/ultradns/ultradns-go-sdk/pkg/helper"
	"github.com/ultradns/ultradns-go-sdk/pkg/webforward"
)

const serviceErrorString = "WebForward service configuration failed"

func TestNewSuccess(t *testing.T) {
	if integration.TestUsername == "" {
		t.Skip("integration credentials not configured")
	}

	conf := integration.GetConfig()

	if _, err := webforward.New(conf); err != nil {
		t.Fatal(err)
	}
}

func TestGetError(t *testing.T) {
	if _, err := webforward.Get(nil); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestCreateWithConfigError(t *testing.T) {
	webForwardService := webforward.Service{}

	if _, _, err := webForwardService.Create("", &webforward.WebForward{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestReadWithConfigError(t *testing.T) {
	webForwardService := webforward.Service{}

	if _, _, err := webForwardService.Read("", ""); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestUpdateWithConfigError(t *testing.T) {
	webForwardService := webforward.Service{}

	if _, err := webForwardService.Update("", "", &webforward.WebForward{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestDeleteWithConfigError(t *testing.T) {
	webForwardService := webforward.Service{}

	if _, err := webForwardService.Delete("", ""); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestListWithConfigError(t *testing.T) {
	webForwardService := webforward.Service{}

	if _, _, err := webForwardService.List("", &helper.QueryInfo{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

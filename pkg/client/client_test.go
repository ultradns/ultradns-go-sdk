package client_test

import (
	"os"
	"testing"

	"github.com/ultradns/ultradns-go-sdk/internal/testing/integration"
	"github.com/ultradns/ultradns-go-sdk/pkg/client"
)

func TestNewClientWithCredentials(t *testing.T) {
	conf := integration.GetConfig()

	if _, err := client.NewClient(conf); err != nil {
		t.Error(err)
	}
}

func TestNewClientOneMissingParam(t *testing.T) {
	os.Unsetenv("ULTRADNS_USERNAME")
	os.Unsetenv("ULTRADNS_PASSWORD")
	os.Unsetenv("ULTRADNS_HOST_URL")
	conf := integration.GetConfig()
	conf.Username = ""
	conf.Password = ""
	conf.HostURL = ""

	if _, err := client.NewClient(conf); err == nil {
		t.Error("expected error, got nil")
	}
}

func TestNewClientTwoMissingParam(t *testing.T) {
	os.Unsetenv("ULTRADNS_USERNAME")
	os.Unsetenv("ULTRADNS_PASSWORD")
	os.Unsetenv("ULTRADNS_HOST_URL")
	conf := integration.GetConfig()
	conf.Username = ""
	conf.Password = ""
	conf.HostURL = ""

	if _, err := client.NewClient(conf); err == nil {
		t.Error("expected error, got nil")
	}
}

func TestNewClientAllMissingParam(t *testing.T) {
	os.Unsetenv("ULTRADNS_USERNAME")
	os.Unsetenv("ULTRADNS_PASSWORD")
	os.Unsetenv("ULTRADNS_HOST_URL")
	conf := client.Config{}
	if _, err := client.NewClient(conf); err == nil {
		t.Error("expected error, got nil")
	}
}

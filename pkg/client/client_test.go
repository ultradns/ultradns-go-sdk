package client_test

import (
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

func TestNewClientWithoutUsername(t *testing.T) {
	conf := integration.GetConfig()
	conf.Username = ""

	if _, err := client.NewClient(conf); err.Error() != "config validation failure: username is missing" {
		t.Error(err)
	}
}

func TestNewClientWithoutPassword(t *testing.T) {
	conf := integration.GetConfig()
	conf.Password = ""

	if _, err := client.NewClient(conf); err.Error() != "config validation failure: password is missing" {
		t.Error(err)
	}
}

func TestNewClientWithoutHost(t *testing.T) {
	conf := integration.GetConfig()
	conf.HostURL = ""

	if _, err := client.NewClient(conf); err.Error() != "config validation failure: host url is missing" {
		t.Error(err)
	}
}

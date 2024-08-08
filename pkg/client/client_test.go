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

func TestNewClientOneMissingParam(t *testing.T) {
	conf := integration.GetConfig()
	conf.Username = ""

	if _, err := client.NewClient(conf); err.Error() != "Missing required parameters: [ username ]" {
		t.Error(err)
	}
}

func TestNewClientTwoMissingParam(t *testing.T) {
	conf := integration.GetConfig()
	conf.Username = ""
	conf.Password = ""

	if _, err := client.NewClient(conf); err.Error() != "Missing required parameters: [ username, password ]" {
		t.Error(err)
	}
}

func TestNewClientAllMissingParam(t *testing.T) {
	if _, err := client.NewClient(client.Config{}); err.Error() != "Missing required parameters: [ username, password, host url ]" {
		t.Error(err)
	}
}

package client_test

import (
	"testing"

	"github.com/ultradns/ultradns-go-sdk/pkg/client"
	"github.com/ultradns/ultradns-go-sdk/pkg/test"
)

func TestNewClientWithCredentials(t *testing.T) {
	conf := test.GetConfig()
	conf.APIVersion = "v2"
	_, err := client.NewClient(conf)

	if err != nil {
		t.Error(err)
	}
}

func TestNewClientWithoutUsername(t *testing.T) {
	conf := test.GetConfig()
	conf.Username = ""
	_, err := client.NewClient(conf)

	if err.Error() != "config validation failure: username is missing" {
		t.Error(err)
	}
}

func TestNewClientWithoutPassword(t *testing.T) {
	conf := test.GetConfig()
	conf.Password = ""
	_, err := client.NewClient(conf)

	if err.Error() != "config validation failure: password is missing" {
		t.Error(err)
	}
}

func TestNewClientWithoutHost(t *testing.T) {
	conf := test.GetConfig()
	conf.HostURL = ""
	_, err := client.NewClient(conf)

	if err.Error() != "config validation failure: host url is missing" {
		t.Error(err)
	}
}

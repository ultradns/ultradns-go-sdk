package ip_test

import (
	"os"
	"testing"

	"github.com/ultradns/ultradns-go-sdk/internal/testing/integration"
	"github.com/ultradns/ultradns-go-sdk/pkg/dirgroup/ip"
	"github.com/ultradns/ultradns-go-sdk/pkg/helper"
)

const serviceErrorString = "DirGroupIP service configuration failed"

func TestNewSuccess(t *testing.T) {
	conf := integration.GetConfig()

	if _, err := ip.New(conf); err != nil {
		t.Fatal(err)
	}
}

func TestNewError(t *testing.T) {
	os.Unsetenv("ULTRADNS_USERNAME")
	os.Unsetenv("ULTRADNS_PASSWORD")
	os.Unsetenv("ULTRADNS_HOST_URL")
	conf := integration.GetConfig()
	conf.Username = ""
	conf.Password = ""
	conf.HostURL = ""

	if _, err := ip.New(conf); err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestGetError(t *testing.T) {
	if _, err := ip.Get(nil); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestCreateDirIPGroupWithConfigError(t *testing.T) {
	ipService := ip.Service{}
	if _, err := ipService.Create(&ip.DirGroupIP{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestReadDirIPGroupWithConfigError(t *testing.T) {
	ipService := ip.Service{}
	if _, _, _, err := ipService.Read(""); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestUpdateDirIPGroupWithConfigError(t *testing.T) {
	ipService := ip.Service{}
	if _, err := ipService.Update(&ip.DirGroupIP{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestPartialUpdateDirIPGroupWithConfigError(t *testing.T) {
	ipService := ip.Service{}
	if _, err := ipService.PartialUpdate(&ip.DirGroupIP{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestDeleteDirIPGroupWithConfigError(t *testing.T) {
	ipService := ip.Service{}
	if _, err := ipService.Delete(""); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestListDirIPGroupWithConfigError(t *testing.T) {
	ipService := ip.Service{}
	if _, _, err := ipService.List(&helper.QueryInfo{}, &ip.DirGroupIP{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestCreateDirIPGroupFailure(t *testing.T) {
	ipService, err := ip.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, err := ipService.Create(&ip.DirGroupIP{}); err.Error() != "Error while creating DirGroupIP: Server error Response - { code: '404', message: 'Status Code 404' }: {key: ''}" {
		t.Fatal(err)
	}
}

func TestReadDirIPGroupFailure(t *testing.T) {
	ipService, err := ip.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, _, _, err := ipService.Read(""); err.Error() != "Error while reading DirGroupIP: Server error Response - { code: '404', message: 'Status Code 404' }: {key: ''}" {
		t.Fatal(err)
	}
}

func TestUpdateDirIPGroupFailure(t *testing.T) {
	ipService, err := ip.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, err := ipService.Update(&ip.DirGroupIP{}); err.Error() != "Error while updating DirGroupIP: Server error Response - { code: '404', message: 'Status Code 404' }: {key: ''}" {
		t.Fatal(err)
	}
}

func TestPartialUpdateDirIPGroupFailure(t *testing.T) {
	ipService, err := ip.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, err := ipService.PartialUpdate(&ip.DirGroupIP{}); err.Error() != "Error while partial updating DirGroupIP: Server error Response - { code: '404', message: 'Status Code 404' }: {key: ''}" {
		t.Fatal(err)
	}
}

func TestDeleteDirIPGroupFailure(t *testing.T) {
	ipService, err := ip.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, err := ipService.Delete(""); err.Error() != "Error while deleting DirGroupIP: Server error Response - { code: '404', message: 'Status Code 404' }: {key: ''}" {
		t.Fatal(err)
	}
}

func TestListDirIPGroupFailure(t *testing.T) {
	ipService, err := ip.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, _, err := ipService.List(&helper.QueryInfo{}, &ip.DirGroupIP{}); err.Error() != "Error while listing DirGroupIP: Server error Response - { code: '404', message: 'Status Code 404' }: {key: 'accounts//dirgroups/ip'}" {
		t.Fatal(err)
	}
}

package zone_test

import (
	"os"
	"testing"

	"github.com/ultradns/ultradns-go-sdk/internal/testing/integration"
	"github.com/ultradns/ultradns-go-sdk/pkg/helper"
	"github.com/ultradns/ultradns-go-sdk/pkg/zone"
)

const serviceErrorString = "Zone service configuration failed"

func TestNewSuccess(t *testing.T) {
	conf := integration.GetConfig()

	if _, err := zone.New(conf); err != nil {
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

	if _, err := zone.New(conf); err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestGetError(t *testing.T) {
	if _, err := zone.Get(nil); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestCreateZoneWithConfigError(t *testing.T) {
	zoneService := zone.Service{}

	if _, err := zoneService.CreateZone(&zone.Zone{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestUpdateZoneWithConfigError(t *testing.T) {
	zoneService := zone.Service{}

	if _, err := zoneService.UpdateZone("", &zone.Zone{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestPartialUpdateZoneWithConfigError(t *testing.T) {
	zoneService := zone.Service{}

	if _, err := zoneService.PartialUpdateZone("", &zone.Zone{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestReadZoneWithConfigError(t *testing.T) {
	zoneService := zone.Service{}

	if _, _, err := zoneService.ReadZone(""); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestDeleteZoneWithConfigError(t *testing.T) {
	zoneService := zone.Service{}

	if _, err := zoneService.DeleteZone(""); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestListZoneWithConfigError(t *testing.T) {
	zoneService := zone.Service{}

	if _, _, err := zoneService.ListZone(&helper.QueryInfo{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestMigrateZoneAccountWithConfigError(t *testing.T) {
	zoneService := zone.Service{}

	if _, err := zoneService.MigrateZoneAccount([]string{}, "", ""); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestCreateZoneFailure(t *testing.T) {
	zoneService, err := zone.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, er := zoneService.CreateZone(&zone.Zone{}); er.Error() != "Error while creating Zone: Server error Response - { code: '55001', message: 'properties is required field.' }: {key: ''}" {
		t.Fatal(er)
	}
}

func TestUpdateZoneFailure(t *testing.T) {
	zoneService, err := zone.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, er := zoneService.UpdateZone("non-existing-zone", &zone.Zone{}); er.Error() != "Error while updating Zone: Server error Response - { code: '1801', message: 'Zone does not exist in the system.' }: {key: 'non-existing-zone'}" {
		t.Fatal(er)
	}
}

func TestPartialUpdateZoneFailure(t *testing.T) {
	zoneService, err := zone.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, er := zoneService.PartialUpdateZone("non-existing-zone", &zone.Zone{}); er.Error() != "Error while partial updating Zone: Server error Response - { code: '1801', message: 'Zone does not exist in the system.' }: {key: 'non-existing-zone'}" {
		t.Fatal(er)
	}
}

func TestReadZoneFailure(t *testing.T) {
	zoneService, err := zone.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, _, er := zoneService.ReadZone("non-existing-zone"); er.Error() != "Error while reading Zone: Server error Response - { code: '1801', message: 'Zone does not exist in the system.' }: {key: 'non-existing-zone'}" {
		t.Fatal(er)
	}
}

func TestDeleteZoneFailure(t *testing.T) {
	zoneService, err := zone.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, er := zoneService.DeleteZone("non-existing-zone"); er.Error() != "Error while deleting Zone: Server error Response - { code: '1801', message: 'Zone does not exist in the system.' }: {key: 'non-existing-zone'}" {
		t.Fatal(er)
	}
}

func TestListZoneFailure(t *testing.T) {
	zoneService, err := zone.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, _, er := zoneService.ListZone(&helper.QueryInfo{Query: "test:test"}); er.Error() != "Error while listing Zone: Server error Response - { code: '53005', message: 'Invalid input: q.test' }: {key: 'v3/zones/?&q=test:test&offset=0&cursor=&limit=100&sort=&reverse=false'}" {
		t.Fatal(er)
	}
}

func TestMigrateZoneAccountFailure(t *testing.T) {
	zoneService, err := zone.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, er := zoneService.MigrateZoneAccount([]string{"non-existing-zone"}, "a", "b"); er.Error() != "Error while migrating Zone: Server error Response - { code: '70002', message: 'Data not found.' }: {key: ''}" {
		t.Fatal(er)
	}
}

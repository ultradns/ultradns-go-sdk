package zone_test

import (
	"testing"

	"github.com/ultradns/ultradns-go-sdk/internal/testing/integration"
	"github.com/ultradns/ultradns-go-sdk/pkg/helper"
	"github.com/ultradns/ultradns-go-sdk/pkg/zone"
)

const serviceErrorString = "Zone service is not properly configured"

func TestNewSuccess(t *testing.T) {
	conf := integration.GetConfig()

	if _, err := zone.New(conf); err != nil {
		t.Fatal(err)
	}
}

func TestNewError(t *testing.T) {
	conf := integration.GetConfig()
	conf.Password = ""

	if _, err := zone.New(conf); err.Error() != "config error while creating Zone service : config validation failure: password is missing" {
		t.Fatal(err)
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

func TestCreateZoneFailure(t *testing.T) {
	zoneService, err := zone.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, er := zoneService.CreateZone(&zone.Zone{}); er.Error() != "error while creating Zone -  : error from api response - error code : 55001 - error message : properties is required field." {
		t.Fatal(er)
	}
}

func TestUpdateZoneFailure(t *testing.T) {
	zoneService, err := zone.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, er := zoneService.UpdateZone("non-existing-zone", &zone.Zone{}); er.Error() != "error while updating Zone - non-existing-zone : error from api response - error code : 1801 - error message : Zone does not exist in the system." {
		t.Fatal(er)
	}
}

func TestPartialUpdateZoneFailure(t *testing.T) {
	zoneService, err := zone.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, er := zoneService.PartialUpdateZone("non-existing-zone", &zone.Zone{}); er.Error() != "error while partial updating Zone - non-existing-zone : error from api response - error code : 1801 - error message : Zone does not exist in the system." {
		t.Fatal(er)
	}
}

func TestReadZoneFailure(t *testing.T) {
	zoneService, err := zone.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, _, er := zoneService.ReadZone("non-existing-zone"); er.Error() != "error while reading Zone - non-existing-zone : error from api response - error code : 1801 - error message : Zone does not exist in the system." {
		t.Fatal(er)
	}
}

func TestDeleteZoneFailure(t *testing.T) {
	zoneService, err := zone.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, er := zoneService.DeleteZone("non-existing-zone"); er.Error() != "error while deleting Zone - non-existing-zone : error from api response - error code : 1801 - error message : Zone does not exist in the system." {
		t.Fatal(er)
	}
}

func TestListZoneFailure(t *testing.T) {
	zoneService, err := zone.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, _, er := zoneService.ListZone(&helper.QueryInfo{Query: "test:test"}); er.Error() != "error while listing Zone : uri - v3/zones/?&q=test:test&offset=0&cursor=&limit=100&sort=&reverse=false : error from api response - error code : 53005 - error message : Invalid input: q.test" {
		t.Fatal(er)
	}
}

func TestListZoneSuccess(t *testing.T) {
	zoneService, err := zone.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	_, zoneListRes, er := zoneService.ListZone(&helper.QueryInfo{Limit: 10})

	if er != nil {
		t.Fatal(er)
	}

	if zoneListRes != nil && zoneListRes.QueryInfo.Limit != 10 {
		t.Fatalf("error while listing zones.")
	}
}

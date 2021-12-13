package zone_test

import (
	"testing"

	"github.com/ultradns/ultradns-go-sdk/internal/test"
	"github.com/ultradns/ultradns-go-sdk/pkg/helper"
	"github.com/ultradns/ultradns-go-sdk/pkg/zone"
)

const (
	primaryZoneType    = "PRIMARY"
	aliasZoneType      = "ALIAS"
	newZoneCreateType  = "NEW"
	serviceErrorString = "Zone service is not properly configured"
)

var (
	testPrimaryZoneName     = ""
	testSpecialCharZoneName = ""
	testAliasZoneName       = ""
)

func TestNewSuccess(t *testing.T) {
	conf := test.GetConfig()

	if _, err := zone.New(conf); err != nil {
		t.Fatal(err)
	}
}

func TestNewError(t *testing.T) {
	conf := test.GetConfig()
	conf.Password = ""

	if _, err := zone.New(conf); err.Error() != "config error while creating Zone service : config validation failure: password is missing" {
		t.Fatal(err)
	}
}

func TestGetSuccess(t *testing.T) {
	if _, err := zone.Get(test.TestClient); err != nil {
		t.Fatal(err)
	}
}

func TestGetError(t *testing.T) {
	if _, err := zone.Get(nil); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestCreateZoneSuccessWithPrimaryZone(t *testing.T) {
	zoneService, err := zone.Get(test.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	testPrimaryZoneName = test.GetRandomZoneName()
	zone := getPrimaryZone(testPrimaryZoneName)

	if _, er := zoneService.CreateZone(zone); er != nil {
		t.Fatal(er)
	}
}

func TestCreateZoneSuccessWithAliasZone(t *testing.T) {
	zoneService, err := zone.Get(test.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	testAliasZoneName = test.GetRandomZoneName()
	zone := getAliasZone(testAliasZoneName)

	if _, er := zoneService.CreateZone(zone); er != nil {
		t.Fatal(er)
	}
}

func TestCreateZoneWithConfigError(t *testing.T) {
	zoneService := zone.Service{}

	if _, err := zoneService.CreateZone(&zone.Zone{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestCreateZoneFailure(t *testing.T) {
	zoneService, err := zone.Get(test.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, er := zoneService.CreateZone(&zone.Zone{}); er.Error() != "error while creating zone -  : error code : 55001 - error message : properties is required field." {
		t.Fatal(er)
	}
}

func TestUpdateZoneSuccess(t *testing.T) {
	zoneService, err := zone.Get(test.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	zone := getPrimaryZone(testPrimaryZoneName)
	zone.PrimaryCreateInfo.RestrictIPList[0].SingleIP = "192.168.1.2"

	if _, er := zoneService.UpdateZone(testPrimaryZoneName, zone); er != nil {
		t.Fatal(er)
	}
}

func TestUpdateZoneWithConfigError(t *testing.T) {
	zoneService := zone.Service{}

	if _, err := zoneService.UpdateZone("", &zone.Zone{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestUpdateZoneFailure(t *testing.T) {
	zoneService, err := zone.Get(test.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, er := zoneService.UpdateZone("non-existing-zone", &zone.Zone{}); er.Error() != "error while updating zone - non-existing-zone : error code : 1801 - error message : Zone does not exist in the system." {
		t.Fatal(er)
	}
}

func TestPartialUpdateZoneSuccess(t *testing.T) {
	zoneService, err := zone.Get(test.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	zone := getPrimaryZone(testPrimaryZoneName)
	zone.PrimaryCreateInfo.RestrictIPList[0].SingleIP = "192.168.1.3"
	zone.PrimaryCreateInfo.NotifyAddresses[0].NotifyAddress = "192.168.1.13"

	if _, er := zoneService.PartialUpdateZone(testPrimaryZoneName, zone); er != nil {
		t.Fatal(er)
	}
}

func TestPartialUpdateZoneWithConfigError(t *testing.T) {
	zoneService := zone.Service{}

	if _, err := zoneService.PartialUpdateZone("", &zone.Zone{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestPartialUpdateZoneFailure(t *testing.T) {
	zoneService, err := zone.Get(test.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, er := zoneService.PartialUpdateZone("non-existing-zone", &zone.Zone{}); er.Error() != "error while partial updating zone - non-existing-zone : error code : 1801 - error message : Zone does not exist in the system." {
		t.Fatal(er)
	}
}

func TestReadZoneSuccess(t *testing.T) {
	zoneService, err := zone.Get(test.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	_, zoneRes, er := zoneService.ReadZone(testPrimaryZoneName)

	if er != nil {
		t.Fatal(er)
	}

	if zoneRes != nil && zoneRes.Properties != nil && zoneRes.Properties.Name != testPrimaryZoneName {
		t.Fatalf("zone name mismatched expected - %v : found - %v", testPrimaryZoneName, zoneRes.Properties.Name)
	}
}

func TestReadZoneWithConfigError(t *testing.T) {
	zoneService := zone.Service{}

	if _, _, err := zoneService.ReadZone(testPrimaryZoneName); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestReadZoneFailure(t *testing.T) {
	zoneService, err := zone.Get(test.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, _, er := zoneService.ReadZone("non-existing-zone"); er.Error() != "error while reading zone - non-existing-zone : error code : 1801 - error message : Zone does not exist in the system." {
		t.Fatal(er)
	}
}

func TestListZoneSuccess(t *testing.T) {
	zoneService, err := zone.Get(test.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	_, zoneListRes, er := zoneService.ListZone(&helper.QueryInfo{Query: "name:" + testPrimaryZoneName})

	if er != nil {
		t.Fatal(er)
	}

	if zoneListRes != nil && len(zoneListRes.Zones) > 0 && zoneListRes.Zones[0].Properties.Name != testPrimaryZoneName {
		t.Fatalf("zone name mismatched expected - %v : found - %v", testPrimaryZoneName, zoneListRes.Zones[0].Properties.Name)
	}
}

func TestListZoneWithConfigError(t *testing.T) {
	zoneService := zone.Service{}

	if _, _, err := zoneService.ListZone(&helper.QueryInfo{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestListZoneFailure(t *testing.T) {
	zoneService, err := zone.Get(test.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, _, er := zoneService.ListZone(&helper.QueryInfo{Query: "test:test"}); er.Error() != "error while listing zone : path and query params - v3/zones/?&q=test:test&offset=0&cursor=&limit=100&sort=&reverse=false : error code : 53005 - error message : Invalid input: q.test" {
		t.Fatal(er)
	}
}

func TestDeleteZoneSuccess(t *testing.T) {
	zoneService, err := zone.Get(test.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, er := zoneService.DeleteZone(testAliasZoneName); er != nil {
		t.Errorf("error while deleting alias zone : %s", er)
	}

	if _, er := zoneService.DeleteZone(testPrimaryZoneName); er != nil {
		t.Errorf("error while deleting primary zone : %s", er)
	}
}

func TestDeleteZoneWithConfigError(t *testing.T) {
	zoneService := zone.Service{}

	if _, err := zoneService.DeleteZone(""); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestDeleteZoneWithNonExistingZone(t *testing.T) {
	zoneService, err := zone.Get(test.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, er := zoneService.DeleteZone("non-existing-zone"); er.Error() != "error while deleting zone - non-existing-zone : error code : 1801 - error message : Zone does not exist in the system." {
		t.Fatal(er)
	}
}

func TestSpecialCharZoneCreate(t *testing.T) {
	zoneService, err := zone.Get(test.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	testSpecialCharZoneName = test.GetRandomZoneNameWithSpecialChar()
	zone := getPrimaryZone(testSpecialCharZoneName)

	if _, er := zoneService.CreateZone(zone); er != nil {
		t.Fatal(er)
	}
}

func TestSpecialCharZoneUpdate(t *testing.T) {
	zoneService, err := zone.Get(test.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	zone := getPrimaryZone(testSpecialCharZoneName)
	zone.PrimaryCreateInfo.RestrictIPList[0].SingleIP = "192.168.1.2"

	if _, er := zoneService.UpdateZone(testSpecialCharZoneName, zone); er != nil {
		t.Fatal(er)
	}
}

func TestSpecialCharZoneRead(t *testing.T) {
	zoneService, err := zone.Get(test.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	_, zoneRes, er := zoneService.ReadZone(testSpecialCharZoneName)

	if er != nil {
		t.Fatal(er)
	}

	if zoneRes != nil && zoneRes.Properties != nil && zoneRes.Properties.Name != testSpecialCharZoneName {
		t.Fatalf("zone name mismatched expected - %v : found - %v", testSpecialCharZoneName, zoneRes.Properties.Name)
	}
}

func TestSpecialCharZoneDelete(t *testing.T) {
	zoneService, err := zone.Get(test.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, er := zoneService.DeleteZone(testSpecialCharZoneName); er != nil {
		t.Errorf("error while deleting special character zone : %s", er)
	}
}

func getPrimaryZone(zoneName string) *zone.Zone {
	restrictIP := &zone.RestrictIP{
		SingleIP: "192.168.1.1",
	}
	notifyAddress := &zone.NotifyAddress{
		NotifyAddress: "192.168.1.11",
	}
	primaryZone := &zone.PrimaryZone{
		CreateType:      newZoneCreateType,
		RestrictIPList:  []*zone.RestrictIP{restrictIP},
		NotifyAddresses: []*zone.NotifyAddress{notifyAddress},
	}

	return &zone.Zone{
		Properties:        test.GetZoneProperties(zoneName, primaryZoneType),
		PrimaryCreateInfo: primaryZone,
	}
}

func getAliasZone(zoneName string) *zone.Zone {
	aliasZone := &zone.AliasZone{
		OriginalZoneName: testPrimaryZoneName,
	}

	return &zone.Zone{
		Properties:      test.GetZoneProperties(zoneName, aliasZoneType),
		AliasCreateInfo: aliasZone,
	}
}

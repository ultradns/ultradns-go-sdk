package zone_test

import (
	"testing"

	"github.com/ultradns/ultradns-go-sdk/pkg/helper"
	"github.com/ultradns/ultradns-go-sdk/pkg/test"
	"github.com/ultradns/ultradns-go-sdk/pkg/zone"
)

const (
	primaryZoneType   = "PRIMARY"
	aliasZoneType     = "ALIAS"
	newZoneCreateType = "NEW"
)

var (
	testPrimaryZoneName = ""
	testAliasZoneName   = ""
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
	_, err := zone.New(conf)

	if err.Error() != "config error while creating Zone service : config validation failure: password is missing" {
		t.Fatal(err)
	}
}

func TestGetSuccess(t *testing.T) {
	_, err := zone.Get(test.TestClient)

	if err != nil {
		t.Fatal(err)
	}
}

func TestGetError(t *testing.T) {
	_, err := zone.Get(nil)

	if err.Error() != "Zone service is not properly configured" {
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

	_, er := zoneService.CreateZone(zone)

	if er != nil {
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

	_, er := zoneService.CreateZone(zone)

	if er != nil {
		t.Fatal(er)
	}
}

func TestCreateZoneWithConfigError(t *testing.T) {
	zoneService := zone.Service{}
	_, err := zoneService.CreateZone(&zone.Zone{})

	if err.Error() != "Zone service is not properly configured" {
		t.Fatal(err)
	}
}

func TestCreateZoneFailure(t *testing.T) {
	zoneService, err := zone.Get(test.TestClient)

	if err != nil {
		t.Fatal(err)
	}
	_, er := zoneService.CreateZone(&zone.Zone{})

	if er.Error() != "error while creating zone -  : error code : 55001 - error message : properties is required field." {
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

	_, er := zoneService.UpdateZone(testPrimaryZoneName, zone)

	if er != nil {
		t.Fatal(er)
	}
}

func TestUpdateZoneWithConfigError(t *testing.T) {
	zoneService := zone.Service{}
	_, err := zoneService.UpdateZone("", &zone.Zone{})

	if err.Error() != "Zone service is not properly configured" {
		t.Fatal(err)
	}
}

func TestUpdateZoneFailure(t *testing.T) {
	zoneService, err := zone.Get(test.TestClient)

	if err != nil {
		t.Fatal(err)
	}
	_, er := zoneService.UpdateZone("non-existing-zone", &zone.Zone{})

	if er.Error() != "error while updating zone - non-existing-zone : error code : 1801 - error message : Zone does not exist in the system." {
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

	_, er := zoneService.PartialUpdateZone(testPrimaryZoneName, zone)

	if er != nil {
		t.Fatal(er)
	}
}

func TestPartialUpdateZoneWithConfigError(t *testing.T) {
	zoneService := zone.Service{}
	_, err := zoneService.PartialUpdateZone("", &zone.Zone{})

	if err.Error() != "Zone service is not properly configured" {
		t.Fatal(err)
	}
}

func TestPartialUpdateZoneFailure(t *testing.T) {
	zoneService, err := zone.Get(test.TestClient)

	if err != nil {
		t.Fatal(err)
	}
	_, er := zoneService.PartialUpdateZone("non-existing-zone", &zone.Zone{})

	if er.Error() != "error while partial updating zone - non-existing-zone : error code : 1801 - error message : Zone does not exist in the system." {
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
	_, _, err := zoneService.ReadZone(testPrimaryZoneName)

	if err.Error() != "Zone service is not properly configured" {
		t.Fatal(err)
	}
}

func TestReadZoneFailure(t *testing.T) {
	zoneService, err := zone.Get(test.TestClient)

	if err != nil {
		t.Fatal(err)
	}
	_, _, er := zoneService.ReadZone("non-existing-zone")

	if er.Error() != "error while reading zone - non-existing-zone : error code : 1801 - error message : Zone does not exist in the system." {
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

	if zoneListRes != nil && zoneListRes.ResultInfo != nil && zoneListRes.ResultInfo.ReturnedCount != 1 {
		t.Fatalf("zone returned count mismatched expected - %v : found - %v", 1, zoneListRes.ResultInfo.ReturnedCount)
	}

	if zoneListRes != nil && len(zoneListRes.Zones) > 0 && zoneListRes.Zones[0].Properties.Name != testPrimaryZoneName {
		t.Fatalf("zone name mismatched expected - %v : found - %v", testPrimaryZoneName, zoneListRes.Zones[0].Properties.Name)
	}
}

func TestListZoneWithConfigError(t *testing.T) {
	zoneService := zone.Service{}
	_, _, err := zoneService.ListZone(&helper.QueryInfo{})

	if err.Error() != "Zone service is not properly configured" {
		t.Fatal(err)
	}
}

func TestListZoneFailure(t *testing.T) {
	zoneService, err := zone.Get(test.TestClient)

	if err != nil {
		t.Fatal(err)
	}
	_, _, er := zoneService.ListZone(&helper.QueryInfo{Query: "test:test"})

	if er.Error() != "error while listing zone : path and query params - v3/zones/?&q=test:test&offset=0&cursor=&limit=100&sort=&reverse=false : error code : 53005 - error message : Invalid input: q.test" {
		t.Fatal(er)
	}
}

func TestDeleteZoneSuccess(t *testing.T) {
	zoneService, err := zone.Get(test.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	_, er := zoneService.DeleteZone(testAliasZoneName)

	if er != nil {
		t.Errorf("error while deleting alias zone : %s", er)
	}

	_, errr := zoneService.DeleteZone(testPrimaryZoneName)

	if errr != nil {
		t.Errorf("error while deleting primary zone : %s", errr)
	}
}

func TestDeleteZoneWithConfigError(t *testing.T) {
	zoneService := zone.Service{}
	_, err := zoneService.DeleteZone("")

	if err.Error() != "Zone service is not properly configured" {
		t.Fatal(err)
	}
}

func TestDeleteZoneWithNonExistingZone(t *testing.T) {
	zoneService, err := zone.Get(test.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	_, er := zoneService.DeleteZone("non-existing-zone")

	if er.Error() != "error while deleting zone - non-existing-zone : error code : 1801 - error message : Zone does not exist in the system." {
		t.Fatal(er)
	}
}

func getZoneProperties(zoneName, zoneType string) *zone.Properties {
	return &zone.Properties{
		Name:        zoneName,
		AccountName: test.TestUsername,
		Type:        zoneType,
	}
}

func getPrimaryZone(zoneName string) *zone.Zone {
	restrictIp := &zone.RestrictIP{
		SingleIP: "192.168.1.1",
	}
	notifyAddress := &zone.NotifyAddress{
		NotifyAddress: "192.168.1.11",
	}
	primaryZone := &zone.PrimaryZone{
		CreateType:      newZoneCreateType,
		RestrictIPList:  []*zone.RestrictIP{restrictIp},
		NotifyAddresses: []*zone.NotifyAddress{notifyAddress},
	}

	return &zone.Zone{
		Properties:        getZoneProperties(zoneName, primaryZoneType),
		PrimaryCreateInfo: primaryZone,
	}
}

func getAliasZone(zoneName string) *zone.Zone {
	aliasZone := &zone.AliasZone{
		OriginalZoneName: testPrimaryZoneName,
	}

	return &zone.Zone{
		Properties:      getZoneProperties(zoneName, aliasZoneType),
		AliasCreateInfo: aliasZone,
	}
}

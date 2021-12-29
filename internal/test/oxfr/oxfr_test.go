package oxfr_test

import (
	"os"
	"strings"
	"testing"

	"github.com/ultradns/ultradns-go-sdk/internal/test"
	"github.com/ultradns/ultradns-go-sdk/internal/test/oxfr"
	"github.com/ultradns/ultradns-go-sdk/pkg/zone"
)

const (
	secondaryZoneType = "SECONDARY"
)

var (
	testSondaryZoneName   = ""
	testPrimaryNameServer = os.Getenv("ULTRADNS_UNIT_TEST_NAME_SERVER")
)

func TestOxfrCreateZone(t *testing.T) {
	testSondaryZoneName = test.GetRandomZoneName()
	oxfr.CreateZone(testSondaryZoneName)
}

func TestCreateZoneSuccessWithSecondaryZone(t *testing.T) {
	zoneService, err := zone.Get(test.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	zone := getSecondaryZone()

	if _, er := zoneService.CreateZone(zone); er != nil {
		t.Fatal(er)
	}
}

func TestDeleteSecondaryZoneSuccess(t *testing.T) {
	zoneService, err := zone.Get(test.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, er := zoneService.DeleteZone(testSondaryZoneName); er != nil {
		t.Fatal(er)
	}
}

func TestOxfrDeleteZone(t *testing.T) {
	oxfr.DeleteZone(testSondaryZoneName)
}

func TestCreateZoneFailureWithSecondaryZone(t *testing.T) {
	zoneService, err := zone.Get(test.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	zone := getSecondaryZone()

	if _, er := zoneService.CreateZone(zone); !strings.Contains(er.Error(), "is not authoritative for zone") {
		t.Fatal(er)
	}
}

func getSecondaryZone() *zone.Zone {
	nameServerIP := &zone.NameServer{
		IP: testPrimaryNameServer,
	}
	nameServerIPList := &zone.NameServerIPList{
		NameServerIP1: nameServerIP,
	}

	primaryNameServer := &zone.PrimaryNameServers{
		NameServerIPList: nameServerIPList,
	}

	secondaryZone := &zone.SecondaryZone{
		PrimaryNameServers: primaryNameServer,
	}

	return &zone.Zone{
		Properties:          test.GetZoneProperties(testSondaryZoneName, secondaryZoneType),
		SecondaryCreateInfo: secondaryZone,
	}
}

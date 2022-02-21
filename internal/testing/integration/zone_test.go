package integration_test

import (
	"testing"

	"github.com/ultradns/ultradns-go-sdk/internal/testing/integration"
	"github.com/ultradns/ultradns-go-sdk/pkg/zone"
)

func TestZoneResources(t *testing.T) {
	t.Parallel()

	it := IntegrationTest{}
	primaryZoneName := integration.GetRandomZoneName()
	secondaryZoneName := integration.GetRandomSecondaryZoneName()
	aliasZoneName := integration.GetRandomZoneNameWithSpecialChar()

	t.Run("TestCreatePrimaryZone",
		func(st *testing.T) {
			it.Test = st
			it.CreatePrimaryZone(primaryZoneName)
		})
	t.Run("TestCreateSecondaryZone",
		func(st *testing.T) {
			it.Test = st
			it.CreateSecondaryZone(secondaryZoneName)
		})
	t.Run("TestCreateAliasZone",
		func(st *testing.T) {
			it.Test = st
			it.CreateAliasZone(aliasZoneName, primaryZoneName)
		})
	t.Run("TestReadPrimaryZone",
		func(st *testing.T) {
			it.Test = st
			it.ReadZone(primaryZoneName)
		})
	t.Run("TestUpdatePrimaryZone",
		func(st *testing.T) {
			it.Test = st
			it.UpdatePrimaryZone(primaryZoneName)
		})
	t.Run("TestPartialUpdatePrimaryZone",
		func(st *testing.T) {
			it.Test = st
			it.PartialUpdatePrimaryZone(primaryZoneName)
		})
	t.Run("TestDeleteAliasZone",
		func(st *testing.T) {
			it.Test = st
			it.DeleteZone(aliasZoneName)
		})
	t.Run("TestSecondaryZone",
		func(st *testing.T) {
			it.Test = st
			it.DeleteZone(secondaryZoneName)
		})
	t.Run("TestDeletePrimaryZone",
		func(st *testing.T) {
			it.Test = st
			it.DeleteZone(primaryZoneName)
		})
}

func (t *IntegrationTest) CreatePrimaryZone(zoneName string) {
	zoneData := integration.GetPrimaryZone(zoneName)
	t.CreateZone(zoneData)
}

func (t *IntegrationTest) CreateSecondaryZone(zoneName string) {
	zoneData := integration.GetSecondaryZone(zoneName)
	t.CreateZone(zoneData)
}

func (t *IntegrationTest) CreateAliasZone(alias, primary string) {
	zoneData := integration.GetAliasZone(alias, primary)
	t.CreateZone(zoneData)
}

func (t *IntegrationTest) UpdatePrimaryZone(zoneName string) {
	zoneData := integration.GetPrimaryZone(zoneName)
	zoneData.PrimaryCreateInfo.RestrictIPList[0].SingleIP = "192.168.1.2"
	t.UpdateZone(zoneName, zoneData)
}

func (t *IntegrationTest) PartialUpdatePrimaryZone(zoneName string) {
	zoneData := integration.GetPrimaryZone(zoneName)
	zoneData.PrimaryCreateInfo.RestrictIPList[0].SingleIP = "192.168.1.3"
	zoneData.PrimaryCreateInfo.NotifyAddresses[0].NotifyAddress = "192.168.1.13"
	t.PartialUpdateZone(zoneName, zoneData)
}

func (t *IntegrationTest) CreateZone(zoneData *zone.Zone) {
	zoneService, err := zone.Get(integration.TestClient)

	if err != nil {
		t.Test.Fatal(err)
	}

	if _, er := zoneService.CreateZone(zoneData); er != nil {
		t.Test.Fatal(er)
	}
}

func (t *IntegrationTest) UpdateZone(zoneName string, zoneData *zone.Zone) {
	zoneService, err := zone.Get(integration.TestClient)

	if err != nil {
		t.Test.Fatal(err)
	}

	if _, er := zoneService.UpdateZone(zoneName, zoneData); er != nil {
		t.Test.Fatal(er)
	}
}

func (t *IntegrationTest) PartialUpdateZone(zoneName string, zoneData *zone.Zone) {
	zoneService, err := zone.Get(integration.TestClient)

	if err != nil {
		t.Test.Fatal(err)
	}

	if _, er := zoneService.PartialUpdateZone(zoneName, zoneData); er != nil {
		t.Test.Fatal(er)
	}
}

func (t *IntegrationTest) ReadZone(zoneName string) {
	zoneService, err := zone.Get(integration.TestClient)

	if err != nil {
		t.Test.Fatal(err)
	}

	_, zoneRes, er := zoneService.ReadZone(zoneName)

	if er != nil {
		t.Test.Fatal(er)
	}

	if zoneRes != nil && zoneRes.Properties != nil && zoneRes.Properties.Name != zoneName {
		t.Test.Fatalf("zone name mismatched expected - %v : found - %v", zoneName, zoneRes.Properties.Name)
	}
}

func (t *IntegrationTest) DeleteZone(zoneName string) {
	zoneService, err := zone.Get(integration.TestClient)

	if err != nil {
		t.Test.Fatal(err)
	}

	if _, er := zoneService.DeleteZone(zoneName); er != nil {
		t.Test.Fatalf("unable to delete zone - %s : error - %s", zoneName, er.Error())
	}
}

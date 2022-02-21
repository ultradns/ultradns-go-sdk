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

func (it *IntegrationTest) CreatePrimaryZone(zoneName string) {
	zoneData := integration.GetPrimaryZone(zoneName)
	it.CreateZone(zoneData)
}

func (it *IntegrationTest) CreateSecondaryZone(zoneName string) {
	zoneData := integration.GetSecondaryZone(zoneName)
	it.CreateZone(zoneData)
}

func (it *IntegrationTest) CreateAliasZone(alias, primary string) {
	zoneData := integration.GetAliasZone(alias, primary)
	it.CreateZone(zoneData)
}

func (it *IntegrationTest) UpdatePrimaryZone(zoneName string) {
	zoneData := integration.GetPrimaryZone(zoneName)
	zoneData.PrimaryCreateInfo.RestrictIPList[0].SingleIP = "192.168.1.2"
	it.UpdateZone(zoneName, zoneData)
}

func (it *IntegrationTest) PartialUpdatePrimaryZone(zoneName string) {
	zoneData := integration.GetPrimaryZone(zoneName)
	zoneData.PrimaryCreateInfo.RestrictIPList[0].SingleIP = "192.168.1.3"
	zoneData.PrimaryCreateInfo.NotifyAddresses[0].NotifyAddress = "192.168.1.13"
	it.PartialUpdateZone(zoneName, zoneData)
}

func (it *IntegrationTest) CreateZone(zoneData *zone.Zone) {
	zoneService, err := zone.Get(integration.TestClient)

	if err != nil {
		it.Test.Fatal(err)
	}

	if _, er := zoneService.CreateZone(zoneData); er != nil {
		it.Test.Fatal(er)
	}
}

func (it *IntegrationTest) UpdateZone(zoneName string, zoneData *zone.Zone) {
	zoneService, err := zone.Get(integration.TestClient)

	if err != nil {
		it.Test.Fatal(err)
	}

	if _, er := zoneService.UpdateZone(zoneName, zoneData); er != nil {
		it.Test.Fatal(er)
	}
}

func (it *IntegrationTest) PartialUpdateZone(zoneName string, zoneData *zone.Zone) {
	zoneService, err := zone.Get(integration.TestClient)

	if err != nil {
		it.Test.Fatal(err)
	}

	if _, er := zoneService.PartialUpdateZone(zoneName, zoneData); er != nil {
		it.Test.Fatal(er)
	}
}

func (it *IntegrationTest) ReadZone(zoneName string) {
	zoneService, err := zone.Get(integration.TestClient)

	if err != nil {
		it.Test.Fatal(err)
	}

	_, zoneRes, er := zoneService.ReadZone(zoneName)

	if er != nil {
		it.Test.Fatal(er)
	}

	if zoneRes != nil && zoneRes.Properties != nil && zoneRes.Properties.Name != zoneName {
		it.Test.Fatalf("zone name mismatched expected - %v : found - %v", zoneName, zoneRes.Properties.Name)
	}
}

func (it *IntegrationTest) DeleteZone(zoneName string) {
	zoneService, err := zone.Get(integration.TestClient)

	if err != nil {
		it.Test.Fatal(err)
	}

	if _, er := zoneService.DeleteZone(zoneName); er != nil {
		it.Test.Fatalf("unable to delete zone - %s : error - %s", zoneName, er.Error())
	}
}

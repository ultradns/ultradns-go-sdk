package integration_test

import (
	"testing"

	"github.com/ultradns/ultradns-go-sdk/internal/testing/integration"
	"github.com/ultradns/ultradns-go-sdk/pkg/rdpool"
	"github.com/ultradns/ultradns-go-sdk/pkg/rrset"
)

func TestRDPoolResources(t *testing.T) {
	it := IntegrationTest{}

	t.Parallel()

	zoneName := integration.GetRandomZoneName()
	ownerName := integration.GetRandomString()

	t.Run("TestCreateRDPoolResourceZone",
		func(st *testing.T) {
			it.Test = st
			it.CreatePrimaryZone(zoneName)
		})
	t.Run("TestCreateRDPoolResourceTypeA",
		func(st *testing.T) {
			it.Test = st
			it.CreateRDPoolTypeA(ownerName, zoneName)
		})
	t.Run("TestUpdateRDPoolResourceTypeA",
		func(st *testing.T) {
			it.Test = st
			it.UpdateRDPoolTypeA(ownerName, zoneName)
		})
	t.Run("TestPartialUpdateRDPoolResourceTypeA",
		func(st *testing.T) {
			it.Test = st
			it.PartialUpdateRDPoolTypeA(ownerName, zoneName)
		})
	t.Run("TestReadRDPoolResourceTypeA",
		func(st *testing.T) {
			it.Test = st
			it.ReadRDPool(integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA))
		})
	t.Run("TestDeleteRDPoolResourceTypeA",
		func(st *testing.T) {
			it.Test = st
			it.DeleteRDPool(integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA))
		})
	t.Run("TestDeleteRDPoolResourceZone",
		func(st *testing.T) {
			it.Test = st
			it.DeleteZone(zoneName)
		})
}

func (it *IntegrationTest) CreateRDPoolTypeA(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA)
	rrSet := getRDPoolTypeA(ownerName)
	it.CreateRDPool(rrSetKey, rrSet)
}

func (it *IntegrationTest) UpdateRDPoolTypeA(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA)
	rrSet := getRDPoolTypeA(ownerName)
	rrSet.RData = []string{"192.168.1.11"}
	it.UpdateRDPool(rrSetKey, rrSet)
}

func (it *IntegrationTest) PartialUpdateRDPoolTypeA(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA)
	rrSet := getRDPoolTypeA(ownerName)
	rrSet.RData = []string{"192.168.1.12"}
	it.PartialUpdateRDPool(rrSetKey, rrSet)
}

func (it *IntegrationTest) CreateRDPool(rrSetKey *rrset.RRSetKey, rrSet *rrset.RRSet) {
	rdPoolService, err := rdpool.Get(integration.TestClient)

	if err != nil {
		it.Test.Fatal(err)
	}

	if _, er := rdPoolService.CreateRDPool(rrSetKey, rrSet); er != nil {
		it.Test.Fatal(er)
	}
}

func (it *IntegrationTest) UpdateRDPool(rrSetKey *rrset.RRSetKey, rrSet *rrset.RRSet) {
	rdPoolService, err := rdpool.Get(integration.TestClient)

	if err != nil {
		it.Test.Fatal(err)
	}

	if _, er := rdPoolService.UpdateRDPool(rrSetKey, rrSet); er != nil {
		it.Test.Fatal(er)
	}
}

func (it *IntegrationTest) PartialUpdateRDPool(rrSetKey *rrset.RRSetKey, rrSet *rrset.RRSet) {
	rdPoolService, err := rdpool.Get(integration.TestClient)

	if err != nil {
		it.Test.Fatal(err)
	}

	if _, er := rdPoolService.PartialUpdateRDPool(rrSetKey, rrSet); er != nil {
		it.Test.Fatal(er)
	}
}

func (it *IntegrationTest) ReadRDPool(rrSetKey *rrset.RRSetKey) {
	rdPoolService, err := rdpool.Get(integration.TestClient)

	if err != nil {
		it.Test.Fatal(err)
	}

	if _, _, er := rdPoolService.ReadRDPool(rrSetKey); er != nil {
		it.Test.Fatal(er)
	}
}

func (it *IntegrationTest) DeleteRDPool(rrSetKey *rrset.RRSetKey) {
	rdPoolService, err := rdpool.Get(integration.TestClient)

	if err != nil {
		it.Test.Fatal(err)
	}

	if _, er := rdPoolService.DeleteRDPool(rrSetKey); er != nil {
		it.Test.Fatal(er)
	}
}

func getRDPoolTypeA(ownerName string) *rrset.RRSet {
	profile := &rdpool.Profile{
		Order: "FIXED",
	}

	return &rrset.RRSet{
		OwnerName: ownerName,
		RRType:    testRecordTypeA,
		RData:     []string{"192.168.1.1"},
		Profile:   profile,
	}
}

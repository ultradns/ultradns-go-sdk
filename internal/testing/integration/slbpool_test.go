package integration_test

import (
	"testing"

	"github.com/ultradns/ultradns-go-sdk/internal/testing/integration"
	"github.com/ultradns/ultradns-go-sdk/pkg/record"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/pool"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/slbpool"
	"github.com/ultradns/ultradns-go-sdk/pkg/rrset"
)

func TestSLBPoolResources(t *testing.T) {
	ownerName := integration.GetRandomString()
	it := IntegrationTest{}

	t.Parallel()

	zoneName := integration.GetRandomZoneName()

	t.Run("TestCreateSLBPoolResourceZone",
		func(st *testing.T) {
			it.Test = st
			it.CreatePrimaryZone(zoneName)
		})
	t.Run("TestCreateSLBPoolResourceTypeA",
		func(st *testing.T) {
			it.Test = st
			it.CreateSLBPoolTypeA(ownerName, zoneName)
		})
	t.Run("TestUpdateSLBPoolResourceTypeA",
		func(st *testing.T) {
			it.Test = st
			it.UpdateSLBPoolTypeA(ownerName, zoneName)
		})
	t.Run("TestPartialUpdateSLBPoolResourceTypeA",
		func(st *testing.T) {
			it.Test = st
			it.PartialUpdateSLBPoolTypeA(ownerName, zoneName)
		})
	t.Run("TestReadSLBPoolResourceTypeA",
		func(st *testing.T) {
			it.Test = st
			it.ReadSLBPool(integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA))
		})
	t.Run("TestDeleteSLBPoolResourceTypeA",
		func(st *testing.T) {
			it.Test = st
			it.DeleteSLBPool(integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA))
		})
	t.Run("TestDeleteSLBPoolResourceZone",
		func(st *testing.T) {
			it.Test = st
			it.DeleteZone(zoneName)
		})
}

func (it *IntegrationTest) CreateSLBPoolTypeA(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA)
	rrSet := getSLBPoolTypeA(ownerName)
	it.CreateSLBPool(rrSetKey, rrSet)
}

func (it *IntegrationTest) UpdateSLBPoolTypeA(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA)
	rrSet := getSLBPoolTypeA(ownerName)
	rrSet.RData = []string{"192.168.1.11"}
	it.UpdateSLBPool(rrSetKey, rrSet)
}

func (it *IntegrationTest) PartialUpdateSLBPoolTypeA(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA)
	rrSet := getSLBPoolTypeA(ownerName)
	rrSet.RData = []string{"192.168.1.12"}
	it.PartialUpdateSLBPool(rrSetKey, rrSet)
}

func (it *IntegrationTest) CreateSLBPool(rrSetKey *rrset.RRSetKey, rrSet *rrset.RRSet) {
	recordService, err := record.Get(integration.TestClient)

	if err != nil {
		it.Test.Fatal(err)
	}

	if _, er := recordService.Create(rrSetKey, rrSet); er != nil {
		it.Test.Fatal(er)
	}
}

func (it *IntegrationTest) UpdateSLBPool(rrSetKey *rrset.RRSetKey, rrSet *rrset.RRSet) {
	recordService, err := record.Get(integration.TestClient)

	if err != nil {
		it.Test.Fatal(err)
	}

	if _, er := recordService.Update(rrSetKey, rrSet); er != nil {
		it.Test.Fatal(er)
	}
}

func (it *IntegrationTest) PartialUpdateSLBPool(rrSetKey *rrset.RRSetKey, rrSet *rrset.RRSet) {
	recordService, err := record.Get(integration.TestClient)

	if err != nil {
		it.Test.Fatal(err)
	}

	if _, er := recordService.PartialUpdate(rrSetKey, rrSet); er != nil {
		it.Test.Fatal(er)
	}
}

func (it *IntegrationTest) ReadSLBPool(rrSetKey *rrset.RRSetKey) {
	recordService, err := record.Get(integration.TestClient)

	if err != nil {
		it.Test.Fatal(err)
	}

	if _, _, er := recordService.Read(rrSetKey); er != nil {
		it.Test.Fatal(er)
	}
}

func (it *IntegrationTest) DeleteSLBPool(rrSetKey *rrset.RRSetKey) {
	recordService, err := record.Get(integration.TestClient)

	if err != nil {
		it.Test.Fatal(err)
	}

	if _, er := recordService.Delete(rrSetKey); er != nil {
		it.Test.Fatal(er)
	}
}

func getSLBPoolTypeA(ownerName string) *rrset.RRSet {
	rdataInfo := &slbpool.RDataInfo{
		ProbingEnabled: false,
	}
	allFailRecord := &slbpool.AllFailRecord{
		RData:   "192.168.0.1",
		Serving: false,
	}
	monitor := &pool.Monitor{
		Method: "GET",
		URL:    integration.TestHost,
	}
	profile := &slbpool.Profile{
		RDataInfo:                []*slbpool.RDataInfo{rdataInfo},
		Monitor:                  monitor,
		AllFailRecord:            allFailRecord,
		RegionFailureSensitivity: "HIGH",
		ServingPreference:        "AUTO_SELECT",
		ResponseMethod:           "ROUND_ROBIN",
	}

	return &rrset.RRSet{
		OwnerName: ownerName,
		RRType:    testRecordTypeA,
		RData:     []string{"192.168.1.1"},
		Profile:   profile,
	}
}

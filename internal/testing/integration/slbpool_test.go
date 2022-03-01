package integration_test

import (
	"testing"

	"github.com/ultradns/ultradns-go-sdk/internal/testing/integration"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/pool"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/slbpool"
	"github.com/ultradns/ultradns-go-sdk/pkg/rrset"
)

func (t *IntegrationTest) TestSLBPoolResources(zoneName string) {
	it := IntegrationTest{}
	ownerName := integration.GetRandomString()

	t.Test.Run("TestCreateSLBPoolResourceTypeA",
		func(st *testing.T) {
			it.Test = st
			it.CreateSLBPoolTypeA(ownerName, zoneName)
		})
	t.Test.Run("TestUpdateSLBPoolResourceTypeA",
		func(st *testing.T) {
			it.Test = st
			it.UpdateSLBPoolTypeA(ownerName, zoneName)
		})
	t.Test.Run("TestPartialUpdateSLBPoolResourceTypeA",
		func(st *testing.T) {
			it.Test = st
			it.PartialUpdateSLBPoolTypeA(ownerName, zoneName)
		})
	t.Test.Run("TestReadSLBPoolResourceTypeA",
		func(st *testing.T) {
			it.Test = st
			it.ReadRecord(integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, pool.SLB))
		})
	t.Test.Run("TestDeleteSLBPoolResourceTypeA",
		func(st *testing.T) {
			it.Test = st
			it.DeleteRecord(integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, ""))
		})
}

func (t *IntegrationTest) CreateSLBPoolTypeA(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, "")
	rrSet := getSLBPoolTypeA(ownerName)
	t.CreateRecord(rrSetKey, rrSet)
}

func (t *IntegrationTest) UpdateSLBPoolTypeA(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, "")
	rrSet := getSLBPoolTypeA(ownerName)
	rrSet.RData = []string{"192.168.1.11"}
	t.UpdateRecord(rrSetKey, rrSet)
}

func (t *IntegrationTest) PartialUpdateSLBPoolTypeA(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, "")
	rrSet := getSLBPoolTypeA(ownerName)
	rrSet.RData = []string{"192.168.1.12"}
	t.PartialUpdateRecord(rrSetKey, rrSet)
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

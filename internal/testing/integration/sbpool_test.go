package integration_test

import (
	"testing"

	"github.com/ultradns/ultradns-go-sdk/internal/testing/integration"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/pool"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/sbpool"
	"github.com/ultradns/ultradns-go-sdk/pkg/rrset"
)

func (t *IntegrationTest) TestSBPoolResources(zoneName, ownerName string) {
	it := IntegrationTest{}

	t.Test.Run("TestCreateSBPoolResourceTypeA",
		func(st *testing.T) {
			it.Test = st
			it.CreateSBPoolTypeA(ownerName, zoneName)
		})
	t.Test.Run("TestUpdateSBPoolResourceTypeA",
		func(st *testing.T) {
			it.Test = st
			it.UpdateSBPoolTypeA(ownerName, zoneName)
		})
	t.Test.Run("TestPartialUpdateSBPoolResourceTypeA",
		func(st *testing.T) {
			it.Test = st
			it.PartialUpdateSBPoolTypeA(ownerName, zoneName)
		})
	t.Test.Run("TestReadSBPoolResourceTypeA",
		func(st *testing.T) {
			it.Test = st
			it.ReadRecord(integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, pool.SB))
		})
}

func (it *IntegrationTest) CreateSBPoolTypeA(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, "")
	rrSet := getSBPoolTypeA(ownerName)
	it.CreateRecord(rrSetKey, rrSet)
}

func (it *IntegrationTest) UpdateSBPoolTypeA(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, "")
	rrSet := getSBPoolTypeA(ownerName)
	rrSet.RData = []string{"192.168.1.11"}
	it.UpdateRecord(rrSetKey, rrSet)
}

func (it *IntegrationTest) PartialUpdateSBPoolTypeA(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, "")
	rrSet := getSBPoolTypeA(ownerName)
	rrSet.RData = []string{"192.168.1.12"}
	it.PartialUpdateRecord(rrSetKey, rrSet)
}

func getSBPoolTypeA(ownerName string) *rrset.RRSet {
	rdataInfo := &pool.RDataInfo{
		State:         "NORMAL",
		RunProbes:     true,
		Priority:      1,
		FailoverDelay: 0,
		Threshold:     1,
	}
	profile := &sbpool.Profile{
		RDataInfo:        []*pool.RDataInfo{rdataInfo},
		RunProbes:        true,
		ActOnProbes:      true,
		FailureThreshold: 0,
		Order:            "FIXED",
		MaxActive:        1,
		MaxServed:        1,
	}

	return &rrset.RRSet{
		OwnerName: ownerName,
		RRType:    testRecordTypeA,
		RData:     []string{"192.168.1.1"},
		Profile:   profile,
	}
}

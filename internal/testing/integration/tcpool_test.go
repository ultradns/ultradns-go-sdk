package integration_test

import (
	"testing"

	"github.com/ultradns/ultradns-go-sdk/internal/testing/integration"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/pool"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/tcpool"
	"github.com/ultradns/ultradns-go-sdk/pkg/rrset"
)

func (t *IntegrationTest) TestTCPoolResources(zoneName, ownerName string) {
	it := IntegrationTest{}

	t.Test.Run("TestCreateTCPoolResourceTypeA",
		func(st *testing.T) {
			it.Test = st
			it.CreateTCPoolTypeA(ownerName, zoneName)
		})
	t.Test.Run("TestUpdateTCPoolResourceTypeA",
		func(st *testing.T) {
			it.Test = st
			it.UpdateTCPoolTypeA(ownerName, zoneName)
		})
	t.Test.Run("TestPartialUpdateTCPoolResourceTypeA",
		func(st *testing.T) {
			it.Test = st
			it.PartialUpdateTCPoolTypeA(ownerName, zoneName)
		})
	t.Test.Run("TestReadTCPoolResourceTypeA",
		func(st *testing.T) {
			it.Test = st
			it.ReadRecord(integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, pool.TC))
		})
}

func (t *IntegrationTest) CreateTCPoolTypeA(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, "")
	rrSet := getTCPoolTypeA(ownerName)
	t.CreateRecord(rrSetKey, rrSet)
}

func (t *IntegrationTest) UpdateTCPoolTypeA(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, "")
	rrSet := getTCPoolTypeA(ownerName)
	rrSet.RData = []string{"192.168.1.11"}
	t.UpdateRecord(rrSetKey, rrSet)
}

func (t *IntegrationTest) PartialUpdateTCPoolTypeA(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, "")
	rrSet := getTCPoolTypeA(ownerName)
	rrSet.RData = []string{"192.168.1.12"}
	t.PartialUpdateRecord(rrSetKey, rrSet)
}

func getTCPoolTypeA(ownerName string) *rrset.RRSet {
	rdataInfo := &pool.RDataInfo{
		State:         "NORMAL",
		RunProbes:     true,
		Priority:      1,
		FailoverDelay: 0,
		Threshold:     1,
		Weight:        2,
	}
	profile := &tcpool.Profile{
		RDataInfo:        []*pool.RDataInfo{rdataInfo},
		RunProbes:        true,
		ActOnProbes:      true,
		MaxToLB:          1,
		FailureThreshold: 0,
	}

	return &rrset.RRSet{
		OwnerName: ownerName,
		RRType:    testRecordTypeA,
		RData:     []string{"192.168.1.1"},
		Profile:   profile,
	}
}

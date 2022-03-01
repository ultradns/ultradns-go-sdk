package integration_test

import (
	"testing"

	"github.com/ultradns/ultradns-go-sdk/internal/testing/integration"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/dirpool"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/pool"
	"github.com/ultradns/ultradns-go-sdk/pkg/rrset"
)

func (t *IntegrationTest) TestDIRPoolResources(zoneName string) {
	it := IntegrationTest{}
	ownerName := integration.GetRandomString()

	t.Test.Run("TestCreateDIRPoolResourceTypeA",
		func(st *testing.T) {
			it.Test = st
			it.CreateDIRPoolTypeA(ownerName, zoneName)
		})
	t.Test.Run("TestUpdateDIRPoolResourceTypeA",
		func(st *testing.T) {
			it.Test = st
			it.UpdateDIRPoolTypeA(ownerName, zoneName)
		})
	t.Test.Run("TestPartialUpdateDIRPoolResourceTypeA",
		func(st *testing.T) {
			it.Test = st
			it.PartialUpdateDIRPoolTypeA(ownerName, zoneName)
		})
	t.Test.Run("TestReadDIRPoolResourceTypeA",
		func(st *testing.T) {
			it.Test = st
			it.ReadRecord(integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, pool.DIR))
		})
	t.Test.Run("TestDeleteDIRPoolResourceTypeA",
		func(st *testing.T) {
			it.Test = st
			it.DeleteRecord(integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, ""))
		})
}

func (t *IntegrationTest) CreateDIRPoolTypeA(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, "")
	rrSet := getDIRPoolTypeA(ownerName)
	t.CreateRecord(rrSetKey, rrSet)
}

func (t *IntegrationTest) UpdateDIRPoolTypeA(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, "")
	rrSet := getDIRPoolTypeA(ownerName)
	rrSet.RData = []string{"192.168.1.11"}
	t.UpdateRecord(rrSetKey, rrSet)
}

func (t *IntegrationTest) PartialUpdateDIRPoolTypeA(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, "")
	rrSet := getDIRPoolTypeA(ownerName)
	geoInfo := &dirpool.GEOInfo{
		Name: "GEOInfoPool",
	}
	rdataInfo := &dirpool.RDataInfo{
		GeoInfo: geoInfo,
	}
	rrSet.Profile.(*dirpool.Profile).RDataInfo = []*dirpool.RDataInfo{rdataInfo}
	rrSet.RData = []string{"192.168.1.11"}
	t.PartialUpdateRecord(rrSetKey, rrSet)
}

func getDIRPoolTypeA(ownerName string) *rrset.RRSet {
	rdataInfo := &dirpool.RDataInfo{
		AllNonConfigured: true,
	}
	profile := &dirpool.Profile{
		RDataInfo: []*dirpool.RDataInfo{rdataInfo},
	}

	return &rrset.RRSet{
		OwnerName: ownerName,
		RRType:    testRecordTypeA,
		RData:     []string{"192.168.1.1"},
		Profile:   profile,
	}
}

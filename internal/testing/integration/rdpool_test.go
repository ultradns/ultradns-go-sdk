package integration_test

import (
	"fmt"
	"testing"

	"github.com/ultradns/ultradns-go-sdk/internal/testing/integration"
	"github.com/ultradns/ultradns-go-sdk/pkg/record"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/pool"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/rdpool"
	"github.com/ultradns/ultradns-go-sdk/pkg/rrset"
)

func (t *IntegrationTest) TestRDPoolResources(zoneName string) {
	it := IntegrationTest{}
	ownerName := integration.GetRandomString()

	t.Test.Run("TestCreateRDPoolResourceTypeA",
		func(st *testing.T) {
			it.Test = st
			it.CreateRDPoolTypeA(ownerName, zoneName)
		})
	t.Test.Run("TestUpdateRDPoolResourceTypeA",
		func(st *testing.T) {
			it.Test = st
			it.UpdateRDPoolTypeA(ownerName, zoneName)
		})
	t.Test.Run("TestPartialUpdateRDPoolResourceTypeA",
		func(st *testing.T) {
			it.Test = st
			it.PartialUpdateRDPoolTypeA(ownerName, zoneName)
		})
	t.Test.Run("TestReadRDPoolResourceTypeA",
		func(st *testing.T) {
			it.Test = st
			it.ReadRecord(integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, pool.RD))
		})
	t.Test.Run("TestReadPoolResourceValidation",
		func(st *testing.T) {
			it.Test = st
			it.ReadPoolValidationFailure(integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, pool.SB))
		})
	t.Test.Run("TestDeleteRDPoolResourceTypeA",
		func(st *testing.T) {
			it.Test = st
			it.DeleteRecord(integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, ""))
		})
}

func (it *IntegrationTest) CreateRDPoolTypeA(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, "")
	rrSet := getRDPoolTypeA(ownerName)
	it.CreateRecord(rrSetKey, rrSet)
}

func (it *IntegrationTest) UpdateRDPoolTypeA(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, "")
	rrSet := getRDPoolTypeA(ownerName)
	rrSet.RData = []string{"192.168.1.11"}
	it.UpdateRecord(rrSetKey, rrSet)
}

func (it *IntegrationTest) PartialUpdateRDPoolTypeA(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, "")
	rrSet := getRDPoolTypeA(ownerName)
	rrSet.RData = []string{"192.168.1.12"}
	it.PartialUpdateRecord(rrSetKey, rrSet)
}

func (it *IntegrationTest) ReadPoolValidationFailure(rrSetKey *rrset.RRSetKey) {
	recordService, err := record.Get(integration.TestClient)

	if err != nil {
		it.Test.Fatal(err)
	}

	if _, _, er := recordService.Read(rrSetKey); er.Error() != fmt.Sprintf("Record resource of type SB_POOL - %v not found", rrSetKey.RecordID()) {
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

package integration_test

import (
	"testing"

	"github.com/ultradns/ultradns-go-sdk/internal/testing/integration"
	"github.com/ultradns/ultradns-go-sdk/pkg/record"
	"github.com/ultradns/ultradns-go-sdk/pkg/rrset"
)

const (
	testRecordTypeA    = "A"
	testRecordTypeAAAA = "AAAA"
)

func TestRecordResources(t *testing.T) {
	it := IntegrationTest{}
	zoneName := integration.GetRandomZoneName()
	ownerName := integration.GetRandomString()

	t.Parallel()

	t.Run("TestCreateRecordResourceZone",
		func(st *testing.T) {
			it.Test = st
			it.CreatePrimaryZone(zoneName)
		})
	t.Run("TestCreateRecordResourceTypeA",
		func(st *testing.T) {
			it.Test = st
			it.CreateRecordTypeA(ownerName, zoneName)
		})
	t.Run("TestUpdateRecordResourceTypeA",
		func(st *testing.T) {
			it.Test = st
			it.UpdateRecordTypeA(ownerName, zoneName)
		})
	t.Run("TestPartialUpdateRecordResourceTypeA",
		func(st *testing.T) {
			it.Test = st
			it.PartialUpdateRecordTypeA(ownerName, zoneName)
		})
	t.Run("TestReadRecordResourceTypeA",
		func(st *testing.T) {
			it.Test = st
			it.ReadRecord(integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA))
		})
	t.Run("TestDeleteRecordResourceTypeA",
		func(st *testing.T) {
			it.Test = st
			it.DeleteRecord(integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA))
		})
	t.Run("TestDeleteRecordResourceZone",
		func(st *testing.T) {
			it.Test = st
			it.DeleteZone(zoneName)
		})
}

func (it *IntegrationTest) CreateRecordTypeA(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA)
	rrSet := getRRSetTypeA(ownerName)
	it.CreateRecord(rrSetKey, rrSet)
}

func (it *IntegrationTest) UpdateRecordTypeA(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA)
	rrSet := getRRSetTypeA(ownerName)
	rrSet.RData = []string{"192.168.1.11"}
	it.UpdateRecord(rrSetKey, rrSet)
}

func (it *IntegrationTest) PartialUpdateRecordTypeA(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA)
	rrSet := getRRSetTypeA(ownerName)
	rrSet.RData = []string{"192.168.1.12"}
	it.PartialUpdateRecord(rrSetKey, rrSet)
}

func (it *IntegrationTest) CreateRecord(rrSetKey *rrset.RRSetKey, rrSet *rrset.RRSet) {
	recordService, err := record.Get(integration.TestClient)

	if err != nil {
		it.Test.Fatal(err)
	}

	if _, er := recordService.Create(rrSetKey, rrSet); er != nil {
		it.Test.Fatal(er)
	}
}

func (it *IntegrationTest) UpdateRecord(rrSetKey *rrset.RRSetKey, rrSet *rrset.RRSet) {
	recordService, err := record.Get(integration.TestClient)

	if err != nil {
		it.Test.Fatal(err)
	}

	if _, er := recordService.Update(rrSetKey, rrSet); er != nil {
		it.Test.Fatal(er)
	}
}

func (it *IntegrationTest) PartialUpdateRecord(rrSetKey *rrset.RRSetKey, rrSet *rrset.RRSet) {
	recordService, err := record.Get(integration.TestClient)

	if err != nil {
		it.Test.Fatal(err)
	}

	if _, er := recordService.PartialUpdate(rrSetKey, rrSet); er != nil {
		it.Test.Fatal(er)
	}
}

func (it *IntegrationTest) ReadRecord(rrSetKey *rrset.RRSetKey) {
	recordService, err := record.Get(integration.TestClient)

	if err != nil {
		it.Test.Fatal(err)
	}

	if _, _, er := recordService.Read(rrSetKey); er != nil {
		it.Test.Fatal(er)
	}
}

func (it *IntegrationTest) DeleteRecord(rrSetKey *rrset.RRSetKey) {
	recordService, err := record.Get(integration.TestClient)

	if err != nil {
		it.Test.Fatal(err)
	}

	if _, er := recordService.Delete(rrSetKey); er != nil {
		it.Test.Fatal(er)
	}
}

func getRRSetTypeA(ownerName string) *rrset.RRSet {
	return &rrset.RRSet{
		OwnerName: ownerName,
		RRType:    testRecordTypeA,
		RData:     []string{"192.168.1.1"},
	}
}

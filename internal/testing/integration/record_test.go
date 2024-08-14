package integration_test

import (
	"testing"

	"github.com/ultradns/ultradns-go-sdk/internal/testing/integration"
	"github.com/ultradns/ultradns-go-sdk/pkg/helper"
	"github.com/ultradns/ultradns-go-sdk/pkg/record"
	"github.com/ultradns/ultradns-go-sdk/pkg/rrset"
)

func (t *IntegrationTest) TestRecordResources(zoneName string) {
	it := IntegrationTest{}
	ownerName := integration.GetRandomString()

	t.Test.Run("TestCreateRecordResourceTypeA",
		func(st *testing.T) {
			it.Test = st
			it.CreateRecordTypeA(ownerName, zoneName)
		})
	t.Test.Run("TestUpdateRecordResourceTypeA",
		func(st *testing.T) {
			it.Test = st
			it.UpdateRecordTypeA(ownerName, zoneName)
		})
	t.Test.Run("TestPartialUpdateRecordResourceTypeA",
		func(st *testing.T) {
			it.Test = st
			it.PartialUpdateRecordTypeA(ownerName, zoneName)
		})
	t.Test.Run("TestReadRecordResourceTypeA",
		func(st *testing.T) {
			it.Test = st
			it.ReadRecord(integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, ""))
		})
	t.Test.Run("TestListRecordResourceTypeA",
		func(st *testing.T) {
			it.Test = st
			it.ListRecord(integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, ""))
		})
	t.Test.Run("TestDeleteRecordResourceTypeA",
		func(st *testing.T) {
			it.Test = st
			it.DeleteRecord(integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, ""))
		})
}

func (t *IntegrationTest) CreateRecordTypeA(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, "")
	rrSet := getRRSetTypeA(ownerName)
	t.CreateRecord(rrSetKey, rrSet)
}

func (t *IntegrationTest) UpdateRecordTypeA(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, "")
	rrSet := getRRSetTypeA(ownerName)
	rrSet.RData = []string{"192.168.1.11"}
	t.UpdateRecord(rrSetKey, rrSet)
}

func (t *IntegrationTest) PartialUpdateRecordTypeA(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, "")
	rrSet := getRRSetTypeA(ownerName)
	rrSet.RData = []string{"192.168.1.12"}
	t.PartialUpdateRecord(rrSetKey, rrSet)
}

func getRRSetTypeA(ownerName string) *rrset.RRSet {
	return &rrset.RRSet{
		OwnerName: ownerName,
		RRType:    testRecordTypeA,
		RData:     []string{"192.168.1.1"},
	}
}

func (t *IntegrationTest) CreateRecord(rrSetKey *rrset.RRSetKey, rrSet *rrset.RRSet) {
	recordService, err := record.Get(integration.TestClient)

	if err != nil {
		t.Test.Fatal(err)
	}

	if _, er := recordService.Create(rrSetKey, rrSet); er != nil {
		t.Test.Fatal(er)
	}
}

func (t *IntegrationTest) UpdateRecord(rrSetKey *rrset.RRSetKey, rrSet *rrset.RRSet) {
	recordService, err := record.Get(integration.TestClient)

	if err != nil {
		t.Test.Fatal(err)
	}

	if _, er := recordService.Update(rrSetKey, rrSet); er != nil {
		t.Test.Fatal(er)
	}
}

func (t *IntegrationTest) PartialUpdateRecord(rrSetKey *rrset.RRSetKey, rrSet *rrset.RRSet) {
	recordService, err := record.Get(integration.TestClient)

	if err != nil {
		t.Test.Fatal(err)
	}

	if _, er := recordService.PartialUpdate(rrSetKey, rrSet); er != nil {
		t.Test.Fatal(er)
	}
}

func (t *IntegrationTest) ReadRecord(rrSetKey *rrset.RRSetKey) {
	recordService, err := record.Get(integration.TestClient)

	if err != nil {
		t.Test.Fatal(err)
	}

	if _, _, er := recordService.Read(rrSetKey); er != nil {
		t.Test.Fatal(er)
	}
}

func (t *IntegrationTest) ListRecord(rrSetKey *rrset.RRSetKey) {
	recordService, err := record.Get(integration.TestClient)

	if err != nil {
		t.Test.Fatal(err)
	}

	if _, _, er := recordService.List(rrSetKey, &helper.QueryInfo{}); er != nil {
		t.Test.Fatal(er)
	}
}

func (t *IntegrationTest) DeleteRecord(rrSetKey *rrset.RRSetKey) {
	recordService, err := record.Get(integration.TestClient)

	if err != nil {
		t.Test.Fatal(err)
	}

	if _, er := recordService.Delete(rrSetKey); er != nil {
		t.Test.Fatal(er)
	}
}

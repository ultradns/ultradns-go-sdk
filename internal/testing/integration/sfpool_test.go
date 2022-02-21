package integration_test

import (
	"testing"

	"github.com/ultradns/ultradns-go-sdk/internal/testing/integration"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/pool"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/sfpool"
	"github.com/ultradns/ultradns-go-sdk/pkg/rrset"
)

func (t *IntegrationTest) TestSFPoolResources(zoneName string) {
	it := IntegrationTest{}
	ownerName := integration.GetRandomString()

	t.Test.Run("TestCreateSFPoolResourceTypeAAAA",
		func(st *testing.T) {
			it.Test = st
			it.CreateSFPoolTypeAAAA(ownerName, zoneName)
		})
	t.Test.Run("TestUpdateSFPoolResourceTypeAAAA",
		func(st *testing.T) {
			it.Test = st
			it.UpdateSFPoolTypeAAAA(ownerName, zoneName)
		})
	t.Test.Run("TestPartialUpdateSFPoolResourceTypeAAAA",
		func(st *testing.T) {
			it.Test = st
			it.PartialUpdateSFPoolTypeAAAA(ownerName, zoneName)
		})
	t.Test.Run("TestReadSFPoolResourceTypeAAAA",
		func(st *testing.T) {
			it.Test = st
			it.ReadRecord(integration.GetRRSetKey(ownerName, zoneName, testRecordTypeAAAA, pool.SF))
		})
	t.Test.Run("TestDeleteSFPoolResourceTypeAAAA",
		func(st *testing.T) {
			it.Test = st
			it.DeleteRecord(integration.GetRRSetKey(ownerName, zoneName, testRecordTypeAAAA, ""))
		})
}

func (it *IntegrationTest) CreateSFPoolTypeAAAA(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeAAAA, "")
	rrSet := getSFPoolTypeAAAA(ownerName)
	it.CreateRecord(rrSetKey, rrSet)
}

func (it *IntegrationTest) UpdateSFPoolTypeAAAA(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeAAAA, "")
	rrSet := getSFPoolTypeAAAA(ownerName)
	rrSet.RData = []string{"aaaa:bbbb:cccc:dddd:eeee:ffff:1:11"}
	it.UpdateRecord(rrSetKey, rrSet)
}

func (it *IntegrationTest) PartialUpdateSFPoolTypeAAAA(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeAAAA, "")
	rrSet := getSFPoolTypeAAAA(ownerName)
	rrSet.RData = []string{"aaaa:bbbb:cccc:dddd:eeee:ffff:1:12"}
	it.PartialUpdateRecord(rrSetKey, rrSet)
}

func getSFPoolTypeAAAA(ownerName string) *rrset.RRSet {
	backupRecord := &sfpool.BackupRecord{
		RData: "aaaa:bbbb:cccc:dddd:eeee:ffff:1:2",
	}
	monitor := &pool.Monitor{
		Method: "GET",
		URL:    integration.TestHost,
	}
	profile := &sfpool.Profile{
		BackupRecord:             backupRecord,
		Monitor:                  monitor,
		RegionFailureSensitivity: "HIGH",
	}

	return &rrset.RRSet{
		OwnerName: ownerName,
		RRType:    testRecordTypeAAAA,
		RData:     []string{"aaaa:bbbb:cccc:dddd:eeee:ffff:1:1"},
		Profile:   profile,
	}
}

package integration_test

import (
	"testing"

	"github.com/ultradns/ultradns-go-sdk/internal/testing/integration"
)

const (
	testRecordTypeA    = "A"
	testRecordTypeAAAA = "AAAA"
)

type IntegrationTest struct {
	Test *testing.T
}

func TestRecordResources(t *testing.T) {
	t.Parallel()

	it := IntegrationTest{}
	zoneName := integration.GetRandomZoneName()

	ownerNameSB := integration.GetRandomString()
	ownerNameTC := integration.GetRandomString()
	groupNameIP := integration.GetRandomString()
	groupNameGeo := integration.GetRandomString()

	t.Run("TestCreateZoneRecordResources",
		func(st *testing.T) {
			integration.TestClient.EnableDefaultDebugLogger()
			it.Test = st
			it.CreatePrimaryZone(zoneName)
		})
	t.Run("TestStandandRecordResources",
		func(st *testing.T) {
			integration.TestClient.DisableLogger()
			it.Test = st
			it.TestRecordResources(zoneName)
		})
	t.Run("TestRDPoolRecordResources",
		func(st *testing.T) {
			integration.TestClient.EnableDefaultWarnLogger()
			it.Test = st
			it.TestRDPoolResources(zoneName)
		})
	t.Run("TestSFPoolRecordResources",
		func(st *testing.T) {
			integration.TestClient.DisableLogger()
			it.Test = st
			it.TestSFPoolResources(zoneName)
		})
	t.Run("TestSLBPoolRecordResources",
		func(st *testing.T) {
			it.Test = st
			it.TestSLBPoolResources(zoneName)
		})
	t.Run("TestSBPoolRecordResources",
		func(st *testing.T) {
			it.Test = st
			it.TestSBPoolResources(zoneName, ownerNameSB)
		})
	t.Run("TestTCPoolRecordResources",
		func(st *testing.T) {
			it.Test = st
			it.TestTCPoolResources(zoneName, ownerNameTC)
		})
	t.Run("TestProbeResources",
		func(st *testing.T) {
			it.Test = st
			it.TestProbeResources(zoneName, ownerNameSB, ownerNameTC)
		})
	t.Run("TestDIRPoolRecordResources",
		func(st *testing.T) {
			it.Test = st
			it.TestDIRPoolResources(zoneName)
		})
	t.Run("TestDIRGroupGeoResources",
		func(st *testing.T) {
			it.Test = st
			it.TestDirGroupGeoResources(getDirGroupGeo(groupNameGeo), groupNameGeo)
		})
	t.Run("TestDIRGroupIPResources",
		func(st *testing.T) {
			it.Test = st
			it.TestDirGroupIPResources(getDirGroupIP(groupNameIP), groupNameIP)
		})
	t.Run("TestDeleteZoneRecordResources",
		func(st *testing.T) {
			integration.TestClient.EnableDefaultTraceLogger()
			it.Test = st
			it.DeleteZone(zoneName)
		})
}

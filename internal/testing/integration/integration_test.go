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

	t.Run("TestCreateZoneRecordResources",
		func(st *testing.T) {
			it.Test = st
			it.CreatePrimaryZone(zoneName)
		})
	t.Run("TestStandandRecordResources",
		func(st *testing.T) {
			it.Test = st
			it.TestRecordResources(zoneName)
		})
	t.Run("TestRDPoolRecordResources",
		func(st *testing.T) {
			it.Test = st
			it.TestRDPoolResources(zoneName)
		})
	t.Run("TestSFPoolRecordResources",
		func(st *testing.T) {
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
	t.Run("TestHTTPProbeResources",
		func(st *testing.T) {
			it.Test = st
			it.TestHTTPProbeResources(zoneName, ownerNameSB)
		})
	t.Run("TestFTPProbeResources",
		func(st *testing.T) {
			it.Test = st
			it.TestFTPProbeResources(zoneName, ownerNameSB)
		})
	t.Run("TestTCPProbeResources",
		func(st *testing.T) {
			it.Test = st
			it.TestTCPProbeResources(zoneName, ownerNameSB)
		})
	t.Run("TestPINGProbeResources",
		func(st *testing.T) {
			it.Test = st
			it.TestPINGProbeResources(zoneName, ownerNameTC)
		})
	t.Run("TestDNSProbeResources",
		func(st *testing.T) {
			it.Test = st
			it.TestDNSProbeResources(zoneName, ownerNameTC)
		})
	t.Run("TestSMTPProbeResources",
		func(st *testing.T) {
			it.Test = st
			it.TestSMTPProbeResources(zoneName, ownerNameTC)
		})
	t.Run("TestSMTPSendProbeResources",
		func(st *testing.T) {
			it.Test = st
			it.TestSMTPSendProbeResources(zoneName, ownerNameSB)
		})
	t.Run("TestDIRPoolRecordResources",
		func(st *testing.T) {
			it.Test = st
			it.TestDIRPoolResources(zoneName)
		})
	t.Run("TestDeleteZoneRecordResources",
		func(st *testing.T) {
			it.Test = st
			it.DeleteZone(zoneName)
		})
}

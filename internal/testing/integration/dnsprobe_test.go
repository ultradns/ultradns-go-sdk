package integration_test

import (
	"testing"

	"github.com/ultradns/ultradns-go-sdk/internal/testing/integration"
	"github.com/ultradns/ultradns-go-sdk/pkg/probe"
	"github.com/ultradns/ultradns-go-sdk/pkg/probe/dns"
	"github.com/ultradns/ultradns-go-sdk/pkg/probe/helper"
)

func (t *IntegrationTest) TestDNSProbeResources(zoneName, ownerName string) {
	it := IntegrationTest{}

	t.Test.Run("TestCreateProbeResourceTypeDNS",
		func(st *testing.T) {
			it.Test = st
			it.CreateProbeTypeDNS(ownerName, zoneName)
		})
	t.Test.Run("TestListProbeResourceTypeDNS",
		func(st *testing.T) {
			it.Test = st
			it.ListProbe(integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, probe.DNS))
		})
	t.Test.Run("TestUpdateProbeResourceTypeDNS",
		func(st *testing.T) {
			it.Test = st
			it.UpdateProbeTypeDNS(ownerName, zoneName)
		})
	t.Test.Run("TestPartialUpdateProbeResourceTypeDNS",
		func(st *testing.T) {
			it.Test = st
			it.PartialUpdateProbeTypeDNS(ownerName, zoneName)
		})
	t.Test.Run("TestReadProbeResourceTypeDNS",
		func(st *testing.T) {
			it.Test = st
			it.ReadProbe(integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, probe.DNS))
		})
	t.Test.Run("TestDeleteProbeResourceTypeDNS",
		func(st *testing.T) {
			it.Test = st
			it.DeleteProbe(integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, probe.DNS))
		})
}

func (it *IntegrationTest) CreateProbeTypeDNS(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, probe.DNS)
	probedata := getProbeTypeDNS()
	it.CreateProbe(rrSetKey, probedata)
}

func (it *IntegrationTest) UpdateProbeTypeDNS(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, probe.DNS)
	probedata := getProbeTypeDNS()
	probedata.Interval = "FIFTEEN_MINUTES"
	it.UpdateProbe(rrSetKey, probedata)
}

func (it *IntegrationTest) PartialUpdateProbeTypeDNS(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, probe.DNS)
	probedata := getProbeTypeDNS()
	limit := &helper.Limit{
		Fail: 20,
	}
	limitInfo := &helper.LimitsInfo{
		Run: limit,
	}
	details := &dns.Details{
		Port:   53,
		Limits: limitInfo,
	}
	probedata.Details = details
	it.PartialUpdateProbe(rrSetKey, probedata)
}

func getProbeTypeDNS() *probe.Probe {
	limit := &helper.Limit{
		Fail: 10,
	}
	limitInfo := &helper.LimitsInfo{
		Run: limit,
	}
	details := &dns.Details{
		Port:   53,
		Limits: limitInfo,
	}
	probedata := &probe.Probe{
		Type:      probe.DNS,
		Interval:  "ONE_MINUTE",
		Agents:    []string{"NEW_YORK", "DALLAS"},
		Threshold: 2,
		Details:   details,
	}
	return probedata
}

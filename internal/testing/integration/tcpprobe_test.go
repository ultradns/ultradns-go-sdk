package integration_test

import (
	"testing"

	"github.com/ultradns/ultradns-go-sdk/internal/testing/integration"
	"github.com/ultradns/ultradns-go-sdk/pkg/probe"
	"github.com/ultradns/ultradns-go-sdk/pkg/probe/helper"
	"github.com/ultradns/ultradns-go-sdk/pkg/probe/tcp"
)

func (t *IntegrationTest) TestTCPProbeResources(zoneName, ownerName string) {
	it := IntegrationTest{}

	t.Test.Run("TestCreateProbeResourceTypeTCP",
		func(st *testing.T) {
			it.Test = st
			it.CreateProbeTypeTCP(ownerName, zoneName)
		})
	t.Test.Run("TestListProbeResourceTypeTCP",
		func(st *testing.T) {
			it.Test = st
			it.ListProbe(integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, probe.TCP))
		})
	t.Test.Run("TestUpdateProbeResourceTypeTCP",
		func(st *testing.T) {
			it.Test = st
			it.UpdateProbeTypeTCP(ownerName, zoneName)
		})
	t.Test.Run("TestPartialUpdateProbeResourceTypeTCP",
		func(st *testing.T) {
			it.Test = st
			it.PartialUpdateProbeTypeTCP(ownerName, zoneName)
		})
	t.Test.Run("TestReadProbeResourceTypeTCP",
		func(st *testing.T) {
			it.Test = st
			it.ReadProbe(integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, probe.TCP))
		})
	t.Test.Run("TestDeleteProbeResourceTypeTCP",
		func(st *testing.T) {
			it.Test = st
			it.DeleteProbe(integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, probe.TCP))
		})
}

func (it *IntegrationTest) CreateProbeTypeTCP(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, probe.TCP)
	probedata := getProbeTypeTCP()
	it.CreateProbe(rrSetKey, probedata)
}

func (it *IntegrationTest) UpdateProbeTypeTCP(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, probe.TCP)
	probedata := getProbeTypeTCP()
	probedata.Interval = "FIFTEEN_MINUTES"
	it.UpdateProbe(rrSetKey, probedata)
}

func (it *IntegrationTest) PartialUpdateProbeTypeTCP(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, probe.TCP)
	probedata := getProbeTypeTCP()
	limit := &helper.Limit{
		Fail: 20,
	}
	limitInfo := &helper.LimitsInfo{
		Connect: limit,
	}
	details := &tcp.Details{
		Port:   53,
		Limits: limitInfo,
	}
	probedata.Details = details
	it.PartialUpdateProbe(rrSetKey, probedata)
}

func getProbeTypeTCP() *probe.Probe {
	limit := &helper.Limit{
		Fail: 10,
	}
	limitInfo := &helper.LimitsInfo{
		Connect: limit,
	}
	details := &tcp.Details{
		Port:   53,
		Limits: limitInfo,
	}
	probedata := &probe.Probe{
		Type:      probe.TCP,
		Interval:  "ONE_MINUTE",
		Agents:    []string{"NEW_YORK", "DALLAS"},
		Threshold: 2,
		Details:   details,
	}
	return probedata
}

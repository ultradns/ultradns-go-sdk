package integration_test

import (
	"testing"

	"github.com/ultradns/ultradns-go-sdk/internal/testing/integration"
	"github.com/ultradns/ultradns-go-sdk/pkg/probe"
	"github.com/ultradns/ultradns-go-sdk/pkg/probe/helper"
	"github.com/ultradns/ultradns-go-sdk/pkg/probe/ping"
)

func (t *IntegrationTest) TestPINGProbeResources(zoneName, ownerName string) {
	it := IntegrationTest{}

	t.Test.Run("TestCreateProbeResourceTypePING",
		func(st *testing.T) {
			it.Test = st
			it.CreateProbeTypePING(ownerName, zoneName)
		})
	t.Test.Run("TestListProbeResourceTypePING",
		func(st *testing.T) {
			it.Test = st
			it.ListProbe(integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, probe.PING))
		})
	t.Test.Run("TestUpdateProbeResourceTypePING",
		func(st *testing.T) {
			it.Test = st
			it.UpdateProbeTypePING(ownerName, zoneName)
		})
	t.Test.Run("TestPartialUpdateProbeResourceTypePING",
		func(st *testing.T) {
			it.Test = st
			it.PartialUpdateProbeTypePING(ownerName, zoneName)
		})
	t.Test.Run("TestReadProbeResourceTypePING",
		func(st *testing.T) {
			it.Test = st
			it.ReadProbe(integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, probe.PING))
		})
	t.Test.Run("TestDeleteProbeResourceTypePING",
		func(st *testing.T) {
			it.Test = st
			it.DeleteProbe(integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, probe.PING))
		})
}

func (it *IntegrationTest) CreateProbeTypePING(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, probe.PING)
	probedata := getProbeTypePING()
	it.CreateProbe(rrSetKey, probedata)
}

func (it *IntegrationTest) UpdateProbeTypePING(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, probe.PING)
	probedata := getProbeTypePING()
	probedata.Interval = "FIFTEEN_MINUTES"
	it.UpdateProbe(rrSetKey, probedata)
}

func (it *IntegrationTest) PartialUpdateProbeTypePING(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, probe.PING)
	probedata := getProbeTypePING()
	limit := &helper.Limit{
		Fail: 20,
	}
	limitInfo := &helper.LimitsInfo{
		LossPercent: limit,
		Total:       limit,
	}
	details := &ping.Details{
		Packets:    5,
		PacketSize: 56,
		Limits:     limitInfo,
	}
	probedata.Details = details
	it.PartialUpdateProbe(rrSetKey, probedata)
}

func getProbeTypePING() *probe.Probe {
	limit := &helper.Limit{
		Fail: 10,
	}
	limitInfo := &helper.LimitsInfo{
		LossPercent: limit,
		Total:       limit,
	}
	details := &ping.Details{
		Packets:    3,
		PacketSize: 56,
		Limits:     limitInfo,
	}
	probedata := &probe.Probe{
		Type:      probe.PING,
		Interval:  "ONE_MINUTE",
		Agents:    []string{"NEW_YORK", "DALLAS"},
		Threshold: 2,
		Details:   details,
	}
	return probedata
}

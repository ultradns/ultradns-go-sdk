package integration_test

import (
	"testing"

	"github.com/ultradns/ultradns-go-sdk/internal/testing/integration"
	"github.com/ultradns/ultradns-go-sdk/pkg/probe"
	"github.com/ultradns/ultradns-go-sdk/pkg/probe/helper"
	"github.com/ultradns/ultradns-go-sdk/pkg/probe/smtp"
)

func (t *IntegrationTest) TestSMTPProbeResources(zoneName, ownerName string) {
	it := IntegrationTest{}

	t.Test.Run("TestCreateProbeResourceTypeSMTP",
		func(st *testing.T) {
			it.Test = st
			it.CreateProbeTypeSMTP(ownerName, zoneName)
		})
	t.Test.Run("TestListProbeResourceTypeSMTP",
		func(st *testing.T) {
			it.Test = st
			it.ListProbe(integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, probe.SMTP))
		})
	t.Test.Run("TestUpdateProbeResourceTypeSMTP",
		func(st *testing.T) {
			it.Test = st
			it.UpdateProbeTypeSMTP(ownerName, zoneName)
		})
	t.Test.Run("TestPartialUpdateProbeResourceTypeSMTP",
		func(st *testing.T) {
			it.Test = st
			it.PartialUpdateProbeTypeSMTP(ownerName, zoneName)
		})
	t.Test.Run("TestReadProbeResourceTypeSMTP",
		func(st *testing.T) {
			it.Test = st
			it.ReadProbe(integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, probe.SMTP))
		})
	t.Test.Run("TestDeleteProbeResourceTypeSMTP",
		func(st *testing.T) {
			it.Test = st
			it.DeleteProbe(integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, probe.SMTP))
		})
}

func (it *IntegrationTest) CreateProbeTypeSMTP(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, probe.SMTP)
	probedata := getProbeTypeSMTP()
	it.CreateProbe(rrSetKey, probedata)
}

func (it *IntegrationTest) UpdateProbeTypeSMTP(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, probe.SMTP)
	probedata := getProbeTypeSMTP()
	probedata.Interval = "FIFTEEN_MINUTES"
	it.UpdateProbe(rrSetKey, probedata)
}

func (it *IntegrationTest) PartialUpdateProbeTypeSMTP(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, probe.SMTP)
	probedata := getProbeTypeSMTP()
	limit := &helper.Limit{
		Fail: 20,
	}
	limitInfo := &helper.LimitsInfo{
		Connect: limit,
	}
	details := &smtp.Details{
		Port:   25,
		Limits: limitInfo,
	}
	probedata.Details = details
	it.PartialUpdateProbe(rrSetKey, probedata)
}

func getProbeTypeSMTP() *probe.Probe {
	limit := &helper.Limit{
		Fail: 30,
	}
	limitInfo := &helper.LimitsInfo{
		Run: limit,
	}
	details := &smtp.Details{
		Port:   25,
		Limits: limitInfo,
	}
	probedata := &probe.Probe{
		Type:      probe.SMTP,
		Interval:  "ONE_MINUTE",
		Agents:    []string{"NEW_YORK", "DALLAS"},
		Threshold: 2,
		Details:   details,
	}
	return probedata
}

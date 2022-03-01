package integration_test

import (
	"testing"

	"github.com/ultradns/ultradns-go-sdk/internal/testing/integration"
	"github.com/ultradns/ultradns-go-sdk/pkg/probe"
	"github.com/ultradns/ultradns-go-sdk/pkg/probe/helper"
	"github.com/ultradns/ultradns-go-sdk/pkg/probe/smtpsend"
)

func (t *IntegrationTest) TestSMTPSendProbeResources(zoneName, ownerName string) {
	it := IntegrationTest{}

	t.Test.Run("TestCreateProbeResourceTypeSMTPSend",
		func(st *testing.T) {
			it.Test = st
			it.CreateProbeTypeSMTPSend(ownerName, zoneName)
		})
	t.Test.Run("TestListProbeResourceTypeSMTPSend",
		func(st *testing.T) {
			it.Test = st
			it.ListProbe(integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, probe.SMTPSend))
		})
	t.Test.Run("TestUpdateProbeResourceTypeSMTPSend",
		func(st *testing.T) {
			it.Test = st
			it.UpdateProbeTypeSMTPSend(ownerName, zoneName)
		})
	t.Test.Run("TestPartialUpdateProbeResourceTypeSMTPSend",
		func(st *testing.T) {
			it.Test = st
			it.PartialUpdateProbeTypeSMTPSend(ownerName, zoneName)
		})
	t.Test.Run("TestReadProbeResourceTypeSMTPSend",
		func(st *testing.T) {
			it.Test = st
			it.ReadProbe(integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, probe.SMTPSend))
		})
	t.Test.Run("TestDeleteProbeResourceTypeSMTPSend",
		func(st *testing.T) {
			it.Test = st
			it.DeleteProbe(integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, probe.SMTPSend))
		})
}

func (t *IntegrationTest) CreateProbeTypeSMTPSend(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, probe.SMTPSend)
	probedata := getProbeTypeSMTPSend()
	t.CreateProbe(rrSetKey, probedata)
}

func (t *IntegrationTest) UpdateProbeTypeSMTPSend(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, probe.SMTPSend)
	probedata := getProbeTypeSMTPSend()
	probedata.Interval = testProbeInterval
	t.UpdateProbe(rrSetKey, probedata)
}

func (t *IntegrationTest) PartialUpdateProbeTypeSMTPSend(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, probe.SMTPSend)
	probedata := getProbeTypeSMTPSend()
	limit := &helper.Limit{
		Fail: 20,
	}
	limitInfo := &helper.LimitsInfo{
		Connect: limit,
	}
	details := &smtpsend.Details{
		Limits: limitInfo,
	}
	probedata.Details = details
	t.PartialUpdateProbe(rrSetKey, probedata)
}

func getProbeTypeSMTPSend() *probe.Probe {
	limit := &helper.Limit{
		Fail: 30,
	}
	limitInfo := &helper.LimitsInfo{
		Run: limit,
	}
	details := &smtpsend.Details{
		Port:    25,
		From:    "from@ultradns.com",
		To:      "to@ultradns.com",
		Message: "Message",
		Limits:  limitInfo,
	}
	probedata := &probe.Probe{
		Type:      probe.SMTPSend,
		Interval:  "ONE_MINUTE",
		Agents:    []string{"NEW_YORK", "DALLAS"},
		Threshold: 2,
		Details:   details,
	}

	return probedata
}

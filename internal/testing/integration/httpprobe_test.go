package integration_test

import (
	"fmt"
	"testing"

	"github.com/ultradns/ultradns-go-sdk/internal/testing/integration"
	"github.com/ultradns/ultradns-go-sdk/pkg/probe"
	"github.com/ultradns/ultradns-go-sdk/pkg/probe/helper"
	"github.com/ultradns/ultradns-go-sdk/pkg/probe/http"
	"github.com/ultradns/ultradns-go-sdk/pkg/rrset"
)

func (t *IntegrationTest) TestHTTPProbeResources(zoneName, ownerName string) {
	it := IntegrationTest{}

	t.Test.Run("TestCreateProbeResourceTypeHTTP",
		func(st *testing.T) {
			it.Test = st
			it.CreateProbeTypeHTTP(ownerName, zoneName)
		})
	t.Test.Run("TestListProbeResourceTypeHTTP",
		func(st *testing.T) {
			it.Test = st
			it.ListProbe(integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, probe.HTTP))
		})
	t.Test.Run("TestUpdateProbeResourceTypeHTTP",
		func(st *testing.T) {
			it.Test = st
			it.UpdateProbeTypeHTTP(ownerName, zoneName)
		})
	t.Test.Run("TestPartialUpdateProbeResourceTypeHTTP",
		func(st *testing.T) {
			it.Test = st
			it.PartialUpdateProbeTypeHTTP(ownerName, zoneName)
		})
	t.Test.Run("TestReadProbeResourceTypeHTTP",
		func(st *testing.T) {
			it.Test = st
			it.ReadProbe(integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, probe.HTTP))
		})
	t.Test.Run("TestReadProbeResourceValidation",
		func(st *testing.T) {
			it.Test = st
			it.ReadProbeValidationFailure(integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, probe.TCP))
		})
	t.Test.Run("TestDeleteProbeResourceTypeHTTP",
		func(st *testing.T) {
			it.Test = st
			it.DeleteProbe(integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, probe.HTTP))
		})
}

func (t *IntegrationTest) CreateProbeTypeHTTP(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, probe.HTTP)
	probedata := getProbeTypeHTTP()
	t.CreateProbe(rrSetKey, probedata)
}

func (t *IntegrationTest) UpdateProbeTypeHTTP(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, probe.HTTP)
	probedata := getProbeTypeHTTP()
	probedata.Interval = testProbeInterval
	t.UpdateProbe(rrSetKey, probedata)
}

func (t *IntegrationTest) PartialUpdateProbeTypeHTTP(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, probe.HTTP)
	probedata := getProbeTypeHTTP()
	limit := &helper.Limit{
		Fail: 20,
	}
	limitInfo := &helper.LimitsInfo{
		Run:     limit,
		Connect: limit,
	}
	transction := &http.Transaction{
		Method:          "GET",
		ProtocolVersion: "HTTP/1.0",
		URL:             integration.TestHost,
		Limits:          limitInfo,
	}
	details := &http.Details{
		Transactions: []*http.Transaction{transction},
	}
	probedata.Details = details
	t.PartialUpdateProbe(rrSetKey, probedata)
}

func (t *IntegrationTest) ReadProbeValidationFailure(rrSetKey *rrset.RRSetKey) {
	probeService, err := probe.Get(integration.TestClient)

	if err != nil {
		t.Test.Fatal(err)
	}

	rrSetKey.ID = probeID

	if _, _, er := probeService.Read(rrSetKey); er.Error() != fmt.Sprintf("Probe resource of type TCP - %v not found", rrSetKey.PID()) {
		t.Test.Fatal(er)
	}
}

func getProbeTypeHTTP() *probe.Probe {
	limit := &helper.Limit{
		Fail: 10,
	}
	limitInfo := &helper.LimitsInfo{
		Run:     limit,
		Connect: limit,
	}
	transction := &http.Transaction{
		Method:          "GET",
		ProtocolVersion: "HTTP/1.0",
		URL:             integration.TestHost,
		Limits:          limitInfo,
	}
	details := &http.Details{
		Transactions: []*http.Transaction{transction},
	}
	probedata := &probe.Probe{
		Type:      probe.HTTP,
		Interval:  "ONE_MINUTE",
		Agents:    []string{"NEW_YORK", "DALLAS"},
		Threshold: 2,
		Details:   details,
	}

	return probedata
}

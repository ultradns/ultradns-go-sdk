package integration_test

import (
	"testing"

	"github.com/ultradns/ultradns-go-sdk/internal/testing/integration"
	"github.com/ultradns/ultradns-go-sdk/pkg/probe"
	"github.com/ultradns/ultradns-go-sdk/pkg/rrset"
)

var probeID = ""

const (
	testProbeInterval = "FIFTEEN_MINUTES"
)

func (t *IntegrationTest) TestProbeResources(zoneName, ownerNameSB, ownerNameTC string) {
	it := IntegrationTest{}

	t.Test.Run("TestHTTPProbeResources",
		func(st *testing.T) {
			it.Test = st
			it.TestHTTPProbeResources(zoneName, ownerNameSB)
		})
	t.Test.Run("TestFTPProbeResources",
		func(st *testing.T) {
			it.Test = st
			it.TestFTPProbeResources(zoneName, ownerNameSB)
		})
	t.Test.Run("TestTCPProbeResources",
		func(st *testing.T) {
			it.Test = st
			it.TestTCPProbeResources(zoneName, ownerNameSB)
		})
	t.Test.Run("TestPINGProbeResources",
		func(st *testing.T) {
			it.Test = st
			it.TestPINGProbeResources(zoneName, ownerNameTC)
		})
	t.Test.Run("TestDNSProbeResources",
		func(st *testing.T) {
			it.Test = st
			it.TestDNSProbeResources(zoneName, ownerNameTC)
		})
	t.Test.Run("TestSMTPProbeResources",
		func(st *testing.T) {
			it.Test = st
			it.TestSMTPProbeResources(zoneName, ownerNameTC)
		})
	t.Test.Run("TestSMTPSendProbeResources",
		func(st *testing.T) {
			it.Test = st
			it.TestSMTPSendProbeResources(zoneName, ownerNameSB)
		})
}

func (t *IntegrationTest) CreateProbe(rrSetKey *rrset.RRSetKey, probeData *probe.Probe) {
	probeService, err := probe.Get(integration.TestClient)

	if err != nil {
		t.Test.Fatal(err)
	}

	if _, er := probeService.Create(rrSetKey, probeData); er != nil {
		t.Test.Fatal(er)
	}
}

func (t *IntegrationTest) UpdateProbe(rrSetKey *rrset.RRSetKey, probeData *probe.Probe) {
	probeService, err := probe.Get(integration.TestClient)

	if err != nil {
		t.Test.Fatal(err)
	}

	rrSetKey.ID = probeID

	if _, er := probeService.Update(rrSetKey, probeData); er != nil {
		t.Test.Fatal(er)
	}
}

func (t *IntegrationTest) PartialUpdateProbe(rrSetKey *rrset.RRSetKey, probeData *probe.Probe) {
	probeService, err := probe.Get(integration.TestClient)

	if err != nil {
		t.Test.Fatal(err)
	}

	rrSetKey.ID = probeID

	if _, er := probeService.PartialUpdate(rrSetKey, probeData); er != nil {
		t.Test.Fatal(er)
	}
}

func (t *IntegrationTest) ReadProbe(rrSetKey *rrset.RRSetKey) {
	probeService, err := probe.Get(integration.TestClient)

	if err != nil {
		t.Test.Fatal(err)
	}

	rrSetKey.ID = probeID

	if _, _, er := probeService.Read(rrSetKey); er != nil {
		t.Test.Fatal(er)
	}
}

func (t *IntegrationTest) DeleteProbe(rrSetKey *rrset.RRSetKey) {
	probeService, err := probe.Get(integration.TestClient)

	if err != nil {
		t.Test.Fatal(err)
	}

	rrSetKey.ID = probeID

	if _, er := probeService.Delete(rrSetKey); er != nil {
		t.Test.Fatal(er)
	}
}

func (t *IntegrationTest) ListProbe(rrSetKey *rrset.RRSetKey) {
	probeService, err := probe.Get(integration.TestClient)

	if err != nil {
		t.Test.Fatal(err)
	}

	query := &probe.Query{}

	_, probesList, er := probeService.List(rrSetKey, query)

	if er != nil {
		t.Test.Fatal(er)
	}

	probeID = probesList.Probes[0].ID
}

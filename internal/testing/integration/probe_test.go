package integration_test

import (
	"github.com/ultradns/ultradns-go-sdk/internal/testing/integration"
	"github.com/ultradns/ultradns-go-sdk/pkg/probe"
	"github.com/ultradns/ultradns-go-sdk/pkg/rrset"
)

var probeID = ""

func (it *IntegrationTest) CreateProbe(rrSetKey *rrset.RRSetKey, probeData *probe.Probe) {
	probeService, err := probe.Get(integration.TestClient)

	if err != nil {
		it.Test.Fatal(err)
	}

	if _, er := probeService.Create(rrSetKey, probeData); er != nil {
		it.Test.Fatal(er)
	}
}

func (it *IntegrationTest) UpdateProbe(rrSetKey *rrset.RRSetKey, probeData *probe.Probe) {
	probeService, err := probe.Get(integration.TestClient)

	if err != nil {
		it.Test.Fatal(err)
	}

	rrSetKey.ID = probeID

	if _, er := probeService.Update(rrSetKey, probeData); er != nil {
		it.Test.Fatal(er)
	}
}

func (it *IntegrationTest) PartialUpdateProbe(rrSetKey *rrset.RRSetKey, probeData *probe.Probe) {
	probeService, err := probe.Get(integration.TestClient)

	if err != nil {
		it.Test.Fatal(err)
	}

	rrSetKey.ID = probeID

	if _, er := probeService.PartialUpdate(rrSetKey, probeData); er != nil {
		it.Test.Fatal(er)
	}
}

func (it *IntegrationTest) ReadProbe(rrSetKey *rrset.RRSetKey) {
	probeService, err := probe.Get(integration.TestClient)

	if err != nil {
		it.Test.Fatal(err)
	}

	rrSetKey.ID = probeID

	if _, _, er := probeService.Read(rrSetKey); er != nil {
		it.Test.Fatal(er)
	}
}

func (it *IntegrationTest) DeleteProbe(rrSetKey *rrset.RRSetKey) {
	probeService, err := probe.Get(integration.TestClient)

	if err != nil {
		it.Test.Fatal(err)
	}

	rrSetKey.ID = probeID

	if _, er := probeService.Delete(rrSetKey); er != nil {
		it.Test.Fatal(er)
	}
}

func (it *IntegrationTest) ListProbe(rrSetKey *rrset.RRSetKey) {
	probeService, err := probe.Get(integration.TestClient)

	if err != nil {
		it.Test.Fatal(err)
	}

	_, probesList, er := probeService.List(rrSetKey)

	if er != nil {
		it.Test.Fatal(er)
	}

	probeID = probesList.Probes[0].ID
}

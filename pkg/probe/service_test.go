package probe_test

import (
	"testing"

	"github.com/ultradns/ultradns-go-sdk/internal/testing/integration"
	"github.com/ultradns/ultradns-go-sdk/pkg/probe"
	"github.com/ultradns/ultradns-go-sdk/pkg/probe/http"
	"github.com/ultradns/ultradns-go-sdk/pkg/rrset"
)

const serviceErrorString = "Probe service is not properly configured"

func TestNewSuccess(t *testing.T) {
	conf := integration.GetConfig()

	if _, err := probe.New(conf); err != nil {
		t.Fatal(err)
	}
}

func TestNewError(t *testing.T) {
	conf := integration.GetConfig()
	conf.Password = ""

	if _, err := probe.New(conf); err.Error() != "config error while creating Probe service : config validation failure: password is missing" {
		t.Fatal(err)
	}
}

func TestGetError(t *testing.T) {
	if _, err := probe.Get(nil); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestCreateProbeWithConfigError(t *testing.T) {
	probeService := probe.Service{}
	if _, err := probeService.Create(&rrset.RRSetKey{}, &probe.Probe{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestUpdateProbeWithConfigError(t *testing.T) {
	probeService := probe.Service{}

	if _, err := probeService.Update(&rrset.RRSetKey{}, &probe.Probe{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestPartialUpdateProbeWithConfigError(t *testing.T) {
	probeService := probe.Service{}

	if _, err := probeService.PartialUpdate(&rrset.RRSetKey{}, &probe.Probe{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestReadProbeWithConfigError(t *testing.T) {
	probeService := probe.Service{}

	if _, _, err := probeService.Read(&rrset.RRSetKey{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestDeleteProbeWithConfigError(t *testing.T) {
	probeService := probe.Service{}

	if _, err := probeService.Delete(&rrset.RRSetKey{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestListProbeWithConfigError(t *testing.T) {
	probeService := probe.Service{}

	if _, _, err := probeService.List(&rrset.RRSetKey{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestCreateProbeFailure(t *testing.T) {
	probeService, err := probe.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	rrSetKey := integration.GetTestRRSetKey()
	rrSetKey.ID = ""

	if _, er := probeService.Create(rrSetKey, testGetHTTPProbe()); er.Error() != "error while creating Probe - www.non-existing-zone.com.:non-existing-zone.com.:A (1): : error from api response - error code : 53006 - error message : Agents must not be empty." {
		t.Fatal(er)
	}
}

func TestUpdateProbeFailure(t *testing.T) {
	probeService, err := probe.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, er := probeService.Update(integration.GetTestRRSetKey(), testGetHTTPProbe()); er.Error() != "error while updating Probe - www.non-existing-zone.com.:non-existing-zone.com.:A (1):id : error from api response - error code : 53006 - error message : Agents must not be empty." {
		t.Fatal(er)
	}
}

func TestPartialUpdateProbeFailure(t *testing.T) {
	probeService, err := probe.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, er := probeService.PartialUpdate(integration.GetTestRRSetKey(), testGetHTTPProbe()); er.Error() != "error while partial updating Probe - www.non-existing-zone.com.:non-existing-zone.com.:A (1):id : error from api response - error code : 2911 - error message : Pool does not exist in the system" {
		t.Fatal(er)
	}
}

func TestReadProbeFailure(t *testing.T) {
	probeService, err := probe.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, _, er := probeService.Read(integration.GetTestRRSetKey()); er.Error() != "error while reading Probe - www.non-existing-zone.com.:non-existing-zone.com.:A (1):id : error from api response - error code : 2911 - error message : Pool does not exist in the system" {
		t.Fatal(er)
	}
}

func TestDeleteProbeFailure(t *testing.T) {
	probeService, err := probe.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, er := probeService.Delete(integration.GetTestRRSetKey()); er.Error() != "error while deleting Probe - www.non-existing-zone.com.:non-existing-zone.com.:A (1):id : error from api response - error code : 2911 - error message : Pool does not exist in the system" {
		t.Fatal(er)
	}
}

func TestListProbeFailure(t *testing.T) {
	probeService, err := probe.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	rrSetKey := integration.GetTestRRSetKey()
	rrSetKey.ID = ""

	if _, _, er := probeService.List(rrSetKey); er.Error() != "error while listing Probe : uri - zones/non-existing-zone.com./rrsets/A/www/probes/ : error from api response - error code : 2911 - error message : Pool does not exist in the system" {
		t.Fatal(er)
	}
}

func TestCreateProbeValidationFailure(t *testing.T) {
	probeService, err := probe.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	rrSetKey := integration.GetTestRRSetKey()
	rrSetKey.ID = ""

	rrSet := testGetHTTPProbe()
	rrSet.Type = probe.FTP

	if _, er := probeService.Create(rrSetKey, rrSet); er.Error() != "type mismatched : expected - *ftp.Details : found - *http.Details" {
		t.Fatal(er)
	}
}

func TestUpdateProbeValidationFailure(t *testing.T) {
	probeService, err := probe.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	rrSet := testGetHTTPProbe()
	rrSet.Type = probe.FTP

	if _, er := probeService.Update(integration.GetTestRRSetKey(), rrSet); er.Error() != "type mismatched : expected - *ftp.Details : found - *http.Details" {
		t.Fatal(er)
	}
}

func TestRRSetKeyURI(t *testing.T) {
	rrSetKey := rrset.RRSetKey{
		Zone:  "a",
		Owner: "b",
	}

	if expectedURI, foundURI := "zones/a/rrsets/ANY/b/probes/", rrSetKey.ProbeURI(); expectedURI != foundURI {
		t.Fatalf("uri mismatched expected - %s : found - %s", expectedURI, foundURI)
	}
}

func TestRRSetKeyID(t *testing.T) {
	rrSetKey := rrset.RRSetKey{
		Zone:       "example.com",
		Owner:      "www",
		RecordType: "A",
		ID:         "id",
	}

	if expectedID, foundID := "www.example.com.:example.com.:A (1):id", rrSetKey.PID(); expectedID != foundID {
		t.Fatalf("rrset id mismatched expected - %s : found - %s", expectedID, foundID)
	}
}

func testGetHTTPProbe() *probe.Probe {
	probedata := &probe.Probe{}
	probedata.Type = probe.HTTP
	probedata.Details = &http.Details{}

	return probedata
}

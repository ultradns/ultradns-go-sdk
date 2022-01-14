package rdpool_test

import (
	"testing"

	"github.com/ultradns/ultradns-go-sdk/internal/test/integration"
	"github.com/ultradns/ultradns-go-sdk/pkg/rdpool"
	"github.com/ultradns/ultradns-go-sdk/pkg/rrset"
)

const serviceErrorString = "RD-Pool service is not properly configured"

var testRRSetKey = &rrset.RRSetKey{
	Name: "www",
	Zone: "non-existing-zone.com.",
	Type: "A",
}

func TestNewSuccess(t *testing.T) {
	conf := integration.GetConfig()

	if _, err := rdpool.New(conf); err != nil {
		t.Fatal(err)
	}
}

func TestNewError(t *testing.T) {
	conf := integration.GetConfig()
	conf.Password = ""

	if _, err := rdpool.New(conf); err.Error() != "config error while creating RD-Pool service : config validation failure: password is missing" {
		t.Fatal(err)
	}
}

func TestGetError(t *testing.T) {
	if _, err := rdpool.Get(nil); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestCreateRDPoolWithConfigError(t *testing.T) {
	rdPoolService := rdpool.Service{}
	if _, err := rdPoolService.CreateRDPool(&rrset.RRSetKey{}, &rrset.RRSet{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestUpdateRDPoolWithConfigError(t *testing.T) {
	rdPoolService := rdpool.Service{}

	if _, err := rdPoolService.UpdateRDPool(&rrset.RRSetKey{}, &rrset.RRSet{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestPartialUpdateRDPoolWithConfigError(t *testing.T) {
	rdPoolService := rdpool.Service{}

	if _, err := rdPoolService.PartialUpdateRDPool(&rrset.RRSetKey{}, &rrset.RRSet{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestReadRDPoolWithConfigError(t *testing.T) {
	rdPoolService := rdpool.Service{}

	if _, _, err := rdPoolService.ReadRDPool(&rrset.RRSetKey{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestDeleteRDPoolWithConfigError(t *testing.T) {
	rdPoolService := rdpool.Service{}

	if _, err := rdPoolService.DeleteRDPool(&rrset.RRSetKey{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestCreateRDPoolFailure(t *testing.T) {
	rdPoolService, err := rdpool.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	profile := &rdpool.Profile{
		Order: "RANDOM",
	}
	rrSet := &rrset.RRSet{
		Profile: profile,
	}

	if _, er := rdPoolService.CreateRDPool(testRRSetKey, rrSet); er.Error() != "error while creating RD-Pool - www.non-existing-zone.com.:non-existing-zone.com.:A (1) : error from api response - error code : 1801 - error message : Zone does not exist in the system." {
		t.Fatal(er)
	}
}

func TestUpdateRDPoolFailure(t *testing.T) {
	rdPoolService, err := rdpool.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	profile := &rdpool.Profile{
		Order: "ROUND_ROBIN",
	}
	rrSet := &rrset.RRSet{
		Profile: profile,
	}

	if _, er := rdPoolService.UpdateRDPool(testRRSetKey, rrSet); er.Error() != "error while updating RD-Pool - www.non-existing-zone.com.:non-existing-zone.com.:A (1) : error from api response - error code : 1801 - error message : Zone does not exist in the system." {
		t.Fatal(er)
	}
}

func TestPartialUpdateRDPoolFailure(t *testing.T) {
	rdPoolService, err := rdpool.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	profile := &rdpool.Profile{
		Order: "FIXED",
	}
	rrSet := &rrset.RRSet{
		Profile: profile,
	}

	if _, er := rdPoolService.PartialUpdateRDPool(testRRSetKey, rrSet); er.Error() != "error while partial updating RD-Pool - www.non-existing-zone.com.:non-existing-zone.com.:A (1) : error from api response - error code : 1801 - error message : Zone does not exist in the system." {
		t.Fatal(er)
	}
}

func TestReadRDPoolFailure(t *testing.T) {
	rdPoolService, err := rdpool.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, _, er := rdPoolService.ReadRDPool(testRRSetKey); er.Error() != "error while reading RD-Pool - www.non-existing-zone.com.:non-existing-zone.com.:A (1) : error from api response - error code : 1801 - error message : Zone does not exist in the system." {
		t.Fatal(er)
	}
}

func TestDeleteRDPoolFailure(t *testing.T) {
	rdPoolService, err := rdpool.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, er := rdPoolService.DeleteRDPool(testRRSetKey); er.Error() != "error while deleting RD-Pool - www.non-existing-zone.com.:non-existing-zone.com.:A (1) : error from api response - error code : 1801 - error message : Zone does not exist in the system." {
		t.Fatal(er)
	}
}

func TestCreateRDPoolWithValidationFailure(t *testing.T) {
	rdPoolService, err := rdpool.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, er := rdPoolService.CreateRDPool(testRRSetKey, &rrset.RRSet{}); er.Error() != "type mismatched : expected - *rdpool.Profile : found - <nil>" {
		t.Fatal(er)
	}
}

func TestUpdateRDPoolWithValidationFailure(t *testing.T) {
	rdPoolService, err := rdpool.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	profile := &rdpool.Profile{
		Order: "TEST",
	}
	rrSet := &rrset.RRSet{
		Profile: profile,
	}

	if _, er := rdPoolService.UpdateRDPool(testRRSetKey, rrSet); er.Error() != "RD-Pool order should be any of the following data [FIXED RANDOM ROUND_ROBIN] : found - TEST" {
		t.Fatal(er)
	}
}

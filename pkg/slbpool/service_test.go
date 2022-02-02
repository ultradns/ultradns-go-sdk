package slbpool_test

import (
	"testing"

	"github.com/ultradns/ultradns-go-sdk/internal/testing/integration"
	"github.com/ultradns/ultradns-go-sdk/pkg/pool"
	"github.com/ultradns/ultradns-go-sdk/pkg/rrset"
	"github.com/ultradns/ultradns-go-sdk/pkg/slbpool"
)

const (
	serviceErrorString = "SLB-Pool service is not properly configured"
	testString         = "TEST"
)

func TestNewSuccess(t *testing.T) {
	conf := integration.GetConfig()

	if _, err := slbpool.New(conf); err != nil {
		t.Fatal(err)
	}
}

func TestNewError(t *testing.T) {
	conf := integration.GetConfig()
	conf.Password = ""

	if _, err := slbpool.New(conf); err.Error() != "config error while creating SLB-Pool service : config validation failure: password is missing" {
		t.Fatal(err)
	}
}

func TestGetError(t *testing.T) {
	if _, err := slbpool.Get(nil); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestCreateSLBPoolWithConfigError(t *testing.T) {
	slbPoolService := slbpool.Service{}
	if _, err := slbPoolService.CreateSLBPool(&rrset.RRSetKey{}, &rrset.RRSet{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestUpdateSLBPoolWithConfigError(t *testing.T) {
	slbPoolService := slbpool.Service{}

	if _, err := slbPoolService.UpdateSLBPool(&rrset.RRSetKey{}, &rrset.RRSet{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestPartialUpdateSLBPoolWithConfigError(t *testing.T) {
	slbPoolService := slbpool.Service{}

	if _, err := slbPoolService.PartialUpdateSLBPool(&rrset.RRSetKey{}, &rrset.RRSet{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestReadSLBPoolWithConfigError(t *testing.T) {
	slbPoolService := slbpool.Service{}

	if _, _, err := slbPoolService.ReadSLBPool(&rrset.RRSetKey{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestDeleteSLBPoolWithConfigError(t *testing.T) {
	slbPoolService := slbpool.Service{}

	if _, err := slbPoolService.DeleteSLBPool(&rrset.RRSetKey{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestCreateSLBPoolFailure(t *testing.T) {
	slbPoolService, err := slbpool.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	profile := getSLBPoolProfile()
	rrSet := &rrset.RRSet{
		Profile: profile,
		RData:   []string{"192.168.1.1"},
	}

	if _, er := slbPoolService.CreateSLBPool(integration.TestRRSetKey, rrSet); er.Error() != "error while creating SLB-Pool - www.non-existing-zone.com.:non-existing-zone.com.:A (1) : error from api response - error code : 1801 - error message : Zone does not exist in the system." {
		t.Fatal(er)
	}
}

func TestUpdateSLBPoolFailure(t *testing.T) {
	slbPoolService, err := slbpool.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	profile := getSLBPoolProfile()
	rrSet := &rrset.RRSet{
		Profile: profile,
		RData:   []string{"192.168.1.1"},
	}

	if _, er := slbPoolService.UpdateSLBPool(integration.TestRRSetKey, rrSet); er.Error() != "error while updating SLB-Pool - www.non-existing-zone.com.:non-existing-zone.com.:A (1) : error from api response - error code : 1801 - error message : Zone does not exist in the system." {
		t.Fatal(er)
	}
}

func TestPartialUpdateSLBPoolFailure(t *testing.T) {
	slbPoolService, err := slbpool.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	profile := getSLBPoolProfile()
	rrSet := &rrset.RRSet{
		Profile: profile,
	}

	if _, er := slbPoolService.PartialUpdateSLBPool(integration.TestRRSetKey, rrSet); er.Error() != "error while partial updating SLB-Pool - www.non-existing-zone.com.:non-existing-zone.com.:A (1) : error from api response - error code : 1801 - error message : Zone does not exist in the system." {
		t.Fatal(er)
	}
}

func TestReadSLBPoolFailure(t *testing.T) {
	slbPoolService, err := slbpool.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, _, er := slbPoolService.ReadSLBPool(integration.TestRRSetKey); er.Error() != "error while reading SLB-Pool - www.non-existing-zone.com.:non-existing-zone.com.:A (1) : error from api response - error code : 1801 - error message : Zone does not exist in the system." {
		t.Fatal(er)
	}
}

func TestDeleteSLBPoolFailure(t *testing.T) {
	slbPoolService, err := slbpool.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, er := slbPoolService.DeleteSLBPool(integration.TestRRSetKey); er.Error() != "error while deleting SLB-Pool - www.non-existing-zone.com.:non-existing-zone.com.:A (1) : error from api response - error code : 1801 - error message : Zone does not exist in the system." {
		t.Fatal(er)
	}
}

func TestCreateSLBPoolWithValidationFailure(t *testing.T) {
	slbPoolService, err := slbpool.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, er := slbPoolService.CreateSLBPool(integration.TestRRSetKey, &rrset.RRSet{}); er.Error() != "type mismatched : expected - *slbpool.Profile : found - <nil>" {
		t.Fatal(er)
	}
}

func TestCreateSLBPoolWithResponseMethodValidationFailure(t *testing.T) {
	slbPoolService, err := slbpool.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	profile := getSLBPoolProfile()
	profile.ResponseMethod = testString
	rrSet := &rrset.RRSet{
		Profile: profile,
	}

	if _, er := slbPoolService.CreateSLBPool(integration.TestRRSetKey, rrSet); er.Error() != "SLB-Pool Response Method should be any of the following data [PRIORITY_HUNT RANDOM ROUND_ROBIN] : found - TEST" {
		t.Fatal(er)
	}
}

func TestCreateSLBPoolWithServingPreferenceValidationFailure(t *testing.T) {
	slbPoolService, err := slbpool.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	profile := getSLBPoolProfile()
	profile.ServingPreference = testString
	rrSet := &rrset.RRSet{
		Profile: profile,
	}

	if _, er := slbPoolService.CreateSLBPool(integration.TestRRSetKey, rrSet); er.Error() != "SLB-Pool Serving Preference should be any of the following data [AUTO_SELECT SERVE_PRIMARY SERVE_ALL_FAIL] : found - TEST" {
		t.Fatal(er)
	}
}

func TestUpdateSLBPoolWithRegionSensitivityValidationFailure(t *testing.T) {
	slbPoolService, err := slbpool.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	profile := getSLBPoolProfile()
	profile.RegionFailureSensitivity = testString
	rrSet := &rrset.RRSet{
		Profile: profile,
	}

	if _, er := slbPoolService.UpdateSLBPool(integration.TestRRSetKey, rrSet); er.Error() != "SLB-Pool Region Failure Sensitivity should be any of the following data [HIGH LOW] : found - TEST" {
		t.Fatal(er)
	}
}

func TestUpdateSLBPoolWithMethodValidationFailure(t *testing.T) {
	slbPoolService, err := slbpool.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	profile := getSLBPoolProfile()
	profile.Monitor.Method = testString
	rrSet := &rrset.RRSet{
		Profile: profile,
	}

	if _, er := slbPoolService.UpdateSLBPool(integration.TestRRSetKey, rrSet); er.Error() != "SLB-Pool Monitor Method should be any of the following data [GET POST] : found - TEST" {
		t.Fatal(er)
	}
}

func getSLBPoolProfile() *slbpool.Profile {
	rdataInfo := &slbpool.RDataInfo{
		ProbingEnabled: false,
	}
	allFailRecord := &slbpool.AllFailRecord{
		RData:   "192.168.1.1",
		Serving: false,
	}
	monitor := &pool.Monitor{
		Method: "GET",
		URL:    integration.TestHost,
	}

	return &slbpool.Profile{
		RDataInfo:                []*slbpool.RDataInfo{rdataInfo},
		Monitor:                  monitor,
		AllFailRecord:            allFailRecord,
		RegionFailureSensitivity: "HIGH",
		ServingPreference:        "AUTO_SELECT",
		ResponseMethod:           "ROUND_ROBIN",
	}
}

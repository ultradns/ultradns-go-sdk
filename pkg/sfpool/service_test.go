package sfpool_test

import (
	"testing"

	"github.com/ultradns/ultradns-go-sdk/internal/testing/integration"
	"github.com/ultradns/ultradns-go-sdk/pkg/pool"
	"github.com/ultradns/ultradns-go-sdk/pkg/rrset"
	"github.com/ultradns/ultradns-go-sdk/pkg/sfpool"
)

const serviceErrorString = "SF-Pool service is not properly configured"

func TestNewSuccess(t *testing.T) {
	conf := integration.GetConfig()

	if _, err := sfpool.New(conf); err != nil {
		t.Fatal(err)
	}
}

func TestNewError(t *testing.T) {
	conf := integration.GetConfig()
	conf.Password = ""

	if _, err := sfpool.New(conf); err.Error() != "config error while creating SF-Pool service : config validation failure: password is missing" {
		t.Fatal(err)
	}
}

func TestGetError(t *testing.T) {
	if _, err := sfpool.Get(nil); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestCreateSFPoolWithConfigError(t *testing.T) {
	sfPoolService := sfpool.Service{}
	if _, err := sfPoolService.CreateSFPool(&rrset.RRSetKey{}, &rrset.RRSet{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestUpdateSFPoolWithConfigError(t *testing.T) {
	sfPoolService := sfpool.Service{}

	if _, err := sfPoolService.UpdateSFPool(&rrset.RRSetKey{}, &rrset.RRSet{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestPartialUpdateSFPoolWithConfigError(t *testing.T) {
	sfPoolService := sfpool.Service{}

	if _, err := sfPoolService.PartialUpdateSFPool(&rrset.RRSetKey{}, &rrset.RRSet{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestReadSFPoolWithConfigError(t *testing.T) {
	sfPoolService := sfpool.Service{}

	if _, _, err := sfPoolService.ReadSFPool(&rrset.RRSetKey{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestDeleteSFPoolWithConfigError(t *testing.T) {
	sfPoolService := sfpool.Service{}

	if _, err := sfPoolService.DeleteSFPool(&rrset.RRSetKey{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestCreateSFPoolFailure(t *testing.T) {
	sfPoolService, err := sfpool.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	profile := getSFPoolProfile()
	rrSet := &rrset.RRSet{
		Profile: profile,
		RData:   []string{"192.168.1.1"},
	}

	if _, er := sfPoolService.CreateSFPool(integration.TestRRSetKey, rrSet); er.Error() != "error while creating SF-Pool - www.non-existing-zone.com.:non-existing-zone.com.:A (1) : error from api response - error code : 1801 - error message : Zone does not exist in the system." {
		t.Fatal(er)
	}
}

func TestUpdateSFPoolFailure(t *testing.T) {
	sfPoolService, err := sfpool.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	profile := getSFPoolProfile()
	rrSet := &rrset.RRSet{
		Profile: profile,
		RData:   []string{"192.168.1.1"},
	}

	if _, er := sfPoolService.UpdateSFPool(integration.TestRRSetKey, rrSet); er.Error() != "error while updating SF-Pool - www.non-existing-zone.com.:non-existing-zone.com.:A (1) : error from api response - error code : 1801 - error message : Zone does not exist in the system." {
		t.Fatal(er)
	}
}

func TestPartialUpdateSFPoolFailure(t *testing.T) {
	sfPoolService, err := sfpool.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	profile := getSFPoolProfile()
	rrSet := &rrset.RRSet{
		Profile: profile,
	}

	if _, er := sfPoolService.PartialUpdateSFPool(integration.TestRRSetKey, rrSet); er.Error() != "error while partial updating SF-Pool - www.non-existing-zone.com.:non-existing-zone.com.:A (1) : error from api response - error code : 1801 - error message : Zone does not exist in the system." {
		t.Fatal(er)
	}
}

func TestReadSFPoolFailure(t *testing.T) {
	sfPoolService, err := sfpool.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, _, er := sfPoolService.ReadSFPool(integration.TestRRSetKey); er.Error() != "error while reading SF-Pool - www.non-existing-zone.com.:non-existing-zone.com.:A (1) : error from api response - error code : 1801 - error message : Zone does not exist in the system." {
		t.Fatal(er)
	}
}

func TestDeleteSFPoolFailure(t *testing.T) {
	sfPoolService, err := sfpool.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, er := sfPoolService.DeleteSFPool(integration.TestRRSetKey); er.Error() != "error while deleting SF-Pool - www.non-existing-zone.com.:non-existing-zone.com.:A (1) : error from api response - error code : 1801 - error message : Zone does not exist in the system." {
		t.Fatal(er)
	}
}

func TestCreateSFPoolWithValidationFailure(t *testing.T) {
	sfPoolService, err := sfpool.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, er := sfPoolService.CreateSFPool(integration.TestRRSetKey, &rrset.RRSet{}); er.Error() != "type mismatched : expected - *sfpool.Profile : found - <nil>" {
		t.Fatal(er)
	}
}

func TestUpdateSFPoolWithRegionSensitivityValidationFailure(t *testing.T) {
	sfPoolService, err := sfpool.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	profile := getSFPoolProfile()
	profile.RegionFailureSensitivity = "TEST"
	rrSet := &rrset.RRSet{
		Profile: profile,
	}

	if _, er := sfPoolService.UpdateSFPool(integration.TestRRSetKey, rrSet); er.Error() != "SF-Pool Region Failure Sensitivity should be any of the following data [HIGH LOW] : found - TEST" {
		t.Fatal(er)
	}
}

func TestUpdateSFPoolWithMethodValidationFailure(t *testing.T) {
	sfPoolService, err := sfpool.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	profile := getSFPoolProfile()
	profile.Monitor.Method = "TEST"
	rrSet := &rrset.RRSet{
		Profile: profile,
	}

	if _, er := sfPoolService.UpdateSFPool(integration.TestRRSetKey, rrSet); er.Error() != "SF-Pool Monitor Method should be any of the following data [GET POST] : found - TEST" {
		t.Fatal(er)
	}
}

func getSFPoolProfile() *sfpool.Profile {
	backupRecord := &pool.BackupRecord{
		RData: "192.168.1.1",
	}
	monitor := &pool.Monitor{
		Method: "GET",
		URL:    integration.TestHost,
	}

	return &sfpool.Profile{
		BackupRecord:             backupRecord,
		Monitor:                  monitor,
		RegionFailureSensitivity: "HIGH",
	}
}

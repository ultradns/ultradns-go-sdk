package record_test

import (
	"testing"

	"github.com/ultradns/ultradns-go-sdk/internal/testing/integration"
	"github.com/ultradns/ultradns-go-sdk/pkg/record"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/dirpool"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/pool"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/rdpool"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/sbpool"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/sfpool"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/slbpool"
	"github.com/ultradns/ultradns-go-sdk/pkg/record/tcpool"
	"github.com/ultradns/ultradns-go-sdk/pkg/rrset"
)

const serviceErrorString = "Record service is not properly configured"

const (
	testString         = "TEST"
	testPoolOrderError = "Pool order should be any of the following data [FIXED RANDOM ROUND_ROBIN] : found - TEST"
)

func TestNewSuccess(t *testing.T) {
	conf := integration.GetConfig()

	if _, err := record.New(conf); err != nil {
		t.Fatal(err)
	}
}

func TestNewError(t *testing.T) {
	conf := integration.GetConfig()
	conf.Password = ""

	if _, err := record.New(conf); err.Error() != "config error while creating Record service : config validation failure: password is missing" {
		t.Fatal(err)
	}
}

func TestGetError(t *testing.T) {
	if _, err := record.Get(nil); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestCreateRecordWithConfigError(t *testing.T) {
	recordService := record.Service{}
	if _, err := recordService.Create(&rrset.RRSetKey{}, &rrset.RRSet{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestUpdateRecordWithConfigError(t *testing.T) {
	recordService := record.Service{}

	if _, err := recordService.Update(&rrset.RRSetKey{}, &rrset.RRSet{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestPartialUpdateRecordWithConfigError(t *testing.T) {
	recordService := record.Service{}

	if _, err := recordService.PartialUpdate(&rrset.RRSetKey{}, &rrset.RRSet{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestReadRecordWithConfigError(t *testing.T) {
	recordService := record.Service{}

	if _, _, err := recordService.Read(&rrset.RRSetKey{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestDeleteRecordWithConfigError(t *testing.T) {
	recordService := record.Service{}

	if _, err := recordService.Delete(&rrset.RRSetKey{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestCreateRecordFailure(t *testing.T) {
	recordService, err := record.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, er := recordService.Create(integration.GetTestRRSetKey(), &rrset.RRSet{}); er.Error() != "error while creating Record - www.non-existing-zone.com.:non-existing-zone.com.:A (1) : error from api response - error code : 70005 - error message : At least one field must be specified: rdata or profile" {
		t.Fatal(er)
	}
}

func TestUpdateRecordFailure(t *testing.T) {
	recordService, err := record.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, er := recordService.Update(integration.GetTestRRSetKey(), &rrset.RRSet{}); er.Error() != "error while updating Record - www.non-existing-zone.com.:non-existing-zone.com.:A (1) : error from api response - error code : 70005 - error message : At least one field must be specified: rdata or profile" {
		t.Fatal(er)
	}
}

func TestPartialUpdateRecordFailure(t *testing.T) {
	recordService, err := record.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, er := recordService.PartialUpdate(integration.GetTestRRSetKey(), &rrset.RRSet{}); er.Error() != "error while partial updating Record - www.non-existing-zone.com.:non-existing-zone.com.:A (1) : error from api response - error code : 1801 - error message : Zone does not exist in the system." {
		t.Fatal(er)
	}
}

func TestReadRecordFailure(t *testing.T) {
	recordService, err := record.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, _, er := recordService.Read(integration.GetTestRRSetKey()); er.Error() != "error while reading Record - www.non-existing-zone.com.:non-existing-zone.com.:A (1) : error from api response - error code : 1801 - error message : Zone does not exist in the system." {
		t.Fatal(er)
	}
}

func TestDeleteRecordFailure(t *testing.T) {
	recordService, err := record.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, er := recordService.Delete(integration.GetTestRRSetKey()); er.Error() != "error while deleting Record - www.non-existing-zone.com.:non-existing-zone.com.:A (1) : error from api response - error code : 1801 - error message : Zone does not exist in the system." {
		t.Fatal(er)
	}
}

func TestUpdateWithValidationFailure(t *testing.T) {
	recordService, err := record.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	profile := &rdpool.Profile{
		Order: testString,
	}
	rrSet := &rrset.RRSet{
		Profile: profile,
	}

	if _, er := recordService.Update(integration.GetTestRRSetKey(), rrSet); er.Error() != testPoolOrderError {
		t.Fatal(er)
	}
}

func TestRRSetKeyURI(t *testing.T) {
	rrSetKey := rrset.RRSetKey{
		Zone:  "a",
		Owner: "b",
	}

	if expectedURI, foundURI := "zones/a/rrsets/ANY/b", rrSetKey.RecordURI(); expectedURI != foundURI {
		t.Fatalf("uri mismatched expected - %s : found - %s", expectedURI, foundURI)
	}
}

func TestRRSetKeyID(t *testing.T) {
	rrSetKey := rrset.RRSetKey{
		Zone:       "example.com",
		Owner:      "www",
		RecordType: "A",
	}

	if expectedID, foundID := "www.example.com.:example.com.:A (1)", rrSetKey.RecordID(); expectedID != foundID {
		t.Fatalf("rrset id mismatched expected - %s : found - %s", expectedID, foundID)
	}
}

func TestCaseInsensitiveRRSetKeyID(t *testing.T) {
	rrSetKey := rrset.RRSetKey{
		Zone:       "EXAMPLE.com",
		Owner:      "wWw",
		RecordType: "A",
	}

	if expectedID, foundID := "www.example.com.:example.com.:A (1)", rrSetKey.RecordID(); expectedID != foundID {
		t.Fatalf("rrset id mismatched expected - %s : found - %s", expectedID, foundID)
	}
}

func TestRDPoolOrderValidationFailure(t *testing.T) {
	recordService, err := record.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	profile := &rdpool.Profile{
		Order: testString,
	}
	rrSet := &rrset.RRSet{
		Profile: profile,
	}

	if _, er := recordService.Create(integration.GetTestRRSetKey(), rrSet); er.Error() != testPoolOrderError {
		t.Fatal(er)
	}
}

func TestSFPoolMonitorMethodValidationFailure(t *testing.T) {
	recordService, err := record.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	monitor := &pool.Monitor{
		Method: testString,
	}
	profile := &sfpool.Profile{
		Monitor: monitor,
	}
	rrSet := &rrset.RRSet{
		Profile: profile,
	}

	if _, er := recordService.Create(integration.GetTestRRSetKey(), rrSet); er.Error() != "Pool Monitor Method should be any of the following data [GET POST] : found - TEST" {
		t.Fatal(er)
	}
}

func TestSFPoolRegionFailureSensitivityValidationFailure(t *testing.T) {
	recordService, err := record.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	monitor := &pool.Monitor{
		Method: "GET",
	}
	profile := &sfpool.Profile{
		Monitor:                  monitor,
		RegionFailureSensitivity: testString,
	}
	rrSet := &rrset.RRSet{
		Profile: profile,
	}

	if _, er := recordService.Create(integration.GetTestRRSetKey(), rrSet); er.Error() != "Pool Region Failure Sensitivity should be any of the following data [HIGH LOW] : found - TEST" {
		t.Fatal(er)
	}
}

func TestSLBPoolMonitorMethodValidationFailure(t *testing.T) {
	recordService, err := record.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	monitor := &pool.Monitor{
		Method: testString,
	}
	profile := &slbpool.Profile{
		Monitor: monitor,
	}
	rrSet := &rrset.RRSet{
		Profile: profile,
	}

	if _, er := recordService.Create(integration.GetTestRRSetKey(), rrSet); er.Error() != "Pool Monitor Method should be any of the following data [GET POST] : found - TEST" {
		t.Fatal(er)
	}
}

func TestSLBPoolRegionFailureSensitivityValidationFailure(t *testing.T) {
	recordService, err := record.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	monitor := &pool.Monitor{
		Method: "GET",
	}
	profile := &slbpool.Profile{
		Monitor:                  monitor,
		RegionFailureSensitivity: testString,
	}
	rrSet := &rrset.RRSet{
		Profile: profile,
	}

	if _, er := recordService.Create(integration.GetTestRRSetKey(), rrSet); er.Error() != "Pool Region Failure Sensitivity should be any of the following data [HIGH LOW] : found - TEST" {
		t.Fatal(er)
	}
}

func TestSLBPoolResponseMethodValidationFailure(t *testing.T) {
	recordService, err := record.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	monitor := &pool.Monitor{
		Method: "GET",
	}
	profile := &slbpool.Profile{
		Monitor:                  monitor,
		RegionFailureSensitivity: "HIGH",
		ResponseMethod:           testString,
	}
	rrSet := &rrset.RRSet{
		Profile: profile,
	}

	if _, er := recordService.Create(integration.GetTestRRSetKey(), rrSet); er.Error() != "Pool Response Method should be any of the following data [PRIORITY_HUNT RANDOM ROUND_ROBIN] : found - TEST" {
		t.Fatal(er)
	}
}

func TestSLBPoolServingPreferenceValidationFailure(t *testing.T) {
	recordService, err := record.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	monitor := &pool.Monitor{
		Method: "GET",
	}
	profile := &slbpool.Profile{
		Monitor:                  monitor,
		RegionFailureSensitivity: "HIGH",
		ResponseMethod:           "ROUND_ROBIN",
		ServingPreference:        testString,
	}
	rrSet := &rrset.RRSet{
		Profile: profile,
	}

	if _, er := recordService.Create(integration.GetTestRRSetKey(), rrSet); er.Error() != "Pool Serving Preference should be any of the following data [AUTO_SELECT SERVE_PRIMARY SERVE_ALL_FAIL] : found - TEST" {
		t.Fatal(er)
	}
}

func TestSBPoolOrderValidationFailure(t *testing.T) {
	recordService, err := record.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	profile := &sbpool.Profile{
		Order: testString,
	}
	rrSet := &rrset.RRSet{
		Profile: profile,
	}

	if _, er := recordService.Create(integration.GetTestRRSetKey(), rrSet); er.Error() != testPoolOrderError {
		t.Fatal(er)
	}
}

func TestSBPoolRecordStateValidationFailure(t *testing.T) {
	recordService, err := record.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	rdInfo := &pool.RDataInfo{
		State: testString,
	}
	profile := &sbpool.Profile{
		Order:     "FIXED",
		RDataInfo: []*pool.RDataInfo{rdInfo},
	}
	rrSet := &rrset.RRSet{
		Profile: profile,
	}

	if _, er := recordService.Create(integration.GetTestRRSetKey(), rrSet); er.Error() != "Pool record state should be any of the following data [NORMAL ACTIVE INACTIVE] : found - TEST" {
		t.Fatal(er)
	}
}

func TestTCPoolRecordStateValidationFailure(t *testing.T) {
	recordService, err := record.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	rdInfo := &pool.RDataInfo{
		State: testString,
	}
	profile := &tcpool.Profile{
		RDataInfo: []*pool.RDataInfo{rdInfo},
	}
	rrSet := &rrset.RRSet{
		Profile: profile,
	}

	if _, er := recordService.Create(integration.GetTestRRSetKey(), rrSet); er.Error() != "Pool record state should be any of the following data [NORMAL ACTIVE INACTIVE] : found - TEST" {
		t.Fatal(er)
	}
}

func TestDIRPoolConflictResolveValidationFailure(t *testing.T) {
	recordService, err := record.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	profile := &dirpool.Profile{
		ConflictResolve: testString,
	}
	rrSet := &rrset.RRSet{
		Profile: profile,
	}

	if _, er := recordService.Create(integration.GetTestRRSetKey(), rrSet); er.Error() != "DIR Pool Resolve Conflict should be any of the following data [GEO IP ] : found - TEST" {
		t.Fatal(er)
	}
}

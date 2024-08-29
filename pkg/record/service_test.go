package record_test

import (
	"testing"

	"github.com/ultradns/ultradns-go-sdk/internal/testing/integration"
	"github.com/ultradns/ultradns-go-sdk/pkg/helper"
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

const serviceErrorString = "Record service configuration failed"

const (
	testString         = "TEST"
	testPoolOrderError = "Invalid input error: { key: 'poolOrder', value: 'TEST', valid_values: [FIXED RANDOM ROUND_ROBIN] }"
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

	if _, err := record.New(conf); err.Error() != "Record service configuration failed: Missing required parameters: [ password ]" {
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

func TestListRecordWithConfigError(t *testing.T) {
	recordService := record.Service{}
	if _, _, err := recordService.List(&rrset.RRSetKey{}, &helper.QueryInfo{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestCreateRecordFailure(t *testing.T) {
	recordService, err := record.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, er := recordService.Create(integration.GetTestRRSetKey(), &rrset.RRSet{}); er.Error() != "Error while creating Record: Server error Response - { code: '70005', message: 'At least one field must be specified: rdata or profile' }: {key: 'www.non-existing-zone.com.:non-existing-zone.com.:A (1)'}" {
		t.Fatal(er)
	}
}

func TestUpdateRecordFailure(t *testing.T) {
	recordService, err := record.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, er := recordService.Update(integration.GetTestRRSetKey(), &rrset.RRSet{}); er.Error() != "Error while updating Record: Server error Response - { code: '70005', message: 'At least one field must be specified: rdata or profile' }: {key: 'www.non-existing-zone.com.:non-existing-zone.com.:A (1)'}" {
		t.Fatal(er)
	}
}

func TestPartialUpdateRecordFailure(t *testing.T) {
	recordService, err := record.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, er := recordService.PartialUpdate(integration.GetTestRRSetKey(), &rrset.RRSet{}); er.Error() != "Error while partial updating Record: Server error Response - { code: '1801', message: 'Zone does not exist in the system.' }: {key: 'www.non-existing-zone.com.:non-existing-zone.com.:A (1)'}" {
		t.Fatal(er)
	}
}

func TestReadRecordFailure(t *testing.T) {
	recordService, err := record.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, _, er := recordService.Read(integration.GetTestRRSetKey()); er.Error() != "Error while reading Record: Server error Response - { code: '1801', message: 'Zone does not exist in the system.' }: {key: 'www.non-existing-zone.com.:non-existing-zone.com.:A (1)'}" {
		t.Fatal(er)
	}
}

func TestDeleteRecordFailure(t *testing.T) {
	recordService, err := record.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, er := recordService.Delete(integration.GetTestRRSetKey()); er.Error() != "Error while deleting Record: Server error Response - { code: '1801', message: 'Zone does not exist in the system.' }: {key: 'www.non-existing-zone.com.:non-existing-zone.com.:A (1)'}" {
		t.Fatal(er)
	}
}

func TestListRecordFailure(t *testing.T) {
	recordService, err := record.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, _, er := recordService.List(integration.GetTestRRSetKey(), &helper.QueryInfo{}); er.Error() != "Error while listing Record: Server error Response - { code: '1801', message: 'Zone does not exist in the system.' }: {key: 'www.non-existing-zone.com.:non-existing-zone.com.:A (1)?&q=&offset=0&cursor=&limit=100&sort=&reverse=false'}" {
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

	if _, er := recordService.Create(integration.GetTestRRSetKey(), rrSet); er.Error() != "Invalid input error: { key: 'monitorMethod', value: 'TEST', valid_values: [GET POST] }" {
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

	if _, er := recordService.Create(integration.GetTestRRSetKey(), rrSet); er.Error() != "Invalid input error: { key: 'regionFailureSensitivity', value: 'TEST', valid_values: [HIGH LOW] }" {
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

	if _, er := recordService.Create(integration.GetTestRRSetKey(), rrSet); er.Error() != "Invalid input error: { key: 'monitorMethod', value: 'TEST', valid_values: [GET POST] }" {
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

	if _, er := recordService.Create(integration.GetTestRRSetKey(), rrSet); er.Error() != "Invalid input error: { key: 'regionFailureSensitivity', value: 'TEST', valid_values: [HIGH LOW] }" {
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

	if _, er := recordService.Create(integration.GetTestRRSetKey(), rrSet); er.Error() != "Invalid input error: { key: 'responseMethod', value: 'TEST', valid_values: [PRIORITY_HUNT RANDOM ROUND_ROBIN] }" {
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

	if _, er := recordService.Create(integration.GetTestRRSetKey(), rrSet); er.Error() != "Invalid input error: { key: 'servingPreference', value: 'TEST', valid_values: [AUTO_SELECT SERVE_PRIMARY SERVE_ALL_FAIL] }" {
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

	if _, er := recordService.Create(integration.GetTestRRSetKey(), rrSet); er.Error() != "Invalid input error: { key: 'poolRecordState', value: 'TEST', valid_values: [NORMAL ACTIVE INACTIVE] }" {
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

	if _, er := recordService.Create(integration.GetTestRRSetKey(), rrSet); er.Error() != "Invalid input error: { key: 'poolRecordState', value: 'TEST', valid_values: [NORMAL ACTIVE INACTIVE] }" {
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

	if _, er := recordService.Create(integration.GetTestRRSetKey(), rrSet); er.Error() != "Invalid input error: { key: 'dirPoolConflict', value: 'TEST', valid_values: [GEO IP ] }" {
		t.Fatal(er)
	}
}

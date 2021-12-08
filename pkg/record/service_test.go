package record_test

import (
	"fmt"
	"testing"

	"github.com/ultradns/ultradns-go-sdk/pkg/record"
	"github.com/ultradns/ultradns-go-sdk/pkg/rrset"
	"github.com/ultradns/ultradns-go-sdk/pkg/test"
)

const (
	recordTypeA = "A"
)

var (
	testPrimaryZoneName = "0-0-0-0-0rohith.com"
	testOwnerName       = ""
	testRRSetKeyA       *rrset.RRSetKey
)

func TestNewSuccess(t *testing.T) {
	conf := test.GetConfig()
	_, err := record.New(conf)

	if err != nil {
		t.Fatal(err)
	}
}

func TestNewError(t *testing.T) {
	conf := test.GetConfig()
	conf.Password = ""
	_, err := record.New(conf)

	if err.Error() != "config error while creating Record service : config validation failure: password is missing" {
		t.Fatal(err)
	}
}

func TestGetSuccess(t *testing.T) {
	_, err := record.Get(test.TestClient)

	if err != nil {
		t.Fatal(err)
	}
}

func TestGetError(t *testing.T) {
	_, err := record.Get(nil)

	if err.Error() != "Record service is not properly configured" {
		t.Fatal(err)
	}
}

func TestCreateRecordSuccess(t *testing.T) {
	recordService, err := record.Get(test.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	testOwnerName = test.GetRandomString()
	testRRSetKeyA = getRRSetKey(testOwnerName, testPrimaryZoneName, recordTypeA)
	rrSet := getRRSetTypeA(testOwnerName)

	_, er := recordService.CreateRecord(testRRSetKeyA, rrSet)

	if er != nil {
		t.Fatal(er)
	}
}

func TestCreateRecordWithConfigError(t *testing.T) {
	recordService := record.Service{}
	_, err := recordService.CreateRecord(&rrset.RRSetKey{}, &rrset.RRSet{})

	if err.Error() != "Record service is not properly configured" {
		t.Fatal(err)
	}
}

func TestCreateRecordFailure(t *testing.T) {
	recordService, err := record.Get(test.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	_, er := recordService.CreateRecord(testRRSetKeyA, &rrset.RRSet{})

	if er.Error() != fmt.Sprintf("error while creating record - %s : error code : 70005 - error message : At least one field must be specified: rdata or profile", testRRSetKeyA.ID()) {
		t.Fatal(er)
	}
}

func TestUpdateRecordSuccess(t *testing.T) {
	recordService, err := record.Get(test.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	rrSet := getRRSetTypeA(testOwnerName)
	rrSet.RData = []string{"192.168.1.2"}

	_, er := recordService.UpdateRecord(testRRSetKeyA, rrSet)

	if er != nil {
		t.Fatal(er)
	}
}

func TestUpdateRecordWithConfigError(t *testing.T) {
	recordService := record.Service{}
	_, err := recordService.UpdateRecord(&rrset.RRSetKey{}, &rrset.RRSet{})

	if err.Error() != "Record service is not properly configured" {
		t.Fatal(err)
	}
}

func TestUpdateRecordFailure(t *testing.T) {
	recordService, err := record.Get(test.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	_, er := recordService.UpdateRecord(testRRSetKeyA, &rrset.RRSet{})

	if er.Error() != fmt.Sprintf("error while updating record - %s : error code : 70005 - error message : At least one field must be specified: rdata or profile", testRRSetKeyA.ID()) {
		t.Fatal(er)
	}
}

func TestPartialUpdateRecordSuccess(t *testing.T) {
	recordService, err := record.Get(test.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	rrSet := getRRSetTypeA(testOwnerName)
	rrSet.RData = []string{"192.168.1.2"}

	_, er := recordService.PartialUpdateRecord(testRRSetKeyA, rrSet)

	if er != nil {
		t.Fatal(er)
	}
}

func TestPartialUpdateRecordWithConfigError(t *testing.T) {
	recordService := record.Service{}
	_, err := recordService.PartialUpdateRecord(&rrset.RRSetKey{}, &rrset.RRSet{})

	if err.Error() != "Record service is not properly configured" {
		t.Fatal(err)
	}
}

func TestPartialUpdateRecordFailure(t *testing.T) {
	recordService, err := record.Get(test.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	_, er := recordService.PartialUpdateRecord(testRRSetKeyA, &rrset.RRSet{TTL: -1})

	if er.Error() != fmt.Sprintf("error while partial updating record - %s : error code : 1000 - error message : Invalid TTL Format.", testRRSetKeyA.ID()) {
		t.Fatal(er)
	}
}

func TestReadRecordSuccess(t *testing.T) {
	recordService, err := record.Get(test.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	_, rrsetList, er := recordService.ReadRecord(testRRSetKeyA)

	if er != nil {
		t.Fatal(er)
	}

	if rrsetList != nil && rrsetList.ResultInfo != nil && rrsetList.ResultInfo.ReturnedCount != 1 {
		t.Fatalf("rrset returned count mismatched expected - %v : found - %v", 1, rrsetList.ResultInfo.ReturnedCount)
	}

	if rrsetList.ZoneName != testPrimaryZoneName {
		t.Fatalf("zone name mismatched expected - %v : found - %v", testPrimaryZoneName, rrsetList.ZoneName)
	}
}

func TestReadRecordWithConfigError(t *testing.T) {
	recordService := record.Service{}
	_, _, err := recordService.ReadRecord(&rrset.RRSetKey{})

	if err.Error() != "Record service is not properly configured" {
		t.Fatal(err)
	}
}

func TestReadRecordFailure(t *testing.T) {
	recordService, err := record.Get(test.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	rrSetKey := getRRSetKey(test.GetRandomString(), testPrimaryZoneName, recordTypeA)
	_, _, er := recordService.ReadRecord(rrSetKey)

	if er.Error() != fmt.Sprintf("error while reading record - %s : error code : 70002 - error message : Data not found.", rrSetKey.ID()) {
		t.Fatal(er)
	}
}

func TestDeleteRecordSuccess(t *testing.T) {
	recordService, err := record.Get(test.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	_, er := recordService.DeleteRecord(testRRSetKeyA)

	if er != nil {
		t.Fatal(er)
	}
}

func TestDeleteRecordWithConfigError(t *testing.T) {
	recordService := record.Service{}
	_, err := recordService.DeleteRecord(&rrset.RRSetKey{})

	if err.Error() != "Record service is not properly configured" {
		t.Fatal(err)
	}
}

func TestDeleteRecordFailure(t *testing.T) {
	recordService, err := record.Get(test.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	_, er := recordService.DeleteRecord(testRRSetKeyA)

	if er.Error() != fmt.Sprintf("error while deleting record - %s : error code : 56001 - error message : Cannot find resource record data for the input zone, record type and owner combination.", testRRSetKeyA.ID()) {
		t.Fatal(er)
	}
}

func getRRSetKey(ownerName, zoneName, recordType string) *rrset.RRSetKey {
	return &rrset.RRSetKey{
		Name: ownerName,
		Zone: zoneName,
		Type: recordType,
	}
}

func getRRSetTypeA(ownerName string) *rrset.RRSet {
	return &rrset.RRSet{
		OwnerName: ownerName,
		RRType:    recordTypeA,
		RData:     []string{"192.168.1.1"},
	}
}

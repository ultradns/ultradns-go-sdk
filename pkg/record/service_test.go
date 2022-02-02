package record_test

import (
	"testing"

	"github.com/ultradns/ultradns-go-sdk/internal/testing/integration"
	"github.com/ultradns/ultradns-go-sdk/pkg/record"
	"github.com/ultradns/ultradns-go-sdk/pkg/rrset"
)

const serviceErrorString = "Record service is not properly configured"

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
	if _, err := recordService.CreateRecord(&rrset.RRSetKey{}, &rrset.RRSet{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestUpdateRecordWithConfigError(t *testing.T) {
	recordService := record.Service{}

	if _, err := recordService.UpdateRecord(&rrset.RRSetKey{}, &rrset.RRSet{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestPartialUpdateRecordWithConfigError(t *testing.T) {
	recordService := record.Service{}

	if _, err := recordService.PartialUpdateRecord(&rrset.RRSetKey{}, &rrset.RRSet{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestReadRecordWithConfigError(t *testing.T) {
	recordService := record.Service{}

	if _, _, err := recordService.ReadRecord(&rrset.RRSetKey{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestDeleteRecordWithConfigError(t *testing.T) {
	recordService := record.Service{}

	if _, err := recordService.DeleteRecord(&rrset.RRSetKey{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestCreateRecordFailure(t *testing.T) {
	recordService, err := record.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, er := recordService.CreateRecord(integration.TestRRSetKey, &rrset.RRSet{}); er.Error() != "error while creating Record - www.non-existing-zone.com.:non-existing-zone.com.:A (1) : error from api response - error code : 70005 - error message : At least one field must be specified: rdata or profile" {
		t.Fatal(er)
	}
}

func TestUpdateRecordFailure(t *testing.T) {
	recordService, err := record.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, er := recordService.UpdateRecord(integration.TestRRSetKey, &rrset.RRSet{}); er.Error() != "error while updating Record - www.non-existing-zone.com.:non-existing-zone.com.:A (1) : error from api response - error code : 70005 - error message : At least one field must be specified: rdata or profile" {
		t.Fatal(er)
	}
}

func TestPartialUpdateRecordFailure(t *testing.T) {
	recordService, err := record.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, er := recordService.PartialUpdateRecord(integration.TestRRSetKey, &rrset.RRSet{}); er.Error() != "error while partial updating Record - www.non-existing-zone.com.:non-existing-zone.com.:A (1) : error from api response - error code : 1801 - error message : Zone does not exist in the system." {
		t.Fatal(er)
	}
}

func TestReadRecordFailure(t *testing.T) {
	recordService, err := record.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, _, er := recordService.ReadRecord(integration.TestRRSetKey); er.Error() != "error while reading Record - www.non-existing-zone.com.:non-existing-zone.com.:A (1) : error from api response - error code : 1801 - error message : Zone does not exist in the system." {
		t.Fatal(er)
	}
}

func TestDeleteRecordFailure(t *testing.T) {
	recordService, err := record.Get(integration.TestClient)

	if err != nil {
		t.Fatal(err)
	}

	if _, er := recordService.DeleteRecord(integration.TestRRSetKey); er.Error() != "error while deleting Record - www.non-existing-zone.com.:non-existing-zone.com.:A (1) : error from api response - error code : 1801 - error message : Zone does not exist in the system." {
		t.Fatal(er)
	}
}

func TestRRSetKeyURI(t *testing.T) {
	rrSetKey := rrset.RRSetKey{
		Zone: "a",
		Name: "b",
	}

	if expectedURI, foundURI := "zones/a/rrsets/ANY/b", rrSetKey.URI(); expectedURI != foundURI {
		t.Fatalf("uri mismatched expected - %s : found - %s", expectedURI, foundURI)
	}
}

func TestRRSetKeyID(t *testing.T) {
	rrSetKey := rrset.RRSetKey{
		Zone: "example.com",
		Name: "www",
		Type: "A",
	}

	if expectedID, foundID := "www.example.com.:example.com.:A (1)", rrSetKey.ID(); expectedID != foundID {
		t.Fatalf("rrset id mismatched expected - %s : found - %s", expectedID, foundID)
	}
}

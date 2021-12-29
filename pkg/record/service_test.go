package record_test

import (
	"testing"

	"github.com/ultradns/ultradns-go-sdk/internal/test"
	"github.com/ultradns/ultradns-go-sdk/pkg/record"
	"github.com/ultradns/ultradns-go-sdk/pkg/rrset"
)

const serviceErrorString = "Record service is not properly configured"

func TestNewSuccess(t *testing.T) {
	conf := test.GetConfig()

	if _, err := record.New(conf); err != nil {
		t.Fatal(err)
	}
}

func TestNewError(t *testing.T) {
	conf := test.GetConfig()
	conf.Password = ""

	if _, err := record.New(conf); err.Error() != "config error while creating Record service : config validation failure: password is missing" {
		t.Fatal(err)
	}
}

func TestGetSuccess(t *testing.T) {
	if _, err := record.Get(test.TestClient); err != nil {
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

func TestRRSetKeyURI(t *testing.T) {
	rrSetKey := rrset.RRSetKey{
		Zone: "a",
		Name: "b",
	}

	expectedURI := "zones/a/rrsets/ANY/b"

	if foundURI := rrSetKey.URI(); expectedURI != foundURI {
		t.Fatalf("uri mismatched expected - %s : found - %s", expectedURI, foundURI)
	}
}

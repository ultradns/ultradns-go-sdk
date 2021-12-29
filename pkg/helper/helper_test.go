package helper_test

import (
	"testing"

	"github.com/ultradns/ultradns-go-sdk/pkg/helper"
)

func TestGetRecordTypeString(t *testing.T) {
	if expected, found := "A", helper.GetRecordTypeString("A (1)"); expected != found {
		t.Fatal("record type mismatched")
	}
}

func TestGetRecordTypeNumber(t *testing.T) {
	if expected, found := "5", helper.GetRecordTypeNumber("CNAME (5)"); expected != found {
		t.Fatal("record type mismatched")
	}
}

package webforward_test

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"github.com/ultradns/ultradns-go-sdk/pkg/webforward"
)

func TestWebForwardUnmarshal(t *testing.T) {
	raw := `{"guid":"1AA2BB3CC4DD5EE6","requestTo":"*.example.com/*","defaultRedirectTo":"https://example.com/en-us/","defaultForwardType":"HTTP_301_REDIRECT"}`

	var wf webforward.WebForward

	if err := json.Unmarshal([]byte(raw), &wf); err != nil {
		t.Fatal(err)
	}

	if wf.GUID != "1AA2BB3CC4DD5EE6" {
		t.Fatalf("guid mismatched expected - 1AA2BB3CC4DD5EE6 : found - %v", wf.GUID)
	}

	if wf.RequestTo != "*.example.com/*" {
		t.Fatalf("requestTo mismatched expected - *.example.com/* : found - %v", wf.RequestTo)
	}

	if wf.DefaultRedirectTo != "https://example.com/en-us/" {
		t.Fatalf("defaultRedirectTo mismatched expected - https://example.com/en-us/ : found - %v", wf.DefaultRedirectTo)
	}

	if wf.DefaultForwardType != webforward.HTTP301Redirect {
		t.Fatalf("defaultForwardType mismatched expected - HTTP_301_REDIRECT : found - %v", wf.DefaultForwardType)
	}
}

func TestHTTPSWebForwardUnmarshal(t *testing.T) {
	raw := `{
		"guid": "09084E85FE7773E2",
		"requestTo": "https://source.example.com/*",
		"defaultRedirectTo": "https://target.example.com/",
		"defaultForwardType": "HTTP_301_REDIRECT",
		"defaultRedirectType": "HTTPS",
		"certificateId": "certificate-guid",
		"certificateName": "Example certificate",
		"expirationDays": "45",
		"state": "Pending",
		"errorDescription": "No DNS TXT record found for example.com."
	}`

	var wf webforward.WebForward

	if err := json.Unmarshal([]byte(raw), &wf); err != nil {
		t.Fatal(err)
	}

	if wf.DefaultRedirectType != webforward.HTTPSRedirect {
		t.Fatalf("defaultRedirectType mismatched expected - HTTPS : found - %v", wf.DefaultRedirectType)
	}

	if wf.CertificateID != "certificate-guid" {
		t.Fatalf("certificateId mismatched expected - certificate-guid : found - %v", wf.CertificateID)
	}

	if wf.CertificateName != "Example certificate" {
		t.Fatalf("certificateName mismatched expected - Example certificate : found - %v", wf.CertificateName)
	}

	if wf.ExpirationDays != "45" {
		t.Fatalf("expirationDays mismatched expected - 45 : found - %v", wf.ExpirationDays)
	}

	if wf.State != "Pending" {
		t.Fatalf("state mismatched expected - Pending : found - %v", wf.State)
	}

	if wf.ErrorDescription == "" {
		t.Fatal("errorDescription should be preserved")
	}
}

func TestHTTPSWebForwardMarshal(t *testing.T) {
	wf := &webforward.WebForward{
		RequestTo:              "https://source.example.com/*",
		DefaultRedirectTo:      "https://target.example.com/",
		DefaultForwardType:     webforward.HTTP301Redirect,
		CertificateManagedType: webforward.EECertificateManagedType,
	}

	data, err := json.Marshal(wf)
	if err != nil {
		t.Fatal(err)
	}

	var payload map[string]interface{}
	if err := json.Unmarshal(data, &payload); err != nil {
		t.Fatal(err)
	}

	if payload["requestTo"] != wf.RequestTo {
		t.Fatalf("requestTo mismatched expected - %v : found - %v", wf.RequestTo, payload["requestTo"])
	}

	if payload["certificateManagedType"] != webforward.EECertificateManagedType {
		t.Fatalf("certificateManagedType mismatched expected - EE : found - %v", payload["certificateManagedType"])
	}

	for _, field := range []string{"certificateId", "certificateName", "expirationDays", "state", "errorDescription"} {
		if _, ok := payload[field]; ok {
			t.Fatalf("%s should be omitted from a request when empty, encoded: %s", field, data)
		}
	}
}

func TestHTTPSWebForwardRoundTrip(t *testing.T) {
	want := &webforward.WebForward{
		GUID:                "09084E85FE7773E2",
		RequestTo:           "https://source.example.com/*",
		DefaultRedirectTo:   "https://target.example.com/",
		DefaultForwardType:  webforward.HTTP302Redirect,
		DefaultRedirectType: webforward.HTTPSRedirect,
		CertificateID:       "certificate-guid",
		CertificateName:     "Example certificate",
		ZoneName:            "example.com.",
		AccountName:         "example-account",
		State:               "Active",
		ExpirationDays:      "45",
		Records: []webforward.Record{{
			RedirectTo:  "http://fallback.example.com/",
			ForwardType: webforward.HTTP302Redirect,
			Priority:    1,
			Rules: []webforward.Rule{{
				Header:          "Accept",
				MatchCriteria:   "CONTAINS",
				Value:           "text/html",
				CaseInsensitive: true,
			}},
		}},
	}

	data, err := json.Marshal(want)
	if err != nil {
		t.Fatal(err)
	}

	var got webforward.WebForward
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(want, &got) {
		t.Fatalf("HTTPS web forward changed during round trip:\nwant: %+v\ngot:  %+v", want, got)
	}
}

func TestWebForwardMarshalOmitsAbsentFields(t *testing.T) {
	wf := &webforward.WebForward{
		RequestTo:          "www.example.com",
		DefaultRedirectTo:  "https://example.com/",
		DefaultForwardType: webforward.HTTP302Redirect,
	}

	data, err := json.Marshal(wf)

	if err != nil {
		t.Fatal(err)
	}

	encoded := string(data)

	if strings.Contains(encoded, "relativeForwardType") {
		t.Fatalf("relativeForwardType should be omitted when empty, encoded: %s", encoded)
	}

	if strings.Contains(encoded, `"guid"`) {
		t.Fatalf("guid should be omitted when empty, encoded: %s", encoded)
	}
}

func TestRelativeForwardTypeRoundTrip(t *testing.T) {
	wf := &webforward.WebForward{
		RequestTo:           "www.example.com",
		DefaultRedirectTo:   "https://example.com/",
		DefaultForwardType:  webforward.HTTP301Redirect,
		RelativeForwardType: webforward.ParameterAndPath,
	}

	data, err := json.Marshal(wf)

	if err != nil {
		t.Fatal(err)
	}

	var decoded webforward.WebForward

	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatal(err)
	}

	if decoded.RelativeForwardType != webforward.ParameterAndPath {
		t.Fatalf("relativeForwardType mismatched expected - PARAMETER_AND_PATH : found - %v", decoded.RelativeForwardType)
	}
}

func TestResponseListUnmarshal(t *testing.T) {
	raw := `{
		"queryInfo": {"sort": "REQUEST_TO", "reverse": false, "limit": 100},
		"resultInfo": {"totalCount": 2, "offset": 0, "returnedCount": 2},
		"webForwards": [
			{"guid":"090348E6F8B46D80","requestTo":"example.com","defaultRedirectTo":"https://www.example.com/en-us/","defaultForwardType":"HTTP_301_REDIRECT"},
			{"guid":"090348E6F8B4D727","requestTo":"www.example.com","defaultRedirectTo":"https://www.example.com/en-us/","defaultForwardType":"HTTP_302_REDIRECT","relativeForwardType":"PARAMETER_AND_PATH"}
		]
	}`

	var list webforward.ResponseList

	if err := json.Unmarshal([]byte(raw), &list); err != nil {
		t.Fatal(err)
	}

	if len(list.WebForwards) != 2 {
		t.Fatalf("webForward count mismatched expected - 2 : found - %v", len(list.WebForwards))
	}

	if list.WebForwards[1].RelativeForwardType != webforward.ParameterAndPath {
		t.Fatalf("relativeForwardType mismatched expected - PARAMETER_AND_PATH : found - %v", list.WebForwards[1].RelativeForwardType)
	}

	if list.ResultInfo.TotalCount != 2 {
		t.Fatalf("totalCount mismatched expected - 2 : found - %v", list.ResultInfo.TotalCount)
	}
}

func TestResponseListUnmarshalHTTPSWebForward(t *testing.T) {
	raw := `{
		"resultInfo": {"totalCount": 1, "offset": 0, "returnedCount": 1},
		"webForwards": [{
			"guid": "09084E85FE7773E2",
			"requestTo": "https://source.example.com",
			"defaultRedirectTo": "https://target.example.com",
			"defaultForwardType": "HTTP_301_REDIRECT",
			"defaultRedirectType": "HTTPS",
			"zoneName": "example.com.",
			"accountName": "example-account",
			"certificateManagedType": "EE",
			"certificateId": "certificate-guid",
			"certificateName": "Example certificate",
			"expirationDays": "45",
			"state": "Pending",
			"errorDescription": "Pending validation"
		}]
	}`

	var list webforward.ResponseList
	if err := json.Unmarshal([]byte(raw), &list); err != nil {
		t.Fatal(err)
	}

	if len(list.WebForwards) != 1 {
		t.Fatalf("webForward count mismatched expected - 1 : found - %v", len(list.WebForwards))
	}

	got := list.WebForwards[0]
	if got.DefaultRedirectType != webforward.HTTPSRedirect ||
		got.CertificateManagedType != webforward.EECertificateManagedType ||
		got.ExpirationDays != "45" ||
		got.ZoneName != "example.com." ||
		got.AccountName != "example-account" {
		t.Fatalf("HTTPS web forward fields were not preserved: %+v", got)
	}
}

func TestResponseListUnmarshalEmpty(t *testing.T) {
	var list webforward.ResponseList

	if err := json.Unmarshal([]byte(`{"webForwards": []}`), &list); err != nil {
		t.Fatal(err)
	}

	if len(list.WebForwards) != 0 {
		t.Fatalf("webForward count mismatched expected - 0 : found - %v", len(list.WebForwards))
	}
}

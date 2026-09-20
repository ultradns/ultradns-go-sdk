package webforward_test

import (
	"encoding/json"
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

func TestResponseListUnmarshalEmpty(t *testing.T) {
	var list webforward.ResponseList

	if err := json.Unmarshal([]byte(`{"webForwards": []}`), &list); err != nil {
		t.Fatal(err)
	}

	if len(list.WebForwards) != 0 {
		t.Fatalf("webForward count mismatched expected - 0 : found - %v", len(list.WebForwards))
	}
}

package batch_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/ultradns/ultradns-go-sdk/pkg/batch"
)

const serviceErrorString = "Batch service configuration failed"

func TestGetError(t *testing.T) {
	if _, err := batch.Get(nil); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestExecuteWithConfigError(t *testing.T) {
	batchService := batch.Service{}

	if _, _, err := batchService.Execute([]*batch.Action{}); err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestActionListMarshalIsBareArray(t *testing.T) {
	actions := []*batch.Action{
		{
			Method: "POST",
			URI:    "/v1/zones/example.com/rrsets/A/www",
			Body: map[string]interface{}{
				"ttl":   300,
				"rdata": []string{"1.2.3.4"},
			},
		},
		{
			Method: "DELETE",
			URI:    "/v1/zones/example.com/rrsets/A/old",
		},
	}

	data, err := json.Marshal(actions)

	if err != nil {
		t.Fatal(err)
	}

	encoded := string(data)

	if !strings.HasPrefix(encoded, "[") {
		t.Fatalf("batch request must marshal to a bare JSON array, got: %s", encoded)
	}

	if !strings.Contains(encoded, `"method":"POST"`) {
		t.Fatalf("method mismatched, encoded: %s", encoded)
	}

	if !strings.Contains(encoded, `"method":"DELETE"`) {
		t.Fatalf("missing delete action, encoded: %s", encoded)
	}
}

func TestResponseListUnmarshal(t *testing.T) {
	raw := `[
		{"status": 201, "response": {"message": "SUCCESSFUL"}},
		{"status": 204}
	]`

	var result []*batch.Response

	if err := json.Unmarshal([]byte(raw), &result); err != nil {
		t.Fatal(err)
	}

	if len(result) != 2 {
		t.Fatalf("response count mismatched expected - 2 : found - %v", len(result))
	}

	if result[0].Status != 201 {
		t.Fatalf("status mismatched expected - 201 : found - %v", result[0].Status)
	}

	if result[1].Status != 204 {
		t.Fatalf("status mismatched expected - 204 : found - %v", result[1].Status)
	}
}

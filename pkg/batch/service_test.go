package batch_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ultradns/ultradns-go-sdk/pkg/batch"
	"github.com/ultradns/ultradns-go-sdk/pkg/client"
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

func TestGetSuccess(t *testing.T) {
	c, err := client.NewClient(client.Config{
		Username: "username",
		Password: "password",
		HostURL:  "http://127.0.0.1",
	})

	if err != nil {
		t.Fatal(err)
	}

	if _, err := batch.Get(c); err != nil {
		t.Fatal(err)
	}
}

func TestExecuteSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/authorization/token":
			w.Header().Set("Content-Type", "application/json")
			_, _ = io.WriteString(w, `{"access_token":"test-token","token_type":"Bearer","expires_in":3600}`)
		case r.Method == http.MethodPost && r.URL.Path == "/v1/batch":
			var actions []*batch.Action
			if err := json.NewDecoder(r.Body).Decode(&actions); err != nil {
				t.Errorf("decode batch actions: %v", err)
			}
			if len(actions) != 1 ||
				actions[0].Method != http.MethodPost ||
				actions[0].URI != "/v1/zones/example.com./rrsets/A/www" {
				t.Errorf("unexpected batch actions: %+v", actions)
			}
			w.Header().Set("Content-Type", "application/json")
			_, _ = io.WriteString(w, `[{"status":201,"response":{"message":"SUCCESSFUL"}}]`)
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	service, err := batch.New(client.Config{
		Username: "username",
		Password: "password",
		HostURL:  server.URL,
	})

	if err != nil {
		t.Fatal(err)
	}

	res, responses, err := service.Execute([]*batch.Action{{
		Method: http.MethodPost,
		URI:    "/v1/zones/example.com./rrsets/A/www",
	}})

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("status code mismatched expected - %d : found - %d", http.StatusOK, res.StatusCode)
	}

	if len(responses) != 1 || responses[0].Status != http.StatusCreated {
		t.Fatalf("batch responses mismatched: %+v", responses)
	}
}

func TestExecuteWithServerError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/authorization/token" {
			w.Header().Set("Content-Type", "application/json")
			_, _ = io.WriteString(w, `{"access_token":"test-token","token_type":"Bearer","expires_in":3600}`)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = io.WriteString(w, `{"errorCode":500,"errorMessage":"Internal Server Error"}`)
	}))
	defer server.Close()

	service, err := batch.New(client.Config{
		Username: "username",
		Password: "password",
		HostURL:  server.URL,
	})

	if err != nil {
		t.Fatal(err)
	}

	_, _, err = service.Execute([]*batch.Action{{
		Method: http.MethodPost,
		URI:    "/v1/zones/example.com./rrsets/A/www",
	}})

	expected := "Error while creating Batch: Server error Response - { code: '500', message: 'Internal Server Error' }: {key: 'batch'}"

	if err == nil || err.Error() != expected {
		t.Fatalf("execute error mismatched expected - %v : found - %v", expected, err)
	}
}

package provider_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ultradns/ultradns-go-sdk/pkg/cdn/provider"
	"github.com/ultradns/ultradns-go-sdk/pkg/client"
)

const serviceErrorString = "CDNProvider service configuration failed"

func TestNewError(t *testing.T) {
	if _, err := provider.New(client.Config{}); err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestGetError(t *testing.T) {
	if _, err := provider.Get(nil); err == nil || err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestCreateWithConfigError(t *testing.T) {
	svc := provider.Service{}
	if _, err := svc.Create("acc1", &provider.Provider{ClientCdnID: "cdn-a"}); err == nil || err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestCreateWithNilPayload(t *testing.T) {
	svc, err := provider.New(client.Config{Username: "u", Password: "p", HostURL: "https://example.com"})
	if err != nil {
		t.Fatal(err)
	}

	if _, err = svc.Create("acc1", nil); err == nil || err.Error() != "Missing required parameters: [payload ]" {
		t.Fatal(err)
	}
}

func TestCreateWithEmptyClientCdnID(t *testing.T) {
	svc, err := provider.New(client.Config{Username: "u", Password: "p", HostURL: "https://example.com"})
	if err != nil {
		t.Fatal(err)
	}

	if _, err = svc.Create("acc1", &provider.Provider{ClientCdnID: "  "}); err == nil || err.Error() != "Missing required parameters: [clientCdnId ]" {
		t.Fatal(err)
	}
}

func TestCreateTrimsClientCdnIDBeforeRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/accounts/acc1/cdn_providers/cdn-a" {
			t.Fatalf("expected trimmed request path, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"message":"ok"}`))
	}))
	defer server.Close()

	svc, err := provider.New(client.Config{Username: "u", Password: "p", HostURL: server.URL})
	if err != nil {
		t.Fatal(err)
	}

	payload := &provider.Provider{ClientCdnID: "  cdn-a  "}
	if _, err = svc.Create("acc1", payload); err != nil {
		t.Fatal(err)
	}
	if payload.ClientCdnID != "cdn-a" {
		t.Fatalf("expected payload ClientCdnID to be normalized, got %q", payload.ClientCdnID)
	}
}

func TestReadWithConfigError(t *testing.T) {
	svc := provider.Service{}
	if _, _, err := svc.Read("acc1", "cdn-a"); err == nil || err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestUpdateWithConfigError(t *testing.T) {
	svc := provider.Service{}
	if _, err := svc.Update("acc1", "cdn-a", &provider.Provider{}); err == nil || err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestDeleteWithConfigError(t *testing.T) {
	svc := provider.Service{}
	if _, err := svc.Delete("acc1", "cdn-a"); err == nil || err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}

func TestListWithConfigError(t *testing.T) {
	svc := provider.Service{}
	if _, _, err := svc.List("acc1"); err == nil || err.Error() != serviceErrorString {
		t.Fatal(err)
	}
}